// Package youbao 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-30 19:30
// @description:
package youbao

import (
	"context"
	"fmt"

	"douyin_video/conf"
	"douyin_video/log"
	"douyin_video/utils"
)

type CommonRsp struct {
	Code int    `json:"code"`
	Show int    `json:"show"`
	Msg  string `json:"msg"`
}

// Schoolo:
// d01cdb015332f142e0e6e510e8855f89
type loginReq struct {
	Terminal int    `json:"terminal"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	SceneId  int    `json:"scene_id"`
	Scene    int    `json:"scene"`
	DeviceId string `json:"deviceId"`
	Account  string `json:"account"`
}

type loginRsp struct {
	CommonRsp
	Data struct {
		Nickname        string      `json:"nickname"`
		Balance         string      `json:"balance"`
		FreezeMoney     int         `json:"freeze_money"`
		WithdrawalMoney string      `json:"withdrawal_money"`
		Money           string      `json:"money"`
		LeaderTel       interface{} `json:"leader_tel"`
		Sn              int         `json:"sn"`
		Mobile          string      `json:"mobile"`
		Avatar          string      `json:"avatar"`
		Token           string      `json:"token"`
		Type            int         `json:"type"`
	} `json:"data"`
}
type userInfo struct {
	CommonRsp
	Data struct {
		Id             int64  `json:"id"`
		Sn             int32  `json:"sn"`
		Sex            string `json:"sex"`
		Account        string `json:"account"`
		Nickname       string `json:"nickname"`
		RealName       string `json:"real_name"`
		Avatar         string `json:"avatar"`
		Mobile         string `json:"mobile"`
		CreateTime     string `json:"create_time"`
		VipId          int32  `json:"vip_id"`
		LeaderId       int32  `json:"leader_id"`
		VipTime        string `json:"vip_time"`
		Money          string `json:"money"`
		Balance        int32  `json:"balance"`
		Type           int32  `json:"type"`
		Wechat         string `json:"wechat"`
		IsAuth         int32  `json:"is_auth"`
		TypeTime       string `json:"type_time"`
		LeaderWechat   string `json:"Leaderwechat"`
		NewsNum        int32  `json:"news_num"`
		RandCode       string `json:"rand_code"`
		IsLogin        bool   `json:"is_login"`
		VipName        string `json:"vip_name"`
		LeaderNickname string `json:"leader_nickname"`
		HasPassword    bool   `json:"has_password"`
		HasAuth        bool   `json:"has_auth"`
		Version        string `json:"version"`
	} `json:"data"`
}

func login(ctx context.Context, phone, password string) {
	// loginUrl := "https://newapi.thed1g.com/api/login/account"
	// params := loginReq{
	// 	Terminal: 1,
	// 	Mobile:   phone,
	// 	Password: password,
	// 	SceneId:  101,
	// 	Scene:    1,
	// 	DeviceId: "17066141572005638757",
	// 	Account:  phone,
	// }
}

func getUserInfo(ctx context.Context) (int64, error) {
	userInfoRsp := new(userInfo)

	_, err := utils.Request(ctx, conf.C.YouBao.UserInfoUrl, "POST", nil, conf.C.YouBao.Header, userInfoRsp)
	if err != nil {
		log.Errorf("[getUserInfo]rsp_error:%v", err)
		return 0, err
	}

	if userInfoRsp.Code != 1 && userInfoRsp.Show != 0 {
		log.Errorf("[getUserInfo]userInfoRsp_code error:%+v", userInfoRsp)
		return 0, fmt.Errorf("[getUserInfo]userInfoRsp_code:%d", userInfoRsp.Code)
	}
	log.Infof("[getUserInfo]userInfoRsp:%+v", userInfoRsp)
	return userInfoRsp.Data.Id, nil
}
