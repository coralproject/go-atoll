package atoll

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseURL string
}

type Response struct {
	Results  interface{}
	Response *http.Response
}

type Payload struct {
	Data interface{} `json:"data"`
}

func (c *Client) Post(data interface{}, endpoint string) (Response, error) {
	resultsResp := Response{}
	payload := Payload{
		Data: data,
	}

	jsonStr, err := json.Marshal(payload)
	if err != nil {
		return resultsResp, err
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return resultsResp, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resultsResp, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resultsResp, err
	}

	var results map[string]interface{}
	if err := json.Unmarshal(body, &results); err != nil {
		return resultsResp, err
	}
	resultsResp.Results = results["results"]
	resultsResp.Response = resp
	return resultsResp, nil
}
