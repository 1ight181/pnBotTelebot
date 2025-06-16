package common

import (
	ctx "context"
	"fmt"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"regexp"
	"strings"
)

func IsSubscribed(userId int64, dbProvider dbifaces.DataBaseProvider) (bool, error) {
	user := dbmodels.User{
		TgId: userId,
	}
	if err := dbProvider.Find(ctx.Background(), &user, user); err != nil {
		return false, err
	}

	if user.IsSubscribed {
		return true, nil
	} else {
		return false, nil
	}
}

func EscapeMarkdownV2(text string) string {
	specialChars := []string{
		"_", "[", "]", "(", ")", "~", "`", "#", "+", "=", "|", "{", "}", ".", "!", "-",
	}

	for _, char := range specialChars {
		text = strings.ReplaceAll(text, char, "\\"+char)
	}

	return text
}

func WrapURLsWithPreviousWord(text string) string {
	urlRegex := regexp.MustCompile(`(https?://[^\s]+)`)

	matches := urlRegex.FindAllStringIndex(text, -1)

	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		start, end := match[0], match[1]
		url := text[start:end]

		wordStart := strings.LastIndexFunc(text[:start], func(r rune) bool {
			return !(r == ' ' || r == '\t' || r == '\n')
		})
		if wordStart == -1 {
			continue
		}

		spaceStart := strings.LastIndexFunc(text[:wordStart+1], func(r rune) bool {
			return r == ' ' || r == '\t' || r == '\n'
		}) + 1

		word := text[spaceStart:start]

		replacement := fmt.Sprintf("[%s](%s)", word, url)

		text = text[:spaceStart] + replacement + text[end:]
	}

	return text
}
