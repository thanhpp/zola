package service

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"

	"github.com/vfluxus/dvergr/logger"
	"github.com/vfluxus/mailservice/repository/entity"
)

type SMTPService struct{}

func GetSMTP() *SMTPService {
	return &SMTPService{}
}

func (s *SMTPService) srvError(op string, err error) error {
	return fmt.Errorf("srv: smtp. Op: %s. Err: %v", op, err)
}

// ------------------------------
// SendEmail ...
func (s *SMTPService) SendEmail(from *entity.MailAccount, mail *entity.Mail, htmlBody []byte) (err error) {
	// pre-exec check
	if from == nil || mail == nil {
		return s.srvError("input check", errors.New("nil account or mail"))
	}

	var msg = new(bytes.Buffer)
	// subject
	msg.WriteString(fmt.Sprintf("Subject: %s\n", mail.Subject))
	// mime
	msg.WriteString("MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n")
	// html body
	msg.Write(htmlBody)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         from.SMTPHost,
	}
	addr := fmt.Sprintf("%s:%s", from.SMTPHost, from.SMTPPort)
	// conn, err := tls.Dial("tcp", addr, tlsConfig)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		logger.Get().Errorf("Dial to mail server error: %v", err)
		return
	}

	// send email
	c, err := smtp.NewClient(conn, from.SMTPHost)
	if err != nil {
		logger.Get().Errorf("Create smtp client error: %v", err)
		return
	}

	if err = c.StartTLS(tlsConfig); err != nil {
		logger.Get().Errorf("StartTLS error: %v", err)
		return
	}

	// Authentication
	// auth := smtp.PlainAuth("", from.Username, from.Password, from.SMTPHost)
	auth := LoginAuth(strings.Split(from.Username, "@")[0], from.Password)
	if err = c.Auth(auth); err != nil {
		logger.Get().Errorf("Authenticate error: %v", err)
		return
	}

	if err = c.Mail(from.Username); err != nil {
		logger.Get().Errorf("Create mail error: %v", err)
		return
	}

	for _, receive := range mail.SendTo {
		if err = c.Rcpt(receive); err != nil {
			logger.Get().Errorf("Create RCPT cmd error: %v", err)
			return
		}
	}

	writer, err := c.Data()
	if err != nil {
		logger.Get().Errorf("Create email writer error: %v", err)
		return
	}

	_, err = writer.Write(msg.Bytes())
	if err != nil {
		logger.Get().Errorf("Write message error: %v", err)
		return
	}

	err = writer.Close()
	if err != nil {
		logger.Get().Errorf("Close email writer error: %v", err)
		return
	}

	c.Quit()

	// if err := smtp.SendMail(addr, auth, from.Username, mail.SendTo, msg.Bytes()); err != nil {
	// 	return s.srvError("SendEmail", err)
	// }

	return nil
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}
