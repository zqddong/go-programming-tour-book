package logic

import (
	"github.com/zqddong/go-programming-tour-book/chatroom/gloable"
	"strings"
)

// FilterSensitive  敏感词过滤
func FilterSensitive(content string) string {
	for _, word := range gloable.SensitiveWords {
		content = strings.ReplaceAll(content, word, "**")
	}

	return content
}
