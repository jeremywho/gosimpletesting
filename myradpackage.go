package myradpackage

import (
	"strings"
)

func GetProtocol(url string) string {
    return strings.Split(url, ":")[0]
}