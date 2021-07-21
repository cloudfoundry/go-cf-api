package helpers

import "strings"

func Redact(input string, secrets []string) string {
	for _, secret := range secrets {
		input = strings.ReplaceAll(input, secret, "REDACTED")
	}
	return input
}
