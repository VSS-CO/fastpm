package resolver

import (
	"strings"
)

func MatchVersion(requested, available string) bool {
    return requested == "latest" || requested == available
}

func ParseRange(r string) ([]string, error) {
    return strings.Split(r, ","), nil
}
