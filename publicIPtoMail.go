package main

import (
	"log"
	"runtime"
	"time"

	"github.com/polera/publicip"
	"github.com/xhit/go-simple-mail"
)

// Check actual IP every 30 minutes and if it changes, then it sends an e-mail
func checkIP() {

	var actualIP = "0.0.0.0"
	for {
		myIpAddr, err := publicip.GetIP()
		if err != nil {
			log.Printf("Error getting IP address: %s\n", err)
		} else {
			if actualIP == myIpAddr {
				log.Printf("IP is the same: %s", actualIP)
			} else {

				actualIP = myIpAddr
				log.Printf("IP changed: %s", actualIP)

				server := mail.NewSMTPClient()

				server.Host = "mail.host.com"
				server.Port = 465
				server.Username = "username@host.com"
				server.Password = "password"
				server.Encryption = mail.EncryptionSSL
				server.KeepAlive = false
				server.ConnectTimeout = 10 * time.Second
				server.SendTimeout = 10 * time.Second

				smtpClient, err := server.Connect()

				if err != nil {
					log.Printf("Couldn't connect to smtp host!")
				}

				email := mail.NewMSG()
				email.SetFrom("User Name <user@host.com>")
				email.AddTo("destination@mail.com")
				email.SetSubject("Public IP changed")
				email.SetBody(mail.TextPlain, myIpAddr)

				err = email.Send(smtpClient)

				if err != nil {
					log.Printf("Couldn't send the e-mail: %s\n", err)
				} else {
					log.Printf("Email sent...")
				}
			}
		}
		time.Sleep(30 * time.Minute)
	}
}

func main() {
	go checkIP()
	runtime.Goexit()
}
