package tools

import (
	"regexp"
)

func MatchIDCardNum(num string) bool {
	//15位数字或18位数字，当为18位数时，最后一位可能为X
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(num)
}

func MatchCorporateNum(num string) bool {
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)|(^.{18}$)"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(num)
}

func MatchPhoneNum(num string) bool {
	regRuler := "^1[3456789]\\d{9}$"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(num)
}

func MatchEmail(email string) bool {
	regRuler := "^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(email)
}

func MatchPostcode(postcode string) bool {
	regRuler := "^[0-9]\\d{5}$"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(postcode)
}

func MatchHttpUrl(url string) bool {
	regRuler := "http(s)?://[^\\s]*"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(url)
}
