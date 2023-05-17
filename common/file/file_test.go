package file

import (
	"encoding/json"
	"log"
	"testing"
)

type (
	LanguageInfo struct {
		LanguageInfos []Language `json:"language_info"`
	}

	Language struct {
		Name        string `json:"language_name" example:"abc"`     //语种名称
		Description string `json:"description" example:"abc"`       //语种描述
		Short       string `json:"language_short" example:"abc"`    //语种简称
		IconUrl     string `json:"language_icon_url" example:"abc"` //语种图标url
	}
)

func TestReadJsonFile(t *testing.T) {
	lans := new(LanguageInfo)
	byteValue := ReadJsonFile("/Users/nicker/sandbox/weiland/plat-base-service/multilingualconf/log", "language_info.json")

	json.Unmarshal([]byte(byteValue), &lans)
	//log.Println("language_info:", string(byteValue))
	log.Println("language_info lans:", lans)

	for _, lan := range lans.LanguageInfos {

		log.Println("language_info lan:", lan)
	}
}
