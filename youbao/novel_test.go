// Package youbao 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-31 09:59
// @description:
package youbao

import (
	"context"
	"testing"

	"douyin_video/conf"
	"douyin_video/log"
)

func init() {
	conf.LoadConfig()
	log.InitLog()
}
func Test_applyFanQieKeyword(t *testing.T) {
	type args struct {
		ctx        context.Context
		bookId     int64
		authorName string
		bookName   string
		keyWord    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "已申请过",
			args: args{
				ctx:        context.Background(),
				bookId:     7321598679939288115,
				authorName: "长安见月",
				bookName:   "我死后，成了他的大体老师",
				keyWord:    []string{"休想离开我", "他的大体老师", "医学男朋友"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ApplyFanQieKeyword(tt.args.ctx, tt.args.bookId, tt.args.authorName, tt.args.bookName, tt.args.keyWord); (err != nil) != tt.wantErr {
				t.Errorf("applyFanQieKeyword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
