package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"shell_chat/config"
	"shell_chat/entity"
)

func Chat(messages entity.Messages) (*entity.OpenaiChatResponse, error) {
	// url
	url := config.Cfg.BaseUrl + "/v1/chat/completions"
	// data
	reqData := entity.OpenaiChatRequest{
		Model:    config.Cfg.GptModel,
		Messages: messages,
		Stream:   false,
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	// headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Cfg.ApiKey)

	// request
	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	// response
	resBody, err := io.ReadAll(response.Body)
	data := new(entity.OpenaiChatResponse)
	err = json.Unmarshal(resBody, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ChatStream(messages entity.Messages) (io.ReadCloser, error) {
	// url
	url := config.Cfg.BaseUrl + "/v1/chat/completions"
	// data
	reqData := entity.OpenaiChatRequest{
		Model:    config.Cfg.GptModel,
		Messages: messages,
		Stream:   true,
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	// headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Cfg.ApiKey)

	// request
	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// response
	return response.Body, nil
}
