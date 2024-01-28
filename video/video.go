// Package video 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-28 19:31
// @description:
package video

import (
	"context"
	"math/rand"
	"os"
	"path/filepath"
	"slices"
	"time"

	"douyin_video/conf"
	"douyin_video/log"
	"douyin_video/utils"
	"douyin_video/video/moviego"
	"github.com/golang-module/carbon/v2"
)

func EditVideo(ctx context.Context, videoTime float64, videoType string, width, length int64, speed, fragDuration float64) error {
	// 遍历目录中所有视频
	videoFiles, _ := getDirAllVideo(conf.C.VideoResourceDir + videoType)
	// 却输出目录及文件名
	outputDir := conf.C.VideoOutputDir + carbon.Now().ToShortDateString() + "/"
	err := utils.MkdirIfNotExist(outputDir)
	if err != nil {
		outputDir = conf.C.VideoOutputDir
		log.Errorf("[EditVideo]MkdirIfNotExist_error err:%v", err)
	}
	fileName := outputDir + carbon.Now().ToShortDateTimeString() + ".mp4"

	firstIndex := getRandomInt(0, int64(len(videoFiles)))
	firstVideo := videoFiles[firstIndex]
	log.Debugf("[EditVideo]firstVideo:%s", firstVideo)
	clip, err := moviego.Load(firstVideo)
	if err != nil {
		log.Errorf("[EditVideo]Load_error video:%s,err:%v", firstVideo, err)
		return err
	}
	var addedVideo []int64
	addedVideo = append(addedVideo, firstIndex)
	if fragDuration != 0 {
		clip = subClip(clip, fragDuration*speed)
		addedVideo = append(addedVideo, firstIndex)
	}
	var videos []moviego.Video
	var allDuration float64
	videos = append(videos, clip)
	allDuration += clip.Duration()
	log.Debugf("[EditVideo]first clip.Duration():%v", clip.Duration())
	for {
		if allDuration >= videoTime {
			break
		}
		// 随机取视频
		index := getRandomInt(0, int64(len(videoFiles)))
		for {
			if slices.Contains(addedVideo, index) {
				index = getRandomInt(0, int64(len(videoFiles)))
			} else {
				break
			}
		}
		video := videoFiles[index]
		log.Debugf("[EditVideo]next_video:%s", video)
		// 随机取视频片段
		clipTmp, err := moviego.Load(video)
		if err != nil {
			log.Errorf("[EditVideo]Load_error video:%s,err:%v", video, err)
			return err
		}
		if fragDuration != 0 {
			clipTmp = subClip(clipTmp, fragDuration*speed)
		}
		videos = append(videos, clipTmp)
		addedVideo = append(addedVideo, index)
		allDuration += clipTmp.Duration()
		log.Debugf("[EditVideo]allDuration:%v, new clip.Duration:%v", allDuration, clipTmp.Duration())
		// // 拼接视频
		// log.Debugf("[EditVideo]all_before clip.Duration:%v, new clip.Duration:%v", clip.Duration(), clipTmp.Duration())
		// clip, err = moviego.Concat([]moviego.Video{clip, clipTmp})
		// if err == nil {
		// 	addedVideo = append(addedVideo, index)
		// }
		// log.Debugf("[EditVideo]all_end clip.Duration:%v, new clip.Duration:%v", clip.Duration(), clipTmp.Duration())
	}

	// 拼接视频
	clip, err = moviego.Concat(videos)
	if err != nil {
		log.Errorf("[EditVideo]Concat_error err:%v", err)
	}
	log.Debugf("[EditVideo]all clip.Duration:%v", clip.Duration())
	if clip.Duration() > videoTime {
		clip = clip.SubClip(0, videoTime)
	}

	err = clip.Resize(width, length).DelAudio().Output(fileName).Run()
	if err != nil {
		log.Errorf("[EditVideo]Output_error err:%v", err)
	}
	return nil
}

func subClip(clip moviego.Video, fragDuration float64) moviego.Video {
	// 视频时长
	d := clip.Duration()
	// 视频片段个数
	fragCount := d / fragDuration
	if fragCount < 1 {
		return clip
	}
	var r int64
	// 视频片段个数小于等于3，取中间
	if fragCount <= 3 {
		r = 1
	} else {
		// 随机取中间的片段
		r = getRandomInt(1, int64(fragCount-2))
	}
	// log.Debugf("[subClip]fragCount:%v, r:%v, Duration:%v", fragCount, r, d)
	// 掐头去尾
	end := float64(r+1) * fragDuration
	if end > d {
		end = d
	}
	clip = clip.SubClip(float64(r)*fragDuration, end)
	return clip
}

func getDirAllVideo(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Errorf("[getDirAllVideo]Walk_error err:%v", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".mp4" {
			files = append(files, path)
			return nil
		}
		return nil
	})
	if err != nil {
		log.Errorf("[getDirAllVideo]Walk_error err:%v", err)
		return nil, err
	}

	return files, nil
}

// getRandomInt 获取指定范围随机数
func getRandomInt(min, max int64) int64 {
	if max-min <= 0 {
		return 0
	}
	// 设置随机数种子
	rand.NewSource(time.Now().UnixNano())
	randomNum := rand.Int63n(max-min) + min
	return randomNum
}
