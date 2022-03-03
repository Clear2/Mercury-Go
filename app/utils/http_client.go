package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPGet get 请求
func HTTPGet(uri string, header map[string]string) ([]byte, error) {
	return HTTPGetContext(context.Background(), uri, header)
}
func HTTPGetContext(ctx context.Context, uri string, header map[string]string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// HTTPPost post 请求
func HTTPPost(uri string, data string, header map[string]string) ([]byte, error) {
	return HTTPPostContext(context.Background(), uri, data, header)
}

// HTTPPostContext post 请求
func HTTPPostContext(ctx context.Context, uri string, data string, header map[string]string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(data))
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, body)
	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// PostJSON post json 数据请求
func PostJSON(uri string, obj interface{}) ([]byte, error) {
	jsonBuf := new(bytes.Buffer)
	enc := json.NewEncoder(jsonBuf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(obj)
	if err != nil {
		return nil, err
	}
	response, err := http.Post(uri, "application/json;charset=utf-8", jsonBuf)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
