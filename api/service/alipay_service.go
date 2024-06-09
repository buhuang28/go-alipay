package service

import (
	"github.com/buhuang28/go-alipay/caches"
	"github.com/buhuang28/go-alipay/utils"
	"net/url"
	"sort"
	"strings"
)

func AlipayVerifySign(data string) error {
	data, _ = url.QueryUnescape(data)
	split := strings.Split(data, "&")
	sign := ""
	var signContent []string
	for _, v := range split {
		if strings.Contains(v, "sign") {
			if len(v) > 20 {
				sign = v[5:]
			}
			continue
		}
		signContent = append(signContent, v)
	}
	sort.Strings(signContent)
	src := strings.Join(signContent, "&")
	return utils.Rsa2PubSign(caches.AlipayConfigCache.AlipayPublicKey, src, sign)
}
