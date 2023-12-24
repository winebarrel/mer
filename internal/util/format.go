package util

import "strings"

func Comma(num string) string {
	afterPont := ""
	parts := strings.SplitN(num, ".", 2)

	if len(parts) > 1 {
		num = parts[0]
		afterPont = parts[1]
	}

	var wComma strings.Builder

	for i, r := range num {
		n := len(num) - i

		if i > 0 && n%3 == 0 {
			wComma.WriteString(",")
		}

		wComma.WriteRune(r)
	}

	if afterPont != "" {
		wComma.WriteString(".")
		wComma.WriteString(afterPont)
	}

	return wComma.String()
}
