package smtp

import (
	"fmt"
	"net/smtp"
)

type MailHogService struct {
	auth smtp.Auth
	host string
	port string
	senderEmail string
}

func NewMailHogService(host string, port string, senderEmail string, password string) *MailHogService {
	auth := smtp.PlainAuth("", senderEmail, password, host)
	return &MailHogService{
		auth: auth,
		host: host,
		port: port,
		senderEmail: senderEmail,
	}
}

func (s *MailHogService) SendOTP(toEmail string, otpCode string) error {
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	subject := "Subject: Your Login Code\r\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`
		<html>
			<body>
				<h3>Your Verification Code</h3>
				<h1 style="color: blue;">%s</h1>
			</body>
		</html>
	`, otpCode)

	msg := []byte(subject + mime + body)

	err := smtp.SendMail(addr, s.auth, s.senderEmail, []string{toEmail}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
