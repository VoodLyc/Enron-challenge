package controller

import (
	"api/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there %s", "visitor")
}

func Search(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	data, err := json.Marshal(searchEmails(term))
	if err != nil {
		fmt.Println("Error encoding struct to JSON:", err)
		return
	}
	fmt.Fprint(w, string(data))
}

func searchEmails(term string) model.SearchResponse {
	query := `{
        "search_type": "match",
        "query":
        {
            "term": "` + term + `",
			"field": "_all"
        },
        "from": 0,
        "max_results": 20,
        "_source": []
    }`
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

	emails := getEmails(searchBody)

	return model.SearchResponse{
		Took:   searchBody.Took,
		Total:  searchBody.Hits.Total.Value,
		Emails: emails,
	}
}

func getEmails(searchBody model.SearchBody) []model.Email {

	emails := []model.Email{}

	for _, hit := range searchBody.Hits.Hits {
		emails = append(emails, hit.Source)
	}

	return emails
}
