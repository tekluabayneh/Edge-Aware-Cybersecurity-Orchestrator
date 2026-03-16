package notification

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sendinblue/APIv3-go-library/lib"
	sib_api_v3_sdk "github.com/sendinblue/APIv3-go-library/lib"
)

func EmailSender(senderName string, senderEmail string, recipientEmail string, subject string, htmlContent string) (lib.CreateSmtpEmail, *http.Response, error) {
	apiKey := os.Getenv("BREVO_API_KEY")

	cfg := sib_api_v3_sdk.NewConfiguration()
	cfg.AddDefaultHeader("api-key", apiKey)

	client := sib_api_v3_sdk.NewAPIClient(cfg)

	email := sib_api_v3_sdk.SendSmtpEmail{
		Sender: &sib_api_v3_sdk.SendSmtpEmailSender{
			Name:  senderName,
			Email: senderEmail,
		},
		To: []sib_api_v3_sdk.SendSmtpEmailTo{
			{Email: recipientEmail},
		},
		Subject:     subject,
		HtmlContent: htmlContent,
	}

	smtRes, httpRes, err := client.TransactionalEmailsApi.SendTransacEmail(context.Background(), email)
	fmt.Println(err)

	return smtRes, httpRes, err
}
