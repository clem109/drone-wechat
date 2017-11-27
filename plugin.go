package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	respFormat      = "Webhook %d\n  URL: %s\n  RESPONSE STATUS: %s\n  RESPONSE BODY: %s\n"
	debugRespFormat = "Webhook %d\n  URL: %s\n  METHOD: %s\n  HEADERS: %s\n  REQUEST BODY: %s\n  RESPONSE STATUS: %s\n  RESPONSE BODY: %s\n"
)

type (
	Repo struct {
		Owner string `json:"owner"`
		Name  string `json:"name"`
	}

	Build struct {
		Tag     string `json:"tag"`
		Event   string `json:"event"`
		Number  int    `json:"number"`
		Commit  string `json:"commit"`
		Ref     string `json:"ref"`
		Branch  string `json:"branch"`
		Author  string `json:"author"`
		Message string `json:"message"`
		Status  string `json:"status"`
		Link    string `json:"link"`
		Started int64  `json:"started"`
		Created int64  `json:"created"`
	}

	Config struct {
		Method      string
		AccessToken string
		Agentid     int    `json:"agentid"`
		MsgType     string `json:"msgtype"`
		URL         string
		MsgURL      string
		BtnTxt      string
		ToUser      string `json:"touser"`
		ToParty     string `json:"toparty"`
		Safe        bool
		ContentType string
		Debug       bool
		SkipVerify  bool
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	Job struct {
		Started int64 `json:"started"`
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {
	var buf bytes.Buffer
	var b []byte

	if p.Config.Title == "" {
		data := struct {
			Repo  Repo  `json:"repo"`
			Build Build `json:"build"`
		}{p.Repo, p.Build}

		if err := json.NewEncoder(&buf).Encode(&data); err != nil {
			fmt.Printf("Error: Failed to encode JSON payload. %s\n", err)
			return err
		}
		b = buf.Bytes()
	} else {
		textCard := struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			MsgURL      string `json:"url"`
			BtnTxt      string `json:"btntext"`
		}{p.Config.Title, p.Config.Description, p.Config.MsgURL, p.Config.BtnTxt}
		data := struct {
			ToUser   string `json:"touser"`
			MsgType  string `json:"msgtype"`
			Agentid  int    `json:"agentid"`
			TextCard struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				MsgURL      string `json:"url"`
				BtnTxt      string `json:"btntext"`
			} `json:"textcard"`
		}{p.Config.ToUser, p.Config.MsgType, p.Config.Agentid, textCard}

		b, _ = json.Marshal(data) // []byte(data)

	}

	// build and execute a request for each url.
	// all auth, headers, method, template (payload),
	// and content_type values will be applied to
	// every webhook request.

	url := p.Config.URL + p.Config.AccessToken
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}

	client := http.DefaultClient
	if p.Config.SkipVerify {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error: Failed to execute the HTTP request. %s\n", err)
		return err
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if p.Config.Debug || resp.StatusCode >= http.StatusBadRequest {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("Error: Failed to read the HTTP response body. %s\n", err)
		}

		if p.Config.Debug {
			fmt.Printf(
				debugRespFormat,
				req.URL,
				req.Method,
				req.Header,
				string(b),
				resp.Status,
				string(body),
			)
		} else {
			fmt.Printf(
				respFormat,
				req.URL,
				resp.Status,
				string(body),
			)
		}
	}
	return nil
}
