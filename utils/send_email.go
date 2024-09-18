package utils

import (
	"calender_reminders/config"
	"calender_reminders/global"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

func SendEmail(email, context string) error {
	// 从 Viper 中获取邮件配置信息
	emailConfig := getEmailConfig()
	fmt.Println(emailConfig)
	// 使用 go-mail 发送邮件
	m := gomail.NewMessage()
	m.SetHeader("From", emailConfig.User)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "context")
	m.SetBody("text/html", fmt.Sprintf("Your calender reminder: %s", context))

	d := gomail.NewDialer(emailConfig.Host, emailConfig.Port, emailConfig.User, emailConfig.Password)

	// 配置TLS
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 生产环境中，应该避免使用 InsecureSkipVerify 为 true

	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}

// 从 Viper 中获取邮件配置信息
func getEmailConfig() *config.Email {
	emailConfig := &config.Email{
		Host:     global.FAST_CONFIG.Email.Host,
		Port:     global.FAST_CONFIG.Email.Port,
		User:     global.FAST_CONFIG.Email.User,
		Password: global.FAST_CONFIG.Email.Password,
		UseSSL:   global.FAST_CONFIG.Email.UseSSL,
		UserTls:  global.FAST_CONFIG.Email.UserTls,
	}
	return emailConfig
}
