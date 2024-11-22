package whitelist

import (
	"regexp"

	"github.com/zeromicro/go-zero/core/logx"
)

func IsInList(path string, whitelist []string) bool {
	for _, v := range whitelist {
		if path == v {
			return true
		}
	}
	return false
}

func IsInListByRegex(list []string, key string) bool {
	for _, s := range list {
		regex, err := regexp.Compile(s)
		if err != nil {
			logx.Error(err)
			return false
		}
		if regex.MatchString(key) {
			return true
		}
	}
	return false
}
