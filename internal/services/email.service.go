package services

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendOTPEmail(to string, otp string) error {
	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	if smtpPort == "" {
		smtpPort = "587" // default fallback
	}

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Your MoniVest Verification Code"

	htmlBody := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>
<body style="margin:0;padding:0;background-color:#f4f6f8;font-family:Arial,sans-serif;">
  <table width="100%%" cellpadding="0" cellspacing="0" style="padding:20px;">
    <tr>
      <td align="center">
        <table width="500" cellpadding="0" cellspacing="0" style="background:#ffffff;border-radius:12px;overflow:hidden;">
          
          <!-- Header -->
          <tr>
            <td style="background:#0A2540;color:#ffffff;padding:20px;text-align:center;font-size:22px;font-weight:bold;">
              MoniVest
            </td>
          </tr>

          <!-- Body -->
          <tr>
            <td style="padding:30px;color:#333;">
              <h2 style="margin-top:0;">Verify Your Email</h2>
              <p style="font-size:14px;line-height:1.6;">
                Use the verification code below to complete your registration.
              </p>

              <!-- OTP Box -->
              <div style="margin:25px 0;text-align:center;">
                <span style="
                  display:inline-block;
                  background:#f0f4f8;
                  padding:15px 25px;
                  font-size:28px;
                  letter-spacing:5px;
                  font-weight:bold;
                  border-radius:8px;
                  color:#0A2540;
                ">
                  %s
                </span>
              </div>

              <p style="font-size:13px;color:#666;">
                This code expires in <strong>5 minutes</strong>.
              </p>

              <p style="font-size:13px;color:#666;">
                If you didn’t request this, you can safely ignore this email.
              </p>
            </td>
          </tr>

          <!-- Footer -->
          <tr>
            <td style="background:#f4f6f8;padding:15px;text-align:center;font-size:12px;color:#999;">
              © %d MoniVest. All rights reserved.
            </td>
          </tr>

        </table>
      </td>
    </tr>
  </table>
</body>
</html>
`, otp, getYear())

	msg := []byte(
		"From: MoniVest <" + from + ">\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
			htmlBody,
	)

	return smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{to},
		msg,
	)
}

func getYear() int {
	return 2026
}