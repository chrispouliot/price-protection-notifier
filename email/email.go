package email

import (
	"bytes"
	"fmt"
	"net/http"

	"mime/multipart"

	"github.com/moxuz/price-protection-notifier/config"
)

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

	body, boundary := getPostFormBody(fromAddress, address, text, "Price Notifier Alert")
	req, err := http.NewRequest("POST", config.MailApiUrl, body)
	req.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", boundary))
	req.SetBasicAuth("api", config.MailAPIKey)

	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		return fmt.Errorf("status code %d: %s (%s) while attempting to send email", resp.StatusCode, resp.Status, buf.String())
	}

	return nil
}

func getPostFormBody(fromAddress, address, text, subject string) (*bytes.Buffer, string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("from", fromAddress)
	writer.WriteField("to", address)
	writer.WriteField("text", text)
	writer.WriteField("subject", subject)
	boundary := writer.Boundary()
	writer.Close()
	return body, boundary
}
