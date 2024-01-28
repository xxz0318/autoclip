// Package douyin_video 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 16:22
// @description:
package main

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"douyin_video/audio"
	"douyin_video/conf"
	"douyin_video/log"
	"douyin_video/novel"
	"douyin_video/utils"
	"github.com/golang-module/carbon/v2"
)

func main() {

	conf.LoadConfig()
	log.InitLog()

	ctx := context.Background()
	var novelObj novel.Novel
	var bookIds []int64
	if conf.C.NovelSource == "fanQie" {
		novelObj = novel.FanQie{}
		bookIds = conf.C.FanQie.CbidList
	} else if conf.C.NovelSource == "dianZhong" {
		novelObj = novel.DianZhong{}
	}
	for _, v := range bookIds {
		txtToAudio(ctx, novelObj, v)
	}

}

func txtToAudio(ctx context.Context, novelObj novel.Novel, bookId int64) {
	txt, err := novelObj.GetChapterContentByBookId(ctx, bookId)
	if err != nil {
		log.Errorf("[txtToAudio]GetChapterContentByBookId_error bookId:%d, err:%v", bookId, err)
		return
	}
	// 去除所有尖括号内的HTML代码，并换成换行符
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	txt = re.ReplaceAllString(txt, "\n")
	log.Debugf("[txtToAudio]txtLength:%+v, txt:%s", len(txt), txt)
	txtSlice := strings.Split(txt, "\n")
	var (
		length      int
		txtTmp      string
		newTxtSlice []string
	)
	for k, v := range txtSlice {
		txtSlice[k] = strings.TrimSpace(v)
		length += len(txtSlice[k])

		if length >= conf.C.Audio.TxtLength {

			newTxtSlice = append(newTxtSlice, txtTmp)
			txtTmp = txtSlice[k]
			length = 0
		} else if length < conf.C.Audio.TxtLength && k == len(txtSlice)-1 {
			newTxtSlice = append(newTxtSlice, txtTmp)
		} else {
			txtTmp = txtTmp + txtSlice[k]
		}
	}
	log.Debugf("[txtToAudio]newtxtLength:%+v", len(newTxtSlice))
	outputDir := conf.C.AudioOutputDir + carbon.Now().Format("Ymd") + "/%d/"
	outputDir = fmt.Sprintf(outputDir, bookId)
	err = utils.MkdirIfNotExist(outputDir)
	if err != nil {
		log.Errorf("[txtToAudio]mkdirIfNotExist_error bookId:%d, err:%v", bookId, err)
		return
	}
	for k, v := range newTxtSlice {

		log.Debugf("[txtToAudio]newTxtSlice======== K:%d, V:%s", k, v)

		fileName := fmt.Sprintf("%d_%d", bookId, k+1)

		err = audio.TxtToAudio(ctx, v, outputDir+fileName)
		if err != nil {
			log.Errorf("[txtToAudio]audio生成失败 bookId:%d, k:%d, err:%v", bookId, k, err)
			return
		}
		fmt.Printf("audio生成成功，fileName:%+v\n", fileName)
	}
}
