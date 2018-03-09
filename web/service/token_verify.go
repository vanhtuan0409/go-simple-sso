package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vanhtuan0409/go-simple-sso/web/model"
)

type TokenVerifyService interface {
	Verify(token string) (*model.User, error)
}

type tokenVerifyService struct {
	ssoURL string
}

func NewTokenVerifyService(url string) TokenVerifyService {
	return &tokenVerifyService{
		ssoURL: url,
	}
}

type verifyResponse struct {
	Success bool        `json:"success"`
	User    *model.User `json:"user"`
	Message string      `json:"message"`
}

func (s *tokenVerifyService) Verify(token string) (*model.User, error) {
	verifyURL := s.ssoURL + "/verify_token"

	data, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", verifyURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonObj := new(verifyResponse)
	if err := json.NewDecoder(resp.Body).Decode(jsonObj); err != nil {
		return nil, err
	}
	if !jsonObj.Success {
		return nil, errors.New(jsonObj.Message)
	}

	return jsonObj.User, nil
}
