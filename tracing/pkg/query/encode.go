package query

import (
	"net/url"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Encode(v url.Values, sep byte, filters []string, fieldNames []string) string {
	if len(v) == 0 {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.QueryEscape(k)
		if slices.Contains(filters, keyEscaped) {
			continue
		}

		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte(sep)
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			_, err := strconv.Atoi(v)
			if err != nil {
				buf.WriteString("%s")
			} else {
				buf.WriteString("%d")
			}
		}
	}

	sort.Strings(fieldNames)
	var bufValue strings.Builder
	for _, k := range fieldNames {

		if slices.Contains(filters, strings.ToLower(k)) {
			continue
		}

		if bufValue.Len() > 0 {
			bufValue.WriteByte(',')
		}
		bufValue.WriteString("q." + k)
	}

	return "fmt.Sprintf(\"" + buf.String() + "\"," + bufValue.String() + ")"
}
