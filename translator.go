package translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	IAM_TOKEN string
	folderId  string
}

type Response struct {
	Code         int
	Message      string
	Lang         string
	Translations []Translations `json:"translations"`
}

type Translations struct {
	Text string `json:"text"`
}

type Body struct {
	Source   string   `json:"sourceLanguageCode"`
	Target   string   `json:"targetLanguageCode"`
	Texts    []string `json:"texts"`
	FolderId string   `json:"folderId"`
}

func NewClient(IAM_TOKEN, folderId string) *Client {
	return &Client{
		IAM_TOKEN: IAM_TOKEN,
		folderId:  folderId}
}

func (c *Client) Translate(texts []string, source, target string) ([]string, error) { //TODO: Language Detection, more functionality
	b := Body{
		FolderId: c.folderId,
		Source:   source,
		Target:   target,
		Texts:    texts}

	jsonStr, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.IAM_TOKEN)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	var resp Response

	if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		fmt.Println(resp.Code, resp.Message, response.StatusCode)
		return nil, err
	}

	var result []string
	for _, v := range resp.Translations {

		result = append(result, v.Text)
	}
	return result, nil
}