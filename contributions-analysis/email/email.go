package email

import (
	"gopkg.in/gomail.v2"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Service interface {
	SendEmail(to, subject, body string) error
}

type service struct {
	config Config
}

func NewService(cfg Config) Service {
	return &service{config: cfg}
}

func (s *service) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.config.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(s.config.Host, s.config.Port, s.config.Username, s.config.Password)

	return d.DialAndSend(m)
}
