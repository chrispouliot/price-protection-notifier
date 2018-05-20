package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moxuz/price-protection-notifier/config"
)

type ReqBody struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

func SendPriceChange(address, checkUrl string, price float64) error {
	return send(config.MailFromAddress, address, checkUrl, nil, price)
}

func SendError(address, checkUrl string, err error) error {
	return send(config.MailFromAddress, address, checkUrl, err, 0)
}

func send(fromAddress, address, checkUrl string, err error, price float64) error {
	text := fmt.Sprintf("Check %s has", checkUrl)
	if err != nil {
		text = fmt.Sprintf("%s error %s", text, err.Error())
	} else {
		text = fmt.Sprintf("%s price change to $%.2f", text, price)
	}
	body := ReqBody{
		To:      address,
		From:    fromAddress,
		Subject: "Price Notifier",
		Text:    text,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", checkUrl, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.SetBasicAuth("api", config.MailAPIKey)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("status code %d: %s while attempting to send email", resp.StatusCode, resp.Status)
	}

	return nil
}
