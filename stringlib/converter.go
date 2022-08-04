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

func StructToMap(s interface{}) (map[string]interface{}, error) {
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(inrec, &inInterface)
	return inInterface, err
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
