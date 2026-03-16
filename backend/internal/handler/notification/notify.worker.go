package notification

import (
	"fmt"
	"net/http"

	"github.com/sendinblue/APIv3-go-library/lib"
)

func Notify(userName string, ReciverEmail string, message string) (lib.CreateSmtpEmail, *http.Response, error) {

	html := fmt.Sprintf(`
<div style="font-family: Arial, sans-serif; background-color:#f7f8fa; padding:30px;">
  
  <div style="max-width:500px; margin:auto; background:#ffffff; border-radius:8px; padding:25px; box-shadow:0 2px 8px rgba(0,0,0,0.05);">

    <h2 style="color:#1f2937; margin-bottom:10px;">
      Edge-Aware Security Alert
    </h2>

    <p style="color:#4b5563; font-size:14px;">
      Hello <strong>%s</strong>,
    </p>

    <p style="color:#4b5563; font-size:14px;">
      Your monitoring agent has detected a new security event on your system.
    </p>

    <div style="background:#f3f4f6; padding:15px; border-radius:6px; margin:15px 0;">
      <p style="margin:0; font-size:14px; color:#111827;">
        <strong>Alert Message:</strong>
      </p>
      <p style="margin-top:6px; font-size:14px; color:#374151;">
        %s
      </p>
    </div>

    <p style="color:#6b7280; font-size:13px;">
      Please review the alert in your dashboard to investigate further.
    </p>

    <hr style="border:none; border-top:1px solid #e5e7eb; margin:20px 0;">

    <p style="font-size:12px; color:#9ca3af;">
      This notification was generated automatically by your Edge-Aware monitoring agent.
    </p>

  </div>

</div>
`, userName, message)

	smtRes, res, err := EmailSender(userName, ReciverEmail, ReciverEmail, "Agent Security Alert", html)
	return smtRes, res, err
}
