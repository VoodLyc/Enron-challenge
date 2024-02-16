package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type Bulk struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}

func main() {
	// CPU profiling
	// f, err := os.Create("cpu_profile.prof")
	// if err != nil {
	// 	log.Fatal("could not create CPU profile: ", err)
	// }
	// defer f.Close()

	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	log.Fatal("could not start CPU profile: ", err)
	// }
	// defer pprof.StopCPUProfile()

	rootDir := "./data/enron_mail_20110402/maildir"
	var emails []Email

	createIndex()

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

		emails = append(emails, email)
		fmt.Printf("append: %s\n", email.MessageID)
		if len(emails) >= 5000 {
			uploadDataToZincSearch(emails)
			emails = nil
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
	}

	uploadDataToZincSearch(emails)
}

func readEmailFile(filePath string) (Email, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Email{}, err
	}

	headerEnd := strings.Index(string(content), "\n\r")

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

func createIndex() {
	index, err := os.ReadFile("./index.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:4080/api/index", bytes.NewReader(index))
	if err != nil {
		fmt.Printf("Error creating the request: %v\n", err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}

func uploadDataToZincSearch(emails []Email) {
	emailData := Bulk{
		Index:   "emails",
		Records: emails,
	}

	jsonData, err := json.Marshal(emailData)
	if err != nil {
		fmt.Printf("Error encoding to JSON: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulkv2", bytes.NewReader(jsonData))
	if err != nil {
		fmt.Printf("Error creating the request: %v\n", err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}
