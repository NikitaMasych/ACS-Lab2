package client

import (
	"sort"
	"strings"
)

func ProcessRequest(context string) string {
	characters := strings.Split(context, "")
	sort.Strings(characters)
	return strings.Join(characters, "")
}
