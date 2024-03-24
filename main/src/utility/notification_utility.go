package utility

import (
	"appointment-notification-sender/main/src/config"
	"appointment-notification-sender/main/src/models"
	"context"
	"fmt"
	brevo "github.com/getbrevo/brevo-go/lib"
	"gopkg.in/gomail.v2"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func SendMessages(customers *[]models.Customer) {
	for i := range *customers {
		if (*customers)[i].IsSMS {
			//err := sendSMS((*customers)[i].CellNumber, "Your appointment has been confirmed.")
			//if err != nil {
			//	(*customers)[i].IsSMSSent = false
			//} else {
			//	(*customers)[i].IsSMSSent = true
			//}
			log.Println("Send SMS is disabled")
			(*customers)[i].IsSMSSent = false
		}
		if (*customers)[i].IsEmail {
			sendEmail((*customers)[i].Email, (*customers)[i].FullName)
			(*customers)[i].IsEmailSent = true
		}
	}
}

func sendEmail(email, fullName string) {
	log.Println("Sending Email to : " + email)
	log.Println("From Email: " + config.GetAppConfig().EmailConfig.FromEmailAddress)
	from := config.GetAppConfig().EmailConfig.FromEmailAddress
	password := config.GetAppConfig().EmailConfig.FromEmailPassword
	to := email

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Appointment Confirmation")

	tpl, err := template.ParseFiles("view/email_template.html")
	if err != nil {
		log.Println(err.Error())
	}

	data := struct {
		FullName string
	}{
		FullName: fullName,
	}

	var emailBody strings.Builder
	if err := tpl.Execute(&emailBody, data); err != nil {
		log.Println("Error executing template:", err)
		return
	}
	log.Print(emailBody.String())
	m.SetBody("text/html", emailBody.String())
	d := gomail.NewDialer(config.GetAppConfig().EmailConfig.SmtpRelayAddress, config.GetAppConfig().EmailConfig.SmtpPort, from, password)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err.Error())
	}
	log.Println("Email Sent Successfully")
}

func sendSMS(toNumber, message string) error {
	log.Println("Sending SMS to: " + toNumber)
	var ctx context.Context
	cfg := brevo.NewConfiguration()
	log.Println("API Key: " + config.GetAppConfig().EmailConfig.SmsApiKey)
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", config.GetAppConfig().EmailConfig.SmsApiKey)

	br := brevo.NewAPIClient(cfg)
	sms, h, err := br.TransactionalSMSApi.SendTransacSms(ctx, brevo.SendTransacSms{
		Sender:             "Danish",
		Recipient:          toNumber,
		Content:            message,
		Type_:              "Marketing",
		Tag:                "Appointment",
		WebUrl:             "",
		UnicodeEnabled:     false,
		OrganisationPrefix: "WebDel",
	})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(h.Body)

	// Read response body
	body, err := io.ReadAll(h.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	// Print response body as string
	fmt.Println(string(body))

	if h.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send SMS: status code %d", h.StatusCode)
	}

	log.Println(sms)
	log.Println("Sms Sent Successfully")

	return nil
}
