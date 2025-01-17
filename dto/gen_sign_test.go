package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buhuang28/go-alipay/logs"
	log "github.com/sirupsen/logrus"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGenReqSign(t *testing.T) {
	var alipayDto AlipayPublicReq
	typeOf := reflect.TypeOf(alipayDto)
	tag2Field := make(map[string]string)

	template := `buf.WriteString("&%s=")
buf.WriteString(a.%s)
`
	template1 := `if a.%s != "" {
	buf.WriteString("&%s=")
	buf.WriteString(a.%s)
}
`

	template2 := `if a.%s != "" {
		buf.WriteString("%s=")
		buf.WriteString(a.%s)
	}
`

	template3 := `buf.WriteString("%s=")
buf.WriteString(a.%s)
`

	var result bytes.Buffer
	result.WriteString("var buf bytes.Buffer\r\n")

	var keys []string
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		tag := field.Tag.Get("json")
		tag2Field[tag] = field.Name
		keys = append(keys, tag)
	}
	sort.Strings(keys)
	for i, v := range keys {
		if i == 0 {
			if strings.HasSuffix(v, ",omitempty") {
				tag := strings.ReplaceAll(v, ",omitempty", "")
				result.WriteString(fmt.Sprintf(template1, tag2Field[v], tag, tag2Field[v]))
			} else {
				result.WriteString(fmt.Sprintf(template3, v, tag2Field[v]))
			}
			continue
		}
		if strings.HasSuffix(v, ",omitempty") {
			tag := strings.ReplaceAll(v, ",omitempty", "")
			result.WriteString(fmt.Sprintf(template2, tag2Field[v], tag, tag2Field[v]))
		} else {
			result.WriteString(fmt.Sprintf(template, v, tag2Field[v]))
		}
	}
	fmt.Println(result.String())
}

func TestGenRespSign(t *testing.T) {
	var alipayDto AlipayTradePrecreateResp
	typeOf := reflect.TypeOf(alipayDto.AlipayTradePrecreateResponse)
	tag2Field := make(map[string]string)

	template := `buf.WriteString("&%s=")
buf.WriteString(a.%s)
`
	template1 := `if a.%s != "" {
	buf.WriteString("&%s=")
	buf.WriteString(a.%s)
}
`

	template2 := `if a.%s != "" {
		buf.WriteString("%s=")
		buf.WriteString(a.%s)
	}
`

	template3 := `buf.WriteString("%s=")
buf.WriteString(a.%s)
`

	var result bytes.Buffer
	result.WriteString("var buf bytes.Buffer\r\n")

	var keys []string
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		tag := field.Tag.Get("json")
		tag2Field[tag] = field.Name
		keys = append(keys, tag)
	}
	sort.Strings(keys)
	for i, v := range keys {
		if i == 0 {
			if strings.HasSuffix(v, ",omitempty") {
				tag := strings.ReplaceAll(v, ",omitempty", "")
				result.WriteString(fmt.Sprintf(template1, tag2Field[v], tag, tag2Field[v]))
			} else {
				result.WriteString(fmt.Sprintf(template3, v, tag2Field[v]))
			}
			continue
		}
		if strings.HasSuffix(v, ",omitempty") {
			tag := strings.ReplaceAll(v, ",omitempty", "")
			result.WriteString(fmt.Sprintf(template2, tag2Field[v], tag, tag2Field[v]))
		} else {
			result.WriteString(fmt.Sprintf(template, v, tag2Field[v]))
		}
	}
	fmt.Println(result.String())
}

