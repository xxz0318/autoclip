// Package douyin_video 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 16:22
// @description:
package main

import (
	"context"
	"fmt"
	"regexp"

	"douyin_video/audio"
	"douyin_video/conf"
	"douyin_video/log"
	"douyin_video/novel"
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
		log.SugarLogger.Errorf("[txtToAudio]GetChapterContentByBookId_error bookId:%d, err:%v", bookId, err)
		return
	}
	// 去除所有尖括号内的HTML代码，并换成换行符
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	txt = re.ReplaceAllString(txt, "\n")
	fileName := fmt.Sprintf("./%d.wav", bookId)
	err = audio.TxtToAudio(ctx, txt, fileName)
	if err != nil {
		log.SugarLogger.Errorf("[txtToAudio]audio生成失败 bookId:%d, err:%v", bookId, err)
		return
	}
	fmt.Printf("audio生成成功，fileName:%+v\n", fileName)
}
