// Package youbao 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-30 19:40
// @description:
package youbao

import (
	"context"
	"fmt"

	"douyin_video/conf"
	"douyin_video/log"
	"douyin_video/utils"
)

type fieldRspStruct struct {
	CommonRsp
	Data struct {
		FormField []FieldData `json:"formField"`
		CopyUrl   string      `json:"copy_url"`
	} `json:"data"`
}

type FieldData struct {
	FrontMsg string   `json:"front_msg"`
	MaxLen   int      `json:"max_len"`
	MinLen   int      `json:"min_len"`
	Name     string   `json:"name"`
	Field    string   `json:"field"`
	VarType  int      `json:"var_type"`
	IsEmpty  int      `json:"is_empty"`
	IsCopy   int      `json:"is_copy"`
	Type     int      `json:"type"`
	Value    string   `json:"value"`
	TypeList []string `json:"typeList"`
}

type addReq struct {
	FormData  [][]FieldData `json:"formData"`
	ProjectId string        `json:"project_id"`
}

// ApplyFanQieKeyword 申请番茄关键词
func ApplyFanQieKeyword(ctx context.Context, bookId int64, authorName, bookName string, keyWord []string) error {
	if len(keyWord) < 1 {
		log.Errorf("[applyFanQieKeyword]keyWord_len_error:%+v", keyWord)
		return fmt.Errorf("[applyFanQieKeyword]keyWord_len_error:%+v", keyWord)
	}
	selectReq := struct {
		Id        int    `json:"id"`
		ProjectId string `json:"project_id"`
	}{
		Id:        0,
		ProjectId: "62",
	}
	fieldReq := struct {
		ProjectId string `json:"project_id"`
	}{
		ProjectId: "62",
	} // 番茄是62
	fieldContentRsp := new(fieldRspStruct)
	// 必须先请求initSelectData接口，否则请求initFormFiled接口无返回数据
	utils.Request(ctx, conf.C.YouBao.InitSelectUrl, "POST", selectReq, conf.C.YouBao.Header, nil)
	_, err := utils.Request(ctx, conf.C.YouBao.GetFiledUrl, "POST", fieldReq, conf.C.YouBao.Header, fieldContentRsp)
	if err != nil {
		log.Errorf("[applyFanQieKeyword]Request_error err:%v", err)
		return err
	}
	if fieldContentRsp.Code != 1 && fieldContentRsp.Show != 0 {
		log.Errorf("[applyFanQieKeyword]fieldContentRsp_code error:%+v", fieldContentRsp)
		return fmt.Errorf("[applyFanQieKeyword]fieldContentRsp_code:%d", fieldContentRsp.Code)
	}

	var keywordData [][]FieldData
	for _, v := range keyWord {
		fieldList := fieldContentRsp.Data.FormField
		for k, fieldInfo := range fieldList {
			if fieldInfo.Field == "keyword" {
				fieldList[k].Value = v
			}
			if fieldInfo.Field == "book_id" {
				fieldList[k].Value = fmt.Sprintf("%d", bookId)
			}
			if fieldInfo.Field == "book_name" {
				fieldList[k].Value = bookName
			}
			if fieldInfo.Field == "author" {
				fieldList[k].Value = authorName
			}
		}
		keywordData = append(keywordData, fieldList)
	}

	log.Debugf("[applyFanQieKeyword]keywordData:%+v", keywordData)

	if len(keywordData) < 1 || len(keywordData[0]) < 1 {
		log.Errorf("[applyFanQKW]keywordData_len_error:%+v", keywordData)
		return fmt.Errorf("[applyFanQKW]keywordData_len_error:%+v", keywordData)
	}
	addReqData := addReq{
		FormData:  keywordData,
		ProjectId: "62",
	}
	addRsp := new(CommonRsp)
	_, err = utils.Request(ctx, conf.C.YouBao.AddKeywordUrl, "POST", addReqData, conf.C.YouBao.Header, addRsp)
	if err != nil {
		log.Errorf("[applyFanQieKeyword]Request_error err:%v", err)
		return err
	}
	log.Infof("[applyFanQieKeyword]addRsp:%+v", addRsp)
	return nil

}
