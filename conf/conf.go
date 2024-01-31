// Package config 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 17:57
// @description:
package conf

import (
	_ "embed"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var C config

const (
	DebugFile = "debug.log"
	InfoFile  = "info.log"
	ErrorFile = "error.log"
	WarnFile  = "warn.log"
)

type config struct {
	VideoResourceDir string `yaml:"videoResourceDir"`
	VideoOutputDir   string `yaml:"videoOutputDir"`
	AudioOutputDir   string `yaml:"audioOutputDir"`
	LogDir           string `yaml:"logDir"`
	NovelSource      string `yaml:"novelSource"`
	ContentType      int32  `yaml:"contentType"`
	FanQie           fanQie `yaml:"fanQie"`
	Audio            audio  `yaml:"audio"`
	Video            video  `yaml:"video"`
	YouBao           youbao `yaml:"youbao"`
}
type fanQie struct {
	Token             string            `yaml:"token"`
	AppId             int64             `yaml:"app_id"`
	MsToken           string            `yaml:"msToken"`
	XBogus            string            `yaml:"X-Bogus"`
	BookListUrl       string            `yaml:"bookListUrl"`
	BookInfoUrl       string            `yaml:"bookInfoUrl"`
	ChapterInfoUrl    string            `yaml:"chapterInfoUrl"`
	ChapterContentUrl string            `yaml:"chapterContentUrl"`
	Header            map[string]string `yaml:"header"`
	CbidList          []int64           `yaml:"cbidList"`
}
type audio struct {
	Speed               int32             `yaml:"speed"`
	AesKey              string            `yaml:"aesKey"`
	TxtLength           int               `yaml:"txtLength"`
	VoiceId             string            `yaml:"voiceId"`
	ConvertUrl          string            `yaml:"convertUrl"`
	GetVoiceAudioUrlWeb string            `yaml:"getVoiceAudioUrlWeb"`
	Sign                string            `yaml:"sign"`
	Header              map[string]string `yaml:"header"`
}
type video struct {
	VideoNum     int     `yaml:"videoNum"`
	VideoTime    float64 `yaml:"videoTime"`
	VideoType    string  `yaml:"videoType"`
	VideoWidth   int64   `yaml:"videoWidth"`
	VideoHeight  int64   `yaml:"videoHeight"`
	Speed        float64 `yaml:"speed"`
	FragDuration float64 `yaml:"fragDuration"`
}
type youbao struct {
	IsAddKeyword  bool                `yaml:"isAddKeyword"`
	ContentType   string              `yaml:"contentType"`
	LoginUrl      string              `yaml:"loginUrl"`
	UserInfoUrl   string              `yaml:"userInfoUrl"`
	InitSelectUrl string              `yaml:"initSelectUrl"`
	GetFiledUrl   string              `yaml:"getFiledUrl"`
	AddKeywordUrl string              `yaml:"addKeywordUrl"`
	Header        map[string]string   `yaml:"header"`
	Keywords      map[int64]ybKeyword `yaml:"keywords"`
}
type ybKeyword struct {
	AuthorName string   `yaml:"authorName"`
	BookName   string   `yaml:"bookName"`
	KeyWord    []string `yaml:"keyword"`
}

func LoadConfig() {
	confDir := os.Getenv("GO_WORKSPACE")
	if confDir == "" {
		confDir = "./"
	} else {
		confDir = confDir + "douyin_video/"
	}

	buf, err := os.ReadFile(confDir + "config.yaml")
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	var conf config
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		log.Fatal(err)
	}
	C = conf
}
