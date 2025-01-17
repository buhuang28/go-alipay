package alipay

import (
	"github.com/buhuang28/go-alipay/caches"
	"github.com/buhuang28/go-alipay/utils"
	"net/url"
	"sort"
	"strings"
)

// 验证异步通知的签名
func VerifySign(data string) error {
	data, _ = url.QueryUnescape(data)
	split := strings.Split(data, "&")
	sign := ""
	//var signBuf bytes.Buffer
	var signContent []string
	for _, v := range split {
		if strings.Contains(v, "sign") {
			if len(v) > 20 {
				sign = v[5:]
			}
			continue
		}
		//signBuf.WriteString(v)
		signContent = append(signContent, v)
	}
	sort.Strings(signContent)
	src := strings.Join(signContent, "&")
	return utils.Rsa2PubSign(caches.AlipayConfig.AlipayPublicKey, src, sign)
}
