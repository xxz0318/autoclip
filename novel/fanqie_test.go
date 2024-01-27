// Package main 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 18:52
// @description:
package novel

import (
	"context"
	"reflect"
	"testing"

	"douyin_video/conf"
)

func init() {
	conf.LoadConfig()
}

func Test_fanQie_GetBookList(t *testing.T) {
	type args struct {
		ctx       context.Context
		pageIndex int32
		pageSize  int32
	}
	tests := []struct {
		name    string
		args    args
		want    []book
		want1   int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				ctx:       context.Background(),
				pageIndex: 0,
				pageSize:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fanQie{}
			got, got1, err := f.GetBookList(tt.args.ctx, tt.args.pageIndex, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBookList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			t.Logf("GetBookList() got = %+v", got)
			// }
			if got1 != tt.want1 {
				t.Errorf("GetBookList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fanQie_GetBookInfo(t *testing.T) {
	type args struct {
		ctx    context.Context
		bookId int64
	}
	tests := []struct {
		name    string
		args    args
		want    book
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "7314551982457359412",
			args: args{
				ctx:    context.Background(),
				bookId: 7314551982457359412,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fanQie{}
			got, err := f.GetBookInfo(tt.args.ctx, tt.args.bookId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBookInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBookInfo() got = %+v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fanQie_GetChapterInfo(t *testing.T) {
	type args struct {
		ctx    context.Context
		bookId int64
	}
	tests := []struct {
		name    string
		args    args
		want    chapter
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "7326809253237228094",
			args: args{
				ctx:    context.Background(),
				bookId: 7326809253237228094,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fanQie{}
			got, err := f.GetChapterInfo(tt.args.ctx, tt.args.bookId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChapterInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChapterInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fanQie_GetChapterContent(t *testing.T) {
	type args struct {
		ctx    context.Context
		bookId int64
		itemId int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "7326809253237228094",
			args: args{
				ctx:    context.Background(),
				bookId: 7326809253237228094,
				itemId: 7326801982235083326,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fanQie{}
			got, err := f.GetChapterContent(tt.args.ctx, tt.args.bookId, tt.args.itemId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChapterContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetChapterContent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
