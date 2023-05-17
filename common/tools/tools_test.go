package tools

import (
	"log"
	"testing"
)

func TestMatchIDCardNum(t *testing.T) {
	//15位数字或18位数字，当为18位数时，最后一位可能为X
	if !MatchIDCardNum("340321199304078765") {
		log.Println("invalid auth_num")
	} else {
		log.Println("auth_num is ok")
	}

	if !MatchIDCardNum("34032119930407876X") {
		log.Println("invalid auth_num")
	} else {
		log.Println("auth_num is ok")
	}
	if !MatchIDCardNum("34032119930407876x") {
		log.Println("invalid auth_num")
	} else {
		log.Println("auth_num is ok")
	}

	//位数不对
	if !MatchIDCardNum("34032119930407") {
		log.Println("invalid auth_num")
	} else {
		log.Println("auth_num is ok")
	}

	//位数不对
	if !MatchIDCardNum("137673245") {
		log.Println("invalid auth_num")
	} else {
		log.Println("auth_num is ok")
	}

	//格式不对
	if !MatchIDCardNum("34032119930407876Y") {
		log.Println("invalid auth_num")
	} else {
		log.Println("auth_num is ok")
	}

	//格式不对
	if !MatchIDCardNum("340321199@@407876Y") {
		log.Println("invalid auth_num")
	} else {
		log.Println("auth_num is ok")
	}
}

func TestMatchHttpUrl(t *testing.T) {
	if !MatchHttpUrl("http:") {
		log.Println("url is invalid")
	} else {
		log.Println("url is ok")
	}

	if !MatchHttpUrl("https:") {
		log.Println("url is invalid")
	} else {
		log.Println("url is ok")
	}

	if !MatchHttpUrl("https://") {
		log.Println("url is invalid")
	} else {
		log.Println("url is ok")
	}

	if !MatchHttpUrl("http://") {
		log.Println("url is invalid")
	} else {
		log.Println("url is ok")
	}

}

func TestMatchEmail(t *testing.T) {
	//	regRuler := "^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$"
	if !MatchEmail("13@qq.com") {
		log.Println("invalid email")
		return
	} else {
		log.Println("email ok")
	}

	if !MatchEmail("13@163.com") {
		log.Println("invalid email")
		return
	} else {
		log.Println("email ok")
	}
	if !MatchEmail("1345314444444444455555555555555@bd5134531444444444445333333333555555555555555555555534354334411.com") {
		log.Println("invalid email")
		return
	} else {
		log.Println("email ok")
	}
}

func TestMatchPhoneNum(t *testing.T) {
	if !MatchPhoneNum("13516355566") {
		log.Println("invalid phone num")
		return
	} else {
		log.Println("phone num is ok")
	}
}
