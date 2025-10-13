package parsinglogfiles

import (
    "fmt"
    "regexp"
    "strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    indexes := re.FindStringIndex(text)
    return indexes != nil && indexes[0] == 0 && indexes[1] == 5
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=\-]*>`)
    sl := re.Split(text, -1)
    fmt.Println(sl)
    return sl
}

func CountQuotedPasswords(lines []string) int {
    count := 0
	re := regexp.MustCompile(`"([^"]*)"`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			insideQuotes := matches[1]
			if strings.Contains(strings.ToLower(insideQuotes), "password") {
				count++
			}
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	var result []string
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			username := matches[1]
			line = fmt.Sprintf("[USR] %s %s", username, line)
		}
		result = append(result, line)
	}
	return result
}
