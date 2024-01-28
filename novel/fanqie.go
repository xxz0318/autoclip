// Package main 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 17:33
// @description:
package novel

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"douyin_video/conf"
	"douyin_video/log"
	"github.com/bytedance/sonic"
	"github.com/spf13/cast"
)

type book struct {
	BookId     string      `json:"book_id"`
	BookName   string      `json:"book_name"`
	WordNum    int         `json:"word_num"`
	Score      float64     `json:"score"`
	ThumbUrl   string      `json:"thumb_url"`
	Roles      interface{} `json:"roles"`
	Categories []struct {
		CategoryId   int    `json:"category_id"`
		CategoryName string `json:"category_name"`
		Status       int    `json:"status"`
	} `json:"categories"`
	BookAbstract                string `json:"book_abstract"`
	CreationStatus              int    `json:"creation_status"`
	Author                      string `json:"author"`
	ChapterNum                  int    `json:"chapter_num"`
	CopyrightType               int    `json:"copyright_type"`
	Genre                       int    `json:"genre"`
	CanRecommendChapterNum      int    `json:"can_recommend_chapter_num"`
	VideoRecommandDurationLimit int    `json:"video_recommand_duration_limit"`
	BookRecommandDurationLimit  int    `json:"book_recommand_duration_limit"`
	Subabstract                 string `json:"subabstract"`
	FirstOnlineTime             string `json:"first_online_time"`
	IsExclusive                 int    `json:"is_exclusive"`
	RecentIncome                int    `json:"recent_income"`
	RecentHighViewCount         int    `json:"recent_high_view_count"`
}

type bookListRsp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	LogId   string `json:"log_id"`
	Data    struct {
		RankBooks []book `json:"rank_books"`
		Total     int64  `json:"total"`
		PageIndex int32  `json:"page_index"`
	} `json:"data"`
}
type bookRsp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	LogId   string `json:"log_id"`
	Data    struct {
		BookList []book `json:"book_list"`
		Total    int64  `json:"total"`
	} `json:"data"`
}

type chapter struct {
	ItemId           string `json:"item_id"`
	Index            int32  `json:"index"`
	ChapterName      string `json:"chapter_name"`
	Content          string `json:"content"`
	PayInfo          bool   `json:"pay_info"`
	AudioUrl         string `json:"audio_url"`
	VideoUrl         string `json:"video_url"`
	VideoDownloadUrl string `json:"video_download_url"`
}
type chapterRsp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	LogId   string `json:"log_id"`
	Data    struct {
		ChapterList []chapter `json:"chapter_list"`
		Total       int64     `json:"total"`
	} `json:"data"`
}

type chapterContentRsp struct {
	Code    int32   `json:"code"`
	Message string  `json:"message"`
	LogId   string  `json:"log_id"`
	Data    chapter `json:"data"`
}

type FanQie struct {
}

func (f FanQie) GetChapterContentByBookId(ctx context.Context, bookId int64) (string, error) {
	chapterInfo, err := f.GetChapterInfo(ctx, bookId)
	if err != nil {
		return "", err
	}
	content, err := f.GetChapterContent(ctx, bookId, cast.ToInt64(chapterInfo.ItemId))
	if err != nil {
		return "", err
	}
	return content, nil
}

// GetBookList 获取书籍列表
func (f FanQie) GetBookList(ctx context.Context, pageIndex, pageSize int32) ([]book, int64, error) {
	if pageIndex < 0 {
		pageIndex = 0
	}
	if pageSize < 0 {
		pageSize = 10
	}
	bookListUrl := replaceFanQiePlaceholder(ctx, conf.C.FanQie.BookListUrl)
	bookListUrl = strings.Replace(bookListUrl, "{page_index}", cast.ToString(pageIndex), -1)
	bookListUrl = strings.Replace(bookListUrl, "{page_size}", cast.ToString(pageSize), -1)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", bookListUrl, nil)
	setFanQieHeader(ctx, req)
	rsp, reqErr := client.Do(req)

	if reqErr != nil {
		log.Errorf("[GetBookList]NewRequest_error reqErr:%v", reqErr)
		return nil, 0, reqErr
	}
	defer rsp.Body.Close()
	var bookInfoListRsp bookListRsp
	content, _ := io.ReadAll(rsp.Body)
	err := sonic.Unmarshal(content, &bookInfoListRsp)
	if err != nil {
		log.Errorf("[GetBookList]Unmarshal_error err:%v", err)
		return nil, 0, err
	}
	return bookInfoListRsp.Data.RankBooks, bookInfoListRsp.Data.Total, nil
}

