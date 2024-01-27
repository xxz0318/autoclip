// Package config 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 17:57
// @description:
package conf

import (
	_ "embed"
	"log"

	"gopkg.in/yaml.v3"
)

var C config

const (
	DebugFile = "./logs/debug.log"
	InfoFile  = "./logs/info.log"
	ErrorFile = "./logs/error.log"
	WarnFile  = "./logs/warn.log"
)

type config struct {
	ResourceDir string `yaml:"resourceDir"`
	OutputDir   string `yaml:"outputDir"`
	NovelSource string `yaml:"novelSource"`
	FanQie      fanQie `yaml:"fanQie"`
	Audio       audio  `yaml:"audio"`
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
	AesKey              string            `yaml:"aesKey"`
	TxtLength           int               `yaml:"txtLength"`
	VoiceId             string            `yaml:"voiceId"`
	ConvertUrl          string            `yaml:"convertUrl"`
	GetVoiceAudioUrlWeb string            `yaml:"getVoiceAudioUrlWeb"`
	Sign                string            `yaml:"sign"`
	Header              map[string]string `yaml:"header"`
}

var (
	//go:embed config.yaml
	configFile string
)

func LoadConfig() {

	var conf config
	err := yaml.Unmarshal([]byte(configFile), &conf)
	if err != nil {
		log.Fatal(err)
	}
	C = conf
}
