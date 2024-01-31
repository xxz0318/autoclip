// Package youbao 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-31 10:17
// @description:
package youbao

import (
	"context"
	"testing"
)

func Test_getUserInfo(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := getUserInfo(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("getUserInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
