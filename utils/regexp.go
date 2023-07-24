package utils

import "regexp"

// PhoneValidate 校验手机号
/// v 手机号
func PhoneValidate(v string) bool {
	regRuler := "^(([^<>()[\\]\\\\.,;:\\s@\"]+(\\.[^<>()[\\]\\\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(v)
}