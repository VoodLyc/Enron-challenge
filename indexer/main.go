package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Email struct {
	MessageID               string `json:"message_id"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Subject                 string `json:"subject"`
	MimeVersion             string `json:"mime_version"`
	ContentType             string `json:"content_type"`
	ContentTransferEncoding string `json:"content_transfer_encoding"`
	XFrom                   string `json:"x_from"`
	XTo                     string `json:"x_to"`
	Xcc                     string `json:"x_cc"`
	Xbcc                    string `json:"x_bcc"`
	XFolder                 string `json:"x_folder"`
	XOrigin                 string `json:"x_origin"`
	XFilename               string `json:"x_filename"`
	Content                 string `json:"content"`
}

func main() {
	rootDir := "./data/enron_mail_20110402/maildir"
	outputDir := "./output"

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		email, err := readEmailFile(path)
		if err != nil {
			fmt.Printf("Error reading email file %s: %v\n", path, err)
			return nil
		}

		jsonData, err := json.Marshal(email)
		if err != nil {
			fmt.Printf("Error encoding email to JSON: %v\n", err)
			return nil
		}

		outputFilePath := filepath.Join(outputDir, fmt.Sprintf("%s.json", email.MessageID[1:len(email.MessageID)-1]))

		if err := os.WriteFile(outputFilePath, jsonData, os.ModePerm); err != nil {
			fmt.Printf("Error writing JSON to file %s: %v\n", outputFilePath, err)
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
	}
}

func readEmailFile(filePath string) (Email, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Email{}, err
	}

	headerEnd := strings.Index(string(content), "\n\r")
	fmt.Println(filePath)
	if headerEnd == -1 {
		return Email{}, fmt.Errorf("invalid email: %s", filePath)
	}

	headers := strings.Split(string(content[:headerEnd]), "\n")

	body := string(content[headerEnd+2:])

	email := Email{}
	for _, line := range headers {
		parseHeaders(line, &email)
	}

	email.Content = body

	return email, nil
}

func parseHeaders(line string, email *Email) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return
	}

	headerName := strings.TrimSpace(parts[0])
	headerValue := strings.TrimSpace(parts[1])

	switch headerName {
	case "Message-ID":
		email.MessageID = headerValue
	case "Date":
		email.Date = headerValue
	case "From":
		email.From = headerValue
	case "To":
		email.To = headerValue
	case "Subject":
		email.Subject = headerValue
	case "Mime-Version":
		email.MimeVersion = headerValue
	case "Content-Type":
		email.ContentType = headerValue
	case "Content-Transfer-Encoding":
		email.ContentTransferEncoding = headerValue
	case "X-From":
		email.XFrom = headerValue
	case "X-To":
		email.XTo = headerValue
	case "X-cc":
		email.Xcc = headerValue
	case "X-bcc":
		email.Xbcc = headerValue
	case "X-Folder":
		email.XFolder = headerValue
	case "X-Origin":
		email.XOrigin = headerValue
	case "X-FileName":
		email.XFilename = headerValue
	}
}
