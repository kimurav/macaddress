package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type MacAddressClient struct {
	BaseUrl   *url.URL
	Client    *http.Client
	AuthToken string
}

type MacDetailResponse struct {
	VendorDetails     `json:"vendorDetails"`
	BlockDetails      `json:"blockDetails"`
	MacAddressDetails `json:"macAddressDetails"`
}
type VendorDetails struct {
	Oui            string `json:"oui"`
	IsPrivate      bool   `json:"isPrivate"`
	CompanyName    string `json:"companyName"`
	CompanyAddress string `json:"companyAddress"`
	CountryCode    string `json:"countryCode"`
}
type BlockDetails struct {
	BlockFound          bool   `json:"blockFound"`
	BorderLeft          string `json:"borderLeft"`
	BorderRight         string `json:"borderRight"`
	BlockSize           int    `json:"blockSize"`
	AssignmentBlockSize string `json:"assignmentBlockSize"`
	DateCreated         string `json:"dateCreated"`
	DateUpdated         string `json:"dateUpdated"`
}

type MacAddressDetails struct {
	SearchTerm         string   `json:"searchTerm"`
	IsValid            bool     `json:"isValid"`
	VirtualMachine     string   `json:"virtualMachine"`
	Applications       []string `json:"applications"`
	TransmissionType   string   `json:"transmissionType"`
	AdministrationType string   `json:"administrationType"`
	WiresharkNotes     string   `json:"wiresharkNotes"`
	Comment            string   `json:"comment"`
}

type ErrorResponse struct {
	Message string `json:"error"`
}

func NewMacAddressClient(baseUrl, authToken string) (*MacAddressClient, error) {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	return &MacAddressClient{
		BaseUrl:   parsedUrl,
		AuthToken: authToken,
		Client: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}

func (m MacAddressClient) Do(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-Authentication-Token", m.AuthToken)
	res, err := m.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
		}
		return errors.New(errRes.Message)
	}

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

func (m MacAddressClient) GetMacAddressDetails(macaddr string) (*MacDetailResponse, error) {
	requestURL := m.BaseUrl
	queryParams := requestURL.Query()
	queryParams.Add("search", macaddr)
	queryParams.Add("output", "json")
	requestURL.RawQuery = queryParams.Encode()

	req, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var resp MacDetailResponse
	if err := m.Do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
