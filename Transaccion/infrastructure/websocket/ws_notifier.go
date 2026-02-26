package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type TransferPayload struct {
	TransferID  string  `json:"transferId"`
	FromAccount string  `json:"fromAccount"`
	ToAccount   string  `json:"toAccount"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Status      string  `json:"status"`
	Timestamp   string  `json:"timestamp"`
}

func NotifyTransfer(transactionID uuid.UUID, fromAccount, toAccount string, amount float64, currency string) error {
	payload := TransferPayload{
		TransferID:  transactionID.String(),
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Amount:      amount,
		Currency:    currency,
		Status:      "success",
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
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