// Package main 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 23:10
// @description:
package main

import (
	"context"
	"reflect"
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
				bookId:   7326809253237228094,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txtToAudio(tt.args.ctx, tt.args.novelObj, tt.args.bookId)
		})
	}
}

func TestAesDecrypt(t *testing.T) {
	type args struct {
		crypted string
		key     string
		iv      string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				crypted: "g2Ts2f59KhPf6TpuqI6SbuH7qJAmbj0PfmspkrDvCc2msI9aKnALGiUhCyTIL9Yd7xW5UPWW1hH6hgRDfXdTO3Pag3FiGofVVDBhCQbdgnuXMAXal1rQ++8GWB0Uhrj86aZjR1iWyP7ODGXe6JbUyQ==",
				key:     "abcdefgabcdefg12",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesDecrypt(tt.args.crypted, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesDecrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
