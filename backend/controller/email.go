package controller

import (
	"api/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Search(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	from := r.URL.Query().Get("from")
	if from == "" {
		from = "0"
	}
	query := `{
        "search_type": "match",
        "query":
        {
            "term": "` + term + `",
			"field": "_all"
        },
        "from": ` + from + `,
        "max_results": 20,
        "_source": []
    }`
	makeSearchRequest(w, query)
}

func GetEmails(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format("2006-01-02T15:04:05.999Z")
	from := r.URL.Query().Get("from")
	if from == "" {
		from = "0"
	}
	query := `{
		"search_type": "daterange",
		"query": {
			"start_time": "2021-12-25T15:08:48.777Z",
			"end_time": "` + formattedTime + `"
		},
		"sort_fields": ["-@timestamp"],
		"from": ` + from + `,
		"max_results": 20,
		"_source": []
	}`
	makeSearchRequest(w, query)
}

func makeSearchRequest(w http.ResponseWriter, query string) {
	data, err := json.Marshal(searchEmails(query))
	if err != nil {
		fmt.Println("Error encoding struct to JSON:", err)
		return
	}
	fmt.Fprint(w, string(data))
}

func searchEmails(query string) model.SearchResponse {
	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", strings.NewReader(query))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var searchBody model.SearchBody
	err = json.Unmarshal(body, &searchBody)
	if err != nil {
		fmt.Println(err)
	}

	emails := getEmailsFromBody(searchBody)

	return model.SearchResponse{
		Took:   searchBody.Took,
		Total:  searchBody.Hits.Total.Value,
		Emails: emails,
	}
}

func getEmailsFromBody(searchBody model.SearchBody) []model.Email {

	emails := []model.Email{}

	for _, hit := range searchBody.Hits.Hits {
		emails = append(emails, hit.Source)
	}

	return emails
}
