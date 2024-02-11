package model

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

type Hit struct {
	Index     string  `json:"_index"`
	Type      string  `json:"_type"`
	ID        string  `json:"_id"`
	Score     float64 `json:"_score"`
	Timestamp string  `json:"@timestamp"`
	Source    Email   `json:"_source"`
	Highlight struct {
		Content []string `json:"content"`
	} `json:"highlight"`
}

type SearchBody struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}

type SearchResponse struct {
	Took   int     `json:"took"`
	Total  int     `json:"total"`
	Emails []Email `json:"emails"`
}