// GetBookInfo 获取书籍信息
func (f FanQie) GetBookInfo(ctx context.Context, bookId int64) (book, error) {
	var bookInfo book
	if bookId <= 0 {
		return bookInfo, nil
	}

	bookInfoUrl := replaceFanQiePlaceholder(ctx, conf.C.FanQie.BookInfoUrl)
	bookInfoUrl = strings.Replace(bookInfoUrl, "{bookid}", cast.ToString(bookId), -1)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", bookInfoUrl, nil)
	setFanQieHeader(ctx, req)
	rsp, reqErr := client.Do(req)

	if reqErr != nil {
		log.Errorf("[GetBookInfo]NewRequest_error reqErr:%v", reqErr)
		return bookInfo, reqErr
	}
	defer rsp.Body.Close()
	var bookInfoRsp bookRsp
	content, _ := io.ReadAll(rsp.Body)
	err := sonic.Unmarshal(content, &bookInfoRsp)
	if err != nil {
		log.Errorf("[GetBookList]Unmarshal_error err:%v", err)
		return bookInfo, err
	}
	if len(bookInfoRsp.Data.BookList) <= 0 {
		return bookInfo, fmt.Errorf("bookInfoRsp.Data.BookList is empty, bookId:%d", bookId)
	}
	return bookInfoRsp.Data.BookList[0], nil
}

// GetChapterInfo 获取章节信息
func (f FanQie) GetChapterInfo(ctx context.Context, bookId int64) (chapter, error) {
	var chapterInfo chapter
	if bookId <= 0 {
		return chapterInfo, nil
	}

	chapterInfoUrl := replaceFanQiePlaceholder(ctx, conf.C.FanQie.ChapterInfoUrl)
	chapterInfoUrl = strings.Replace(chapterInfoUrl, "{bookid}", cast.ToString(bookId), -1)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", chapterInfoUrl, nil)
	setFanQieHeader(ctx, req)
	rsp, reqErr := client.Do(req)

	if reqErr != nil {
		log.Errorf("[GetChapterInfo]NewRequest_error reqErr:%v", reqErr)
		return chapterInfo, reqErr
	}
	defer rsp.Body.Close()
	var chapterInfoRsp chapterRsp
	content, _ := io.ReadAll(rsp.Body)

	err := sonic.Unmarshal(content, &chapterInfoRsp)
	if err != nil {
		log.Errorf("[GetChapterInfo]Unmarshal_error err:%v", err)
		return chapterInfo, err
	}
	if len(chapterInfoRsp.Data.ChapterList) <= 0 {
		return chapterInfo, fmt.Errorf("chapterInfoRsp.Data.ChapterList, bookId:%d", bookId)
	}
	return chapterInfoRsp.Data.ChapterList[0], nil
}

// GetChapterContent 获取章节内容
func (f FanQie) GetChapterContent(ctx context.Context, bookId, itemId int64) (string, error) {
	if bookId <= 0 || itemId <= 0 {
		return "", nil
	}

	chapterContentUrl := replaceFanQiePlaceholder(ctx, conf.C.FanQie.ChapterContentUrl)
	chapterContentUrl = strings.Replace(chapterContentUrl, "{bookid}", cast.ToString(bookId), -1)
	chapterContentUrl = strings.Replace(chapterContentUrl, "{itemid}", cast.ToString(itemId), -1)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", chapterContentUrl, nil)
	setFanQieHeader(ctx, req)
	rsp, reqErr := client.Do(req)

	if reqErr != nil {
		log.Errorf("[GetChapterContent]NewRequest_error reqErr:%v", reqErr)
		return "", reqErr
	}
	defer rsp.Body.Close()
	var chapterInfoRsp chapterContentRsp
	content, _ := io.ReadAll(rsp.Body)

	err := sonic.Unmarshal(content, &chapterInfoRsp)
	if err != nil {
		log.Errorf("[GetChapterContent]Unmarshal_error err:%v", err)
		return "", err
	}

	return chapterInfoRsp.Data.Content, nil
}

func setFanQieHeader(ctx context.Context, req *http.Request) {
	header := conf.C.FanQie.Header
	for k, v := range header {
		req.Header.Set(k, replaceFanQiePlaceholder(ctx, v))
	}
}

// replacePlaceholder 替换占位符
func replaceFanQiePlaceholder(ctx context.Context, str string) string {
	if strings.Contains(str, "{X-Bogus}") {
		str = strings.Replace(str, "{X-Bogus}", conf.C.FanQie.XBogus, -1)
	}
	if strings.Contains(str, "{token}") {
		str = strings.Replace(str, "{token}", conf.C.FanQie.Token, -1)
	}
	if strings.Contains(str, "{appid}") {
		str = strings.Replace(str, "{appid}", cast.ToString(conf.C.FanQie.AppId), -1)
	}
	if strings.Contains(str, "{msToken}") {
		str = strings.Replace(str, "{msToken}", conf.C.FanQie.MsToken, -1)
	}
	return str
}
