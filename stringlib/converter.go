package stringlib

import (
	"encoding/json"
	"fmt"
	"sort"
)

func MapToSortedStr(mapS map[string]interface{}) string {
	keys := make([]string, 0, len(mapS))

	for k := range mapS {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))

	res := ""
	for _, k := range keys {
		res = fmt.Sprintf("%s%s%v", res, k, mapS[k])
	}

	return res
}

// Bug : format change if not string
func StructToMap(s interface{}) (map[string]interface{}, error) {
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(inrec, &inInterface)
	return inInterface, err
}

func StructToJsonString(strct interface{}) (string, error) {
	jsonStr, err := json.Marshal(strct)
	if err != nil {
		return "", err
	}

	return string(jsonStr), err
}

func MapToParam(url string, mapS map[string]interface{}) string {
	res := url
	for k, v := range mapS {
		if res == url {
			res = fmt.Sprintf("%s%s=%v", res, k, v)
			continue
		}
		res = fmt.Sprintf("%s&%s=%v", res, k, v)
	}

	return res
}

func ListToString[V int | string | float32 | float64](l []V) string {
	var str string
	str = fmt.Sprintf("%s[", str)
	for i, s := range l {
		if i == 0 {
			str = fmt.Sprintf("%s%v", str, s)
			continue
		}
		str = fmt.Sprintf("%s,%v", str, s)
	}
	str = fmt.Sprintf("%s]", str)

	return str
}
