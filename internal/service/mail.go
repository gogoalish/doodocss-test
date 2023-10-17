package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/smtp"
	"os"
	"strings"

	"github.com/gogoalish/doodocs-test/utils"
)

type MailService interface {
	SendMessage(m *MailParams) error
}

type Mail struct{}

type MailParams struct {
	To        []string
	CC        []string
	BCC       []string
	Subject   string
	Body      string
	FilePaths []string
	FileNames []string
}

func (*Mail) SendMessage(params *MailParams) error {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("Subject: %s\n", params.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(params.To, ",")))
	if len(params.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(params.CC, ",")))
	}
	if len(params.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(params.BCC, ",")))
	}
	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if len(params.FilePaths) > 0 {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	}

	buf.WriteString(params.Body)
	if len(params.FilePaths) > 0 {
		for i, path := range params.FilePaths {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", utils.DetectMimeType(params.FileNames[i])))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", params.FileNames[i]))
			fileData, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			b := make([]byte, base64.StdEncoding.EncodedLen(len(fileData)))
			base64.StdEncoding.Encode(b, fileData)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}
	var (
		host       = os.Getenv("EMAIL_HOST")
		login      = os.Getenv("EMAIL_LOGIN")
		password   = os.Getenv("EMAIL_PASSWORD")
		portNumber = os.Getenv("EMAIL_PORT")
	)
	fmt.Printf("host: %v\n", host)
	auth := smtp.PlainAuth("", login, password, host)
	return smtp.SendMail(fmt.Sprintf("%s:%s", host, portNumber), auth, login, params.To, buf.Bytes())
}
