package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

//Configuration contains the config from json.config
type Configuration struct {
	Auth Authorization `json:"authorization"`
}

//Authorization is a struct for authenticating with mail server
type Authorization struct {
	Username string
	Password string
	Server   string
}

// Mail is a mail
type Mail struct {
	Sender string `json:"sender"`
	Body   string `json:"body"`
}

func config() Configuration {
	c := Configuration{}

	file, _ := os.Open("../config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&c)
	if err != nil {
		fmt.Println("error:", err)
	}

	return c
}

func main() {

	c := config()

	// Set up authentication information.
	myAuth := c.Auth
	auth := smtp.PlainAuth("", myAuth.Username, myAuth.Password, myAuth.Server)

	//incoming mail
	str := `{"sender": "yannichiodi@gmail.com", "body": "Hello, world!"}`
	m := Mail{}
	json.Unmarshal([]byte(str), &m)

	from := m.Sender
	to := []string{myAuth.Username}
	msg := []byte(m.Body)

	// Connect to the server, authenticate, set the sender and recipient,
	err := smtp.SendMail(myAuth.Server, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("succes")
}
