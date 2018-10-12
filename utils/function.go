package utils

import (
	"regexp"

	"github.com/satori/go.uuid"
)

// GetUUID returns uuid
func GetUUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

// ReplaceFileSuffix
func ReplaceFileSuffix(s string) string {
	re, _ := regexp.Compile(".(jpg|jpeg|png|gif|exe|doc|docx|ppt|pptx|xls|xlsx)")
	suffix := re.ReplaceAllString(s, "")
	return suffix
}