func TestJSON(t *testing.T) {

	js := `[
    {
        "areaCode": "440200000000",
        "areaName": "韶关市",
        "pointNum": 1,
        "pointResponses": [
            {
                "id": "79319010c2d7314a78a2fcc43b0fcf90",
                "cityCode": "440200000000",
                "cityName": "韶关市",
                "areaCode": "440203000000",
                "areaName": "武江区",
                "objectId": "230fb5c965da4dacb1c6649960733c83",
                "name": "韶关武江区定点单位002",
                "property": "ME",
                "periods": null,
                "taskName": null,
                "priceAddress": null,
                "itemNum": 1,
                "orgId": null,
                "parentAreaCode": "440200000000"
            }
        ]
    },
    {
        "areaCode": "440100000000",
        "areaName": "广州市",
        "pointNum": 4,
        "pointResponses": [
            {
                "id": "1e8df4854556947769cb7193bae36e62",
                "cityCode": "440100000000",
                "cityName": "广州市",
                "areaCode": "440106000000",
                "areaName": "天河区",
                "objectId": "3e6826c85c28632df7cb4fcc64b4f5fa",
                "name": "小鲜驿站（车陂店）",
                "property": "ME",
                "periods": null,
                "taskName": null,
                "priceAddress": null,
                "itemNum": 1,
                "orgId": null,
                "parentAreaCode": "440100000000"
            },
            {
                "id": "ce768299555f8477a9d7210b9a3ee67d",
                "cityCode": "440100000000",
                "cityName": "广州市",
                "areaCode": "440111000000",
                "areaName": "白云区",
                "objectId": "e1ed7bad44dc4ba58e4232d1cb3ea44d",
                "name": "白云区定点单位001",
                "property": "ME",
                "periods": "DAY",
                "taskName": "陈罗志2024监测任务(不要删)",
                "priceAddress": null,
                "itemNum": 1,
                "orgId": null,
                "parentAreaCode": "440100000000"
            },
            {
                "id": "16961e65ebfb8da829a9dd7d361eb608",
                "cityCode": "440100000000",
                "cityName": "广州市",
                "areaCode": "440106000000",
                "areaName": "天河区",
                "objectId": "3e6826c85c28632df7cb4fcc64b4f5fa",
                "name": "小鲜驿站（车陂店）",
                "property": "ME",
                "periods": "MONTH",
                "taskName": "蔬菜价格监测2024",
                "priceAddress": null,
                "itemNum": 1,
                "orgId": null,
                "parentAreaCode": "440100000000"
            },
            {
                "id": "883a3094edee8e896cfa7ce8c32e1435",
                "cityCode": "440100000000",
                "cityName": "广州市",
                "areaCode": "440106000000",
                "areaName": "天河区",
                "objectId": "3e6826c85c28632df7cb4fcc64b4f5fa",
                "name": "小鲜驿站（车陂店）",
                "property": "ME",
                "periods": "DAY",
                "taskName": "梁伟乐测试",
                "priceAddress": null,
                "itemNum": 1,
                "orgId": null,
                "parentAreaCode": "440100000000"
            }
        ]
    }
]`

	logs.LogInit()
	m := make(map[string]any)
	var mer []Mer
	err := json.Unmarshal([]byte(js), &mer)
	if err != nil {
		panic(err)
	}
	for _, v := range mer {
		m2 := make(map[string][]PointResponses)
		for _, v2 := range v.PointResponses {
			m2[v2.Name] = append(m2[v2.Name], v2)
		}
		m[v.AreaCode] = m2
	}
	marshal, _ := json.Marshal(m)
	log.Info(string(marshal))
}

type Mer struct {
	AreaCode       string           `json:"areaCode"`
	AreaName       string           `json:"areaName"`
	PointNum       int              `json:"pointNum"`
	PointResponses []PointResponses `json:"pointResponses"`
}

type PointResponses struct {
	ID             string      `json:"id"`
	CityCode       string      `json:"cityCode"`
	CityName       string      `json:"cityName"`
	AreaCode       string      `json:"areaCode"`
	AreaName       string      `json:"areaName"`
	ObjectID       string      `json:"objectId"`
	Name           string      `json:"name"`
	Property       string      `json:"property"`
	Periods        interface{} `json:"periods"`
	TaskName       interface{} `json:"taskName"`
	PriceAddress   interface{} `json:"priceAddress"`
	ItemNum        int         `json:"itemNum"`
	OrgID          interface{} `json:"orgId"`
	ParentAreaCode string      `json:"parentAreaCode"`
}
