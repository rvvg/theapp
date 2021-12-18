package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

// SendMail send mail
func SendMail(ip string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "blacklist@example.com")
	m.SetHeader("To", "test@examle.com")
	m.SetHeader("Subject", "New IP Blacklisted"+ip)
	m.SetBody("text/plain", "IP: "+ip)

	d := gomail.NewDialer("smtp.example.com", 587, "blacklist@example.com", "password")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
