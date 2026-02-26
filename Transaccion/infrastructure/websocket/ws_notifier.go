package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TransferPayload struct {
	Success  bool  `json:"success"`
}

func NotifyTransfer(success bool) error {
	payload := TransferPayload{
		Success : success,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("❌ Error marshaling payload: %s\n", err.Error())
		return err
	}

	wsServerURL := os.Getenv("WS_SERVER_URL")
	fmt.Printf("📡 Notificando a: %s/transfer/notify\n", wsServerURL)
	fmt.Printf("📦 Payload: %s\n", string(body))
 
	resp, err := http.Post(
		fmt.Sprintf("%s/transfer/notify", wsServerURL),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {

		fmt.Printf("❌ Error notificando al WS server: %s\n", err.Error())
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("✅ WS server respondió con status: %d\n", resp.StatusCode)
	return nil
}