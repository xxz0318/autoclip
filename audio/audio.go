// Package audio 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 21:40
// @description:
package audio

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"douyin_video/conf"
	"douyin_video/log"
	"github.com/bytedance/sonic"
	"github.com/forgoer/openssl"
)

type audioConvertReq struct {
	Pitch       string `json:"pitch"`
	Speed       int32  `json:"speed"`
	StyleDegree int32  `json:"styleDegree"`
	Text        string `json:"text"`
	Voice       string `json:"voice"`
	Sign        string `json:"sign"`
	Version     string `json:"version"`
}
type audioRsp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// TxtToAudio 文本转语音 最多8000字
func TxtToAudio(ctx context.Context, content, audioFileName string) error {
	// ==================提交配音任务==================
	taskParams := audioConvertReq{
		Pitch:       "0",
		Speed:       5,
		StyleDegree: 0,
		Text:        content,
		Voice:       conf.C.Audio.VoiceId,
		Sign:        conf.C.Audio.Sign,
		Version:     "28.0",
	}
	taskBody, err := sonic.Marshal(&taskParams)
	if err != nil {
		log.SugarLogger.Errorf("[TxtToAudio]Marshal_error err:%v", err)
		return err
	}
	client := &http.Client{}
	taskReq, _ := http.NewRequest("POST", conf.C.Audio.ConvertUrl, bytes.NewReader(taskBody))
	setAudioHeader(ctx, taskReq, false)
	taskRsp, taskRspErr := client.Do(taskReq)
	if taskRspErr != nil {
		log.SugarLogger.Errorf("[TxtToAudio]TaskRequest_error taskRspErr:%v", taskRspErr)
		return taskRspErr
	}
	defer taskRsp.Body.Close()
	if taskRsp.StatusCode != http.StatusOK {
		log.SugarLogger.Errorf("[TxtToAudio]TaskRsp_error rsp.StatusCode:%d", taskRsp.StatusCode)
		return fmt.Errorf("TaskRsp_error rsp.StatusCode:%d", taskRsp.StatusCode)
	}
	var audioTaskRsp audioRsp
	taskRspBody, _ := io.ReadAll(taskRsp.Body)

	log.SugarLogger.Infof("[TxtToAudio] taskRspBody:%s", string(taskRspBody))

	err = sonic.Unmarshal(taskRspBody, &audioTaskRsp)
	if err != nil {
		log.SugarLogger.Errorf("[TxtToAudio]Unmarshal_error err:%v", err)
		return err
	}

	// ==================获取音频地址==================

	taskId := audioTaskRsp.Data
	// taskId := "051f3a10-dbfc-42f2-8781-6adf77bc3f20"
	fmt.Printf("taskId:%s\n", taskId)

	urlParams := struct {
		TaskId string `json:"taskId"`
	}{TaskId: taskId}
	urlBody, err := sonic.Marshal(&urlParams)
	urlReq, _ := http.NewRequest("POST", conf.C.Audio.GetVoiceAudioUrlWeb, bytes.NewReader(urlBody))
	setAudioHeader(ctx, urlReq, false)
	var audioUrlRsp audioRsp

	for {

		urlRsp, urlRspErr := client.Do(urlReq)
		if urlRspErr != nil {
			log.SugarLogger.Errorf("[TxtToAudio]UrlRequest_error urlRspErr:%v", urlRspErr)
			return urlRspErr
		}
		defer urlRsp.Body.Close()
		if urlRsp.StatusCode != http.StatusOK {
			log.SugarLogger.Errorf("[TxtToAudio]UrlRsp_status_error rsp.StatusCode:%d", urlRsp.StatusCode)
			return fmt.Errorf("UrlRsp_status_error rsp.StatusCode:%d", urlRsp.StatusCode)
		}
		urlRspBody, _ := io.ReadAll(urlRsp.Body)

		log.SugarLogger.Infof("[TxtToAudio]urlRspBody:%s", string(urlRspBody))

		err = sonic.Unmarshal(urlRspBody, &audioUrlRsp)
		if err != nil {
			log.SugarLogger.Errorf("[TxtToAudio]Unmarshal_error err:%v", err)
			return err
		}
		if audioUrlRsp.Msg == "配音生成中" || audioUrlRsp.Msg != "success" {
			time.Sleep(time.Second * 10) // 停10秒后重试
			continue
		}
		if audioUrlRsp.Msg == "success" {
			break
		}
	}
	// ==================下载音频==================
	urlCrypt := audioUrlRsp.Data
	videoUrl, err := aesDecrypt(urlCrypt, conf.C.Audio.AesKey)
	if err != nil {
		log.SugarLogger.Errorf("[TxtToAudio]aesDecrypt_error err:%v", err)
		return err
	}

	log.SugarLogger.Infof("videoUrl:%s", videoUrl)

	dwReq, _ := http.NewRequest("GET", videoUrl, nil)
	setAudioHeader(ctx, dwReq, true)
	dwRsp, dwRspErr := client.Do(dwReq)
	if dwRspErr != nil {
		log.SugarLogger.Errorf("[TxtToAudio]DwRequest_error dwRspErr:%v", dwRspErr)
		return dwRspErr
	}
	defer dwRsp.Body.Close()
	// 创建一个文件用于保存
	out, err := os.Create(audioFileName)
	if err != nil {
		log.SugarLogger.Errorf("[TxtToAudio]CreateAudioFile_error err:%v", err)
		return err
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, dwRsp.Body)
	if err != nil {
		log.SugarLogger.Errorf("[TxtToAudio]Copy_error err:%v", err)
		return err
	}
	return nil
}
func setAudioHeader(ctx context.Context, req *http.Request, isUrl bool) {
	header := conf.C.Audio.Header
	for k, v := range header {
		req.Header.Set(k, v)
	}
	if isUrl {
		req.Header.Set("range", "bytes=0-")
	}
}

// aesDecrypt aes解密
func aesDecrypt(data, key string) (string, error) {
	src, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		src, err = base64.StdEncoding.DecodeString(data + "=")
		if err != nil {
			return "", err
		}
	}
	dst, err := openssl.AesECBDecrypt(src, []byte(key), openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}
