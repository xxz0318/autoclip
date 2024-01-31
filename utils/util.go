// Package utils 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-28 20:54
// @description:
package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"douyin_video/log"
	"github.com/bytedance/sonic"
)

// MkdirIfNotExist makes directories if the input path is not exists
func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

func Request(ctx context.Context, url, method string, body interface{}, header map[string]string, rspStruct any) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	var req *http.Request
	if method == "POST" {
		bodyByte, _ := sonic.Marshal(&body)
		req, _ = http.NewRequestWithContext(ctx, method, url, bytes.NewReader(bodyByte))
	} else {
		req, _ = http.NewRequestWithContext(ctx, method, url, nil)
	}
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request rsp.StatusCode:%d", rsp.StatusCode)
	}
	rspBody, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	if rspStruct != nil {
		err = sonic.Unmarshal(rspBody, rspStruct)
		if err != nil {
			log.Errorf("[Request]Unmarshal_error err:%v", err)
			return "", err
		}
	}
	return string(rspBody), nil
}
