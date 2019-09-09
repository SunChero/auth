package service

import (
	"bytes"
	"fmt"
	"net"
	"net/smtp"
	"net/url"
	"os"
	"text/template"
)

var magicLinkMailTmpl *template.Template

//SendVerificationCode this should be moved to rabbitmq
func (s *Service) SendVerificationCode(code string, email string) error {

	// if magicLinkMailTmpl == nil {
	magicLinkMailTmpl, _ = template.ParseFiles("mail.template.html")
	// 	if err != nil {
	// 		return fmt.Errorf("could not parse magic link mail template: %v", err)
	// 	}
	// }
	link, _ := url.Parse(os.Getenv("ROOT_URL"))
	link.Path = "/auth/login_with_code"
	q := link.Query()
	q.Set("verification_code", code)
	link.RawQuery = q.Encode()
	var b bytes.Buffer
	if err := magicLinkMailTmpl.Execute(&b, map[string]interface{}{
		"MagicLink": link,
		"Minutes":   int(verificationCodeLifespan.Minutes()),
		"Code":      code,
	}); err != nil {
		return fmt.Errorf("could not execute magic link mail template: %v", err)
	}

	// sender := mailing.NewSMTPSender(
	// 	"noreply@laval.ca", //,
	// 	"mail.laval.ca",    //smtpHost,
	// 	"25",               //strconv.Itoa(smtpPort),
	// 	"",                 //smtpUsername,
	// 	"",                 //smtpPassword,
	// )
	if err := sendMail(email, "Magic Link", b.String()); err != nil {
		return fmt.Errorf("could not send magic link: %v", err)
	}

	return nil
}

func sendMail(dst string, title string, msg string) error {
	body := fmt.Sprintf("From: %s\r\n", "noreply@laval.ca") +
		fmt.Sprintf("To: %s\r\n", dst) +
		fmt.Sprintf("Subject: %s\r\n", title) +
		"Content-Type: text/html; charset=utf-8\r\n" +
		"\r\n" + msg
	return smtp.SendMail(
		net.JoinHostPort(os.Getenv("SMTP_SERVER"), os.Getenv("SMTP_PORT")),
		nil,
		"auth@laval.ca",
		[]string{dst},
		[]byte(body),
	)
}
