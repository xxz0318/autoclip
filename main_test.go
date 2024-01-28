// Package main 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 23:10
// @description:
package main

import (
	"context"
	"testing"

	"douyin_video/conf"
	"douyin_video/log"
	"douyin_video/novel"
)

func init() {
	conf.LoadConfig()
	log.InitLog()
}
func Test_txtToAudio(t *testing.T) {
	type args struct {
		ctx      context.Context
		novelObj novel.Novel
		bookId   int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				ctx:      context.Background(),
				novelObj: novel.FanQie{},
				bookId:   7327542584295820350,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txtToAudio(tt.args.ctx, tt.args.novelObj, tt.args.bookId)
		})
	}
}
