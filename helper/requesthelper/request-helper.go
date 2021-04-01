package requesthelper

import (
	"net/http"
	"time"

	"github.com/onkarsutar/BankAccount/helper/loghelper"
)

func NewRequest(url, reqType, contentType, token string, data []byte, ClientRequestTimeOut int) (*http.Response, error) {

	req, err := http.NewRequest(reqType, url, nil)
	if err != nil {
		loghelper.LogError("NewRequest Error: ", err)
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	// req.Header.Add("authorization", token)

	client := &http.Client{
		Timeout: time.Second * time.Duration(ClientRequestTimeOut),
	}

	resp, err := client.Do(req)
	if err != nil {
		loghelper.LogError("NewRequest Error: ", err, req)
		return nil, err
	}
	return resp, nil
}
