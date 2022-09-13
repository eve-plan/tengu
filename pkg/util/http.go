package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrHttpRespNot200 = errors.New("http response not 200")
)

// TODO log

func HttpGet(urlPath string, queryParams map[string]string, headers map[string]string) (respBody []byte, err error) {
	queryurl := ""
	// TODO 当params长度大于5之后使用strings.Builder{}
	for k, v := range queryParams {
		queryurl = queryurl + "&" + k + "=" + url.PathEscape(v)
	}
	if len(queryParams) > 0 {
		queryurl = "?" + queryurl[1:]
	}
	urlPath = urlPath + queryurl
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrHttpRespNot200
	}

	return io.ReadAll(resp.Body)
}

func HttpPost(urlPath string, queryParams, headers map[string]string, params interface{}) (respBody []byte, err error) {
	queryurl := ""
	// TODO 当params长度大于5之后使用strings.Builder{}
	if queryParams != nil {
		for k, v := range queryParams {
			queryurl = queryurl + "&" + k + "=" + url.PathEscape(v)
		}
	}
	if len(queryParams) > 0 {
		queryurl = "?" + queryurl[1:]
	}
	urlPath = urlPath + queryurl

	bodyBytes, _ := json.Marshal(params)

	// resp, err := http.Post(urlPath, "application/json", bytes.NewReader(bodyBytes))
	req, err := http.NewRequest("POST", queryurl, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, ErrHttpRespNot200
	}

	return io.ReadAll(resp.Body)
}
