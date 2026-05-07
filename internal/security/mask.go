package security

import "regexp"

var sensitivePatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)(password\s*=\s*)("[^"]+"|\S+)`),
	regexp.MustCompile(`(?i)(token\s*=\s*)("[^"]+"|\S+)`),
	regexp.MustCompile(`(?i)(secret\s*=\s*)("[^"]+"|\S+)`),
	regexp.MustCompile(`(?i)(access_key\s*=\s*)("[^"]+"|\S+)`),
	regexp.MustCompile(`(?i)(secret_key\s*=\s*)("[^"]+"|\S+)`),
	regexp.MustCompile(`(?i)(api_key\s*=\s*)("[^"]+"|\S+)`),
	regexp.MustCompile(`(?i)(client_secret\s*=\s*)("[^"]+"|\S+)`),
	regexp.MustCompile(`(?i)(private_key\s*=\s*)("[^"]+"|\S+)`),
}

func MaskLine(line string) string {
	out := line

	for _, pattern := range sensitivePatterns {
		out = pattern.ReplaceAllString(out, `${1}<masked>`)
	}

	return out
}
