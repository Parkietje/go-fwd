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
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     string `json:"port"`
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
		fmt.Println("config error:", err)
	}

	return c
}

func main() {

	c := config()

	// Set up authentication information.
	myAuth := c.Auth
	auth := smtp.PlainAuth("", myAuth.Username, myAuth.Password, myAuth.Server)

	fmt.Println(myAuth.Port)
	//generate a JSON email and parse into Mail format
	str := `{"sender": "yannichiodi@gmail.com", "body": "Hello, world!"}`
	m := Mail{}
	json.Unmarshal([]byte(str), &m)

	//incoming mail
	from := m.Sender
	to := []string{myAuth.Username}
	msg := []byte(m.Body)
	fmt.Println("from: [" + from + "] to [" + to[0] + "] msg: [" + string(msg) + "]")

	// Connect to the server, authenticate, set the sender and recipient,
	err := smtp.SendMail(myAuth.Server+":"+myAuth.Port, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("succes")
}
