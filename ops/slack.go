package ops

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendSlackMessage(url string, color string, title string, titleLink string, message string) error {
	attachment := Attachment{
		Text:      message,
		Color:     color,
		Title:     title,
		TitleLink: titleLink,
	}
	payload := Payload{
		Attachments: []Attachment{attachment},
	}

	jsonStr, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return err
}

type Attachment struct {
	Color     string `json:"color"`
	Text      string `json:"text"`
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
}

type Payload struct {
	Attachments []Attachment `json:"attachments,omitempty"`
}
