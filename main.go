package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

func main() {
	http.HandleFunc("/", Mail)

	fmt.Println("starting server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//Configuration contain config values defined in config.json
type Configuration struct {
	Authentication Authentication `json:"authentication"`
}

//Authentication is a child of Configuration defined in config.json
type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     string `json:"port"`
}

//Email contains sender and message body received from client
type Email struct {
	Sender string `json:"sender"`
	Body   string `json:"body"`
}

//Mail is a httphandler which constructs Email from client POST and forwards to Authentication.server
func Mail(w http.ResponseWriter, r *http.Request) {
	c := config()

	if r.URL.Path != "/" {
		http.Error(w, "404 not found. ", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error paring POST: %v", err)
			return
		}
		m := Email{r.FormValue("email"), r.FormValue("message")}

		send(m, c.Authentication)

	default:
		fmt.Fprintf(w, "Please POST an html form with keys \"email\" and \"message\"")
	}
}

func send(m Email, a Authentication) {
	// Generate message
	from := m.Sender
	to := a.Username
	subj := "[go-fwd]"
	body := m.Body

	headers := make(map[string]string)
	headers["From: "] = from
	headers["To: "] = to
	headers["Subject: "] = subj

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Set up authentication information.
	auth := smtp.PlainAuth("", a.Username, a.Password, a.Server)

	/* cert, err := tls.LoadX509KeyPair("server-cert.pem", "server-key.pem")
	if err != nil {
		log.Panic(err)
	} */

	tlsconfig := &tls.Config{
		//Certificates: []tls.Certificate{cert},
		ServerName: a.Server,
	}

	c, err := smtp.Dial(a.Server + ":" + a.Port)
	if err != nil {
		log.Panic(err)
	}
	c.StartTLS(tlsconfig)

	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = c.Mail(from); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to); err != nil {
		log.Panic(err)
	}

	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}

func config() Configuration {
	c := Configuration{}

	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&c)
	if err != nil {
		fmt.Println("config error:", err)
	}

	return c
}
