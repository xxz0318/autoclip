// Package video 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-28 19:56
// @description:
package video

import (
	"context"
	"reflect"
	"testing"

	"douyin_video/conf"
	"douyin_video/log"
)

func init() {
	conf.LoadConfig()
	log.InitLog()
}
func Test_getDirAllVideo(t *testing.T) {
	type args struct {
		dirPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				dirPath: conf.C.VideoResourceDir + "jieya/美甲-高清720/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDirAllVideo(tt.args.dirPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDirAllVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDirAllVideo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEditVideo(t *testing.T) {
	type args struct {
		ctx          context.Context
		videoTime    float64
		videoType    string
		width        int64
		length       int64
		speed        float64
		fragDuration float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.Background(),
				videoTime:    1200,
				videoType:    "jieya/切沙子太空沙塑形/",
				width:        1080,
				length:       1920,
				speed:        1,
				fragDuration: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EditVideo(tt.args.ctx, tt.args.videoTime, tt.args.videoType, tt.args.width, tt.args.length, tt.args.speed, tt.args.fragDuration); (err != nil) != tt.wantErr {
				t.Errorf("EditVideo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
