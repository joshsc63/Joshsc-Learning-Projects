package main

import (
	"bytes"
	"html/template"

	"github.com/vanng822/go-premailer/premailer"
)

type Mail struct {
	Domain string
	Host string
	Port int
	Username string
	Password string
	Encryption string
	FromAddress string
	FromName string
}

type Message struct {
	From string
	FromName string
	To string
	Subject string
	Attachments []string //each entry is pathname to an attachment 
	Data any
	DataMap map[string]any
}

func (m *Mail) SendSMTPMessage(msg Message) error {
	if msg.From == "" {
		msg.From = m.FromAddress
	}
	
	if msg.FromName == "" {
		msg.FromName = m.FromName
	}
	
	// call golang templates for http & plain text mail then pass data to them
	data := map[string]any {
		"message": msg.Data,
	}
	
	msg.DataMap = data
	
	// get two forms of message
	formattedMessage, err := m.buildHTMLMessage(msg)  // html emailformat
}

func(m *Mail) buildHTMLMessage(msg Message) (string, error) {
	templateToRender := "./templates/mail.html.gohtml"
	
	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}
	
	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}
	
	formattedMessage := tpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}
	
	return formattedMessage, nil
}

func (m *Mail) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses: false,
	}
}