package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	bancoinfra "banco-api/Banco/infrastructure"
	cuentainfra "banco-api/Cuenta/infrastructure"
	transaccioninfra "banco-api/Transaccion/infrastructure"
	usuarioinfra "banco-api/Usuario/infrastructure"
	"banco-api/core"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró archivo .env, usando variables de entorno del sistema")
	}

	dbConfig := core.GetDatabaseConfig()
	db, err := core.GetDBPool(dbConfig)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	router.Use(core.SetupCORS())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now(),
		})
	})


	bancoDeps := bancoinfra.NewBancoDependencies(db)
	usuarioDeps := usuarioinfra.NewUsuarioDependencies(db)
	cuentaDeps := cuentainfra.NewCuentaDependencies(db)
	transaccionDeps := transaccioninfra.NewTransaccionDependencies(db)


	bancoDeps.Setup(router)
	usuarioDeps.Setup(router)
	cuentaDeps.Setup(router)
	transaccionDeps.Setup(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}


	go func() {
		log.Printf("Servidor iniciado en puerto %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error iniciando servidor: %v", err)
		}
	}()


	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Recibida señal de terminación, cerrando servidor...")


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Error durante shutdown: %v", err)
	}


	bancoDeps.Shutdown()
	usuarioDeps.Shutdown()
	cuentaDeps.Shutdown()
	transaccionDeps.Shutdown()

	log.Println("Servidor cerrado correctamente")
}
