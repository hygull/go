package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func main() {
	send("Welcom to jfashion...created by Jeevitha")
}

func send(body string) {
	from := "rishikesh0014051992@gmail.com"
	fmt.Println(os.Args[0], os.Args[1], len(os.Args))
	if len(os.Args) < 2 {
		fmt.Println("Your email password is required as cmd args")
		return
	}
	pass := os.Args[1]
	to := "rishikesh0011115067@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: JFASHION\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	//log.Print("sent, visit http://foobarbazz.mailinator.com")
	fmt.Println("Email successfuly sent")
}
