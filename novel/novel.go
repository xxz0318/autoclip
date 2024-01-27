// Package novel 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 21:21
// @description:
package novel

import "context"

type Novel interface {
	GetBookList(ctx context.Context, pageIndex, pageSize int32) ([]book, int64, error)
	GetBookInfo(ctx context.Context, bookId int64) (book, error)
	GetChapterInfo(ctx context.Context, bookId int64) (chapter, error)
	GetChapterContent(ctx context.Context, bookId int64, itemId int64) (string, error)
	GetChapterContentByBookId(ctx context.Context, bookId int64) (string, error)
}
