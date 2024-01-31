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
	"douyin_video/video"
	"douyin_video/youbao"
	"golang.org/x/sync/errgroup"
)

func main() {

	conf.LoadConfig()
	log.InitLog()

	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)
	if conf.C.ContentType == 1 || conf.C.ContentType == 3 {
		var novelObj novel.Novel
		var bookIds []int64
		if conf.C.NovelSource == "fanQie" {
			novelObj = novel.FanQie{}
			bookIds = conf.C.FanQie.CbidList
		} else if conf.C.NovelSource == "dianZhong" {
			novelObj = novel.DianZhong{}
		}
		group.Go(func() error {
			var err error
			for _, v := range bookIds {
				bookId := v
				err = txtToAudio(ctx, novelObj, bookId)
				if err != nil {
					break
				}
			}
			return err
		})

	}
	if conf.C.ContentType == 2 || conf.C.ContentType == 3 {
		for i := 0; i < conf.C.Video.VideoNum; i++ {
			group.Go(func() error {
				err := video.EditVideo(ctx, conf.C.Video.VideoTime,
					conf.C.Video.VideoType, conf.C.Video.VideoWidth,
					conf.C.Video.VideoHeight, conf.C.Video.Speed,
					conf.C.Video.FragDuration)
				return err
			})
		}
	}
	if conf.C.YouBao.IsAddKeyword {
		for bookId, v := range conf.C.YouBao.Keywords {
			v := v
			if len(v.KeyWord) < 1 {
				continue
			}
			bookId := bookId
			group.Go(func() error {
				err := youbao.ApplyFanQieKeyword(ctx, bookId, v.AuthorName, v.BookName, v.KeyWord)
				return err
			})
		}
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

func txtToAudio(ctx context.Context, novelObj novel.Novel, bookId int64) error {
	txt, err := novelObj.GetChapterContentByBookId(ctx, bookId)
	if err != nil {
		log.Errorf("[txtToAudio]GetContent_error bookId:%d, err:%v", bookId, err)
		return err
	}
	// 去除所有尖括号内的HTML代码，并换成换行符
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	txt = re.ReplaceAllString(txt, "\n")
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
	outputDir := conf.C.AudioOutputDir + "/%d/"
	outputDir = fmt.Sprintf(outputDir, bookId)
	err = utils.MkdirIfNotExist(outputDir)
	if err != nil {
		log.Errorf("[txtToAudio]mkdirIfNotExist_error bookId:%d, err:%v", bookId, err)
		return err
	}
	for k, v := range newTxtSlice {

		log.Debugf("[txtToAudio]newTxtSlice======== bookId:%d, k:%d, v:%s", bookId, k, v)

		fileName := fmt.Sprintf("%d_%d", bookId, k+1)
		err = audio.TxtToAudio(ctx, v, outputDir+fileName)
		if err != nil {
			log.Errorf("[txtToAudio]audio生成失败 bookId:%d, k:%d, err:%v", bookId, k, err)
			return fmt.Errorf("[txtToAudio]audio生成失败 bookId:%d, k:%d, err:%v", bookId, k, err)
		}
		log.Infof("audio生成成功，bookId:%d, fileName:%+v\n", bookId, fileName)
	}
	return nil
}
