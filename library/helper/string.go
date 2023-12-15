package helper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// IsEmpty - check if string is empty
func IsEmpty(s string) bool {
	return strings.Trim(s, " ") == ""
}

// IsValidHash validate hex address
func IsValidHash(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{64}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

func Add0xPrefix(data string) string {
	return fmt.Sprintf("0x%s", data)
}
