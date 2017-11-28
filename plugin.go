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
		CorpID      string
		CorpSecret  string
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

	Response struct {
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	Job struct {
		Started int64 `json:"started"`
	}

	Plugin struct {
		Repo     Repo
		Build    Build
		Config   Config
		Response Response
		Job      Job
	}
)

func getToken(body []byte) (*Response, error) {
	var s = new(Response)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func (p Plugin) Exec() error {

	var buf bytes.Buffer
	var b []byte

	// GET request to get the access token
	accessURL := p.Config.URL + p.Config.CorpID + "&corpsecret=" + p.Config.CorpSecret
	fmt.Println("URL:>", accessURL)

	req, err := http.NewRequest("GET", accessURL, bytes.NewBuffer(b))
	var client = http.DefaultClient
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

	s, err := getToken(body)

	// POST Request to WeChat work
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

	// POST URL
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + s.AccessToken

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	request.Header.Set("Content-Type", "application/json")

	client = http.DefaultClient
	if p.Config.SkipVerify {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("Error: Failed to execute the HTTP request. %s\n", err)
		return err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	responseBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(responseBody))

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
