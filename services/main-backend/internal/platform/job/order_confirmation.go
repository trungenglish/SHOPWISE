package job

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"strings"

	"shopwise/retail/internal/platform/config"

	"github.com/hibiken/asynq"
)

func HandleOrderConfirmation(cfg *config.Config) func(*asynq.Task) error {
	return func(task *asynq.Task) error {
		var payload OrderConfirmationPayload
		if err := json.Unmarshal(task.Payload(), &payload); err != nil {
			return fmt.Errorf("decode order confirmation: %w", err)
		}
		body := fmt.Sprintf("Hello %s,\n\nInvoice: %s\nTotal: %d VND\nDelivery estimate: %s to %s\n\n%s\n\nSupport: %s\n", payload.CustomerName, payload.OrderID, payload.TotalAmount, payload.DeliveryFrom.Format("2006-01-02"), payload.DeliveryTo.Format("2006-01-02"), strings.Join(payload.Items, "\n"), cfg.SMTPFrom)
		address := cfg.SMTPHost + ":" + cfg.SMTPPort
		var auth smtp.Auth
		if cfg.SMTPUser != "" {
			auth = smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)
		}
		if err := smtp.SendMail(address, auth, cfg.SMTPFrom, []string{payload.CustomerEmail}, []byte("Subject: ShopWise order confirmation\r\n\r\n"+body)); err != nil {
			return fmt.Errorf("send order confirmation: %w", err)
		}
		return nil
	}
}
