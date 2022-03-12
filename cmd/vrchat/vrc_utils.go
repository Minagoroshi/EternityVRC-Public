package vrchat

import (
	"strings"
)

func TagsConverter(tags []byte) string {
	trust := "Visitor"
	if strings.Contains(string(tags), "system_trust_basic") {
		trust = "New User"
	}
	if strings.Contains(string(tags), "system_trust_known") {
		trust = "User"
	}
	if strings.Contains(string(tags), "system_trust_trusted") {
		trust = "Known User"
	}
	if strings.Contains(string(tags), "system_trust_veteran") {
		trust = "Trusted User"
	}
	if strings.Contains(string(tags), "system_trust_legend") {
		trust = "Veteran User"
	}
	if strings.Contains(string(tags), "system_legend") {
		trust = "Legendary User"
	}

	return trust
}
