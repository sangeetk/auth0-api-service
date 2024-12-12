package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"auth0-api-service/config"
	"auth0-api-service/models"
)

type AuthService struct {
	config *config.Config
	client *http.Client
}

func NewAuthService(config *config.Config) *AuthService {
	return &AuthService{
		config: config,
		client: &http.Client{},
	}
}

func (s *AuthService) Signup(req *models.SignupRequest) error {
	auth0Payload := map[string]interface{}{
		"email":      req.Email,
		"password":   req.Password,
		"connection": "Username-Password-Authentication",
	}

	jsonPayload, err := json.Marshal(auth0Payload)
	if err != nil {
		return err
	}

	auth0Req, err := http.NewRequest(
		"POST",
		"https://"+s.config.Auth0Domain+"/dbconnections/signup",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return err
	}

	auth0Req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(auth0Req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to create user")
	}

	return nil
}

func (s *AuthService) Signin(req *models.SigninRequest) (*models.Auth0Response, error) {
	auth0Payload := map[string]interface{}{
		"grant_type":    "password",
		"username":      req.Email,
		"password":      req.Password,
		"client_id":     s.config.Auth0ClientID,
		"client_secret": s.config.Auth0ClientSecret,
		"audience":      s.config.Auth0Audience,
		"scope":         "openid profile email",
	}

	jsonPayload, err := json.Marshal(auth0Payload)
	if err != nil {
		return nil, err
	}

	auth0Req, err := http.NewRequest(
		"POST",
		"https://"+s.config.Auth0Domain+"/oauth/token",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return nil, err
	}

	auth0Req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(auth0Req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("authentication failed")
	}

	var auth0Response models.Auth0Response
	if err := json.NewDecoder(resp.Body).Decode(&auth0Response); err != nil {
		return nil, err
	}

	return &auth0Response, nil
}
