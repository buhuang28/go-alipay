package dto

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGenSign(t *testing.T) {
	var alipayDto AlipayPublicDTO
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
				tag := strings.TrimRight(v, ",omitempty")
				result.WriteString(fmt.Sprintf(template1, tag2Field[v], tag, tag2Field[v]))
			} else {
				result.WriteString(fmt.Sprintf(template3, v, tag2Field[v]))
			}
			continue
		}
		if strings.HasSuffix(v, ",omitempty") {
			tag := strings.TrimRight(v, ",omitempty")
			result.WriteString(fmt.Sprintf(template2, tag2Field[v], tag, tag2Field[v]))
		} else {
			result.WriteString(fmt.Sprintf(template, v, tag2Field[v]))
		}
	}
	fmt.Println(result.String())
}
