package document

import (
	"regexp"
	"strings"
)

var goizerNameRegex = regexp.MustCompile("[a-zA-Z0-9]+")

func isUpperToken(s string) bool {
	switch strings.ToLower(s) {
	case
		"api", "ascii", "cpu", "css", "dns", "eof", "guid", "html", "http",
		"https", "id", "ip", "json", "lhs", "qps", "ram", "rhs", "rpc",
		"sla", "smtp", "ssh", "tls", "ttl", "ui", "uid", "uuid", "uri",
		"url", "utf8", "vm", "xml", "xsrf", "xss":
		return true
	default:
		return false
	}
}

func GoNamePublic(name string) string {
	matches := goizerNameRegex.FindAllString(name, -1)
	n := ""
	for _, s := range matches {
		if isUpperToken(s) {
			n += strings.ToUpper(s)
			continue
		}
		n += strings.ToUpper(s[:1]) + s[1:]
	}
	return n
}

func GoNamePrivate(name string) string {
	matches := goizerNameRegex.FindAllString(name, -1)
	n := ""
	for i, s := range matches {
		if i == 0 {
			n += strings.ToLower(s[:1]) + s[1:]
		} else {
			if isUpperToken(s) {
				n += strings.ToUpper(s)
				continue
			}
			n += strings.ToUpper(s[:1]) + s[1:]
		}
	}
	return n
}
