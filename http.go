package go_example

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func httpGet() string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		panic(fmt.Sprintf("http get error: %v", err))
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func httpPost() string {
	client := &http.Client{Timeout: 5 * time.Second}
	data, _ := json.Marshal([]string{"hello"})
	resp, err := client.Post("http://www.baidu.com", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(fmt.Sprintf("http post error: %v", err))
	}

	defer resp.Body.Close()

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func httpPost2() string {
	client := &http.Client{Timeout: 5 * time.Second}
	data, _ := json.Marshal([]string{"hello"})
	req, err := http.NewRequest("POST", "http://www.baidu.com", bytes.NewBuffer(data))
	if err != nil {
		panic(fmt.Sprintf("http post error: %v", err))
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func httpPost3() string {
	client := &http.Client{Timeout: 5 * time.Second}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("test", "/tmp/test.png")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	file, err := os.Open("/tmp/test.png")
	defer file.Close()
	_, err = io.Copy(part, file)

	params := make(map[string]string)
	params["test"] = "hello"
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	req, err := http.NewRequest("POST", "http://www.baidu.com", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)

	defer resp.Body.Close()

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}
