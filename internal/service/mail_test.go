package service

import (
	"os"
	"testing"

	"github.com/gogoalish/doodocs-test/utils"
	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	utils.ParseEnv("../../config/credentials.env")

	f, err := os.CreateTemp(".", "*.tmp")
	defer os.Remove(f.Name())
	require.NoError(t, err)
	mail := &Mail{}
	params := &MailParams{
		To:        []string{"aalisherh@gmail.com"},
		CC:        []string{"aalisherh@gmail.com"},
		BCC:       []string{"aalisherh@gmail.com"},
		Subject:   "Test Subject",
		Body:      "Test Body",
		FilePaths: []string{f.Name()},
		FileNames: []string{f.Name()},
	}
	err = mail.SendMessage(params)
	require.NoError(t, err)
}
