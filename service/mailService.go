package service

import (
	"log"
	"romanm/web-service-gin/config"

	"github.com/sendgrid/sendgrid-go"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailService interface {
	CreateMail(mailReq *Mail) []byte
	SendMail(mailReq *Mail) error
	NewMail(from string, to []string, subject string, mailType MailType, data *MailData) *Mail
}

type MailType int

// List of Mail Types we are going to send.
const (
	MailConfirmation MailType = iota + 1
	PassReset
)

// MailData represents the data to be sent to the template of the mail.
type MailData struct {
	Username string
	Link     string
}

// Mail represents a email request
type Mail struct {
	from    string
	to      string
	subject string
	mtype   MailType
	data    *MailData
}

// SGMailService is the sendgrid implementation of our MailService.
type SGMailService struct {
	configs *config.Config
}

func NewSGMailService(configs *config.Config) *SGMailService {
	return &SGMailService{configs}
}

// CreateMail takes in a mail request and constructs a sendgrid mail type.
func (ms *SGMailService) CreateMail(mailReq *Mail) []byte {

	from := mail.NewEmail("Weather Updates Service", mailReq.from)

	m := mail.NewV3Mail()

	m.SetTemplateID(ms.configs.MailVerifTemplateID)

	m.SetFrom(from)

	p := mail.NewPersonalization()

	p.AddTos(mail.NewEmail(mailReq.data.Username, mailReq.to))

	p.SetDynamicTemplateData("Username", mailReq.data.Username)
	p.SetDynamicTemplateData("Link", mailReq.data.Link)

	m.AddPersonalizations(p)
	return mail.GetRequestBody(m)
}

// SendMail creates a sendgrid mail from the given mail request and sends it.
func (ms *SGMailService) SendMail(mailReq *Mail) error {

	log.Println("sending mail to: ", mailReq.to)
	log.Println("mail subject: ", mailReq.subject)
	log.Println("mail type: ", mailReq.mtype)
	log.Println("mail data: ", mailReq.data)
	log.Println("mail from: ", mailReq.from)
	log.Println("mail template id: ", ms.configs.MailVerifTemplateID)

	request := sendgrid.GetRequest(ms.configs.SendGridApiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = ms.CreateMail(mailReq)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		log.Printf("unable to send mail, error: %v", err)
		return err
	}

	if response.StatusCode != 202 {
		log.Printf("mail not sent, status code: %v", response.StatusCode)
		log.Printf("mail not sent, response body: %v", response.Body)
	} else {
		log.Printf("mail sent successfully, status code: %v", response.StatusCode)
	}
	return nil
}

// NewMail returns a new mail request.
func (ms *SGMailService) NewMail(from string, to string, subject string, mailType MailType, data *MailData) *Mail {
	return &Mail{
		from:    from,
		to:      to,
		subject: subject,
		mtype:   mailType,
		data:    data,
	}
}
