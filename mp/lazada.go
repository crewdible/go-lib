package mp

import (
	"fmt"
	"strings"

	_enclib "github.com/crewdible/go-lib/encryption"
	_strlib "github.com/crewdible/go-lib/stringlib"
)

// Sign generator for Lazada
func GenerateSHA256WithParams(secret, path string, params map[string]interface{}) string {
	signStr := fmt.Sprintf("%s%s", path, _strlib.MapToSortedStr(params))
	signHash := _enclib.GenerateHMACSHA256(signStr, secret)
	return strings.ToUpper(signHash)
}
