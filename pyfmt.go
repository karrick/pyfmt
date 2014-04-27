package pyfmt

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

func Sprintf(format string, a ...interface{}) string {
	if len(a) == 1 {
		switch a[0].(type) {
		case map[string]string:
			return form(format, a[0].(map[string]string))
		}
	}
	dict := make(map[string]string)
	for index, item := range a {
		dict[strconv.Itoa(index)] = fmt.Sprintf("%v", item)
	}
	return form(format, dict)
}

func form(format string, dict map[string]string) string {
	var buf, token bytes.Buffer
	var capture bool
	for _, rune := range format {
		switch capture {
		case false:
			if rune == '{' {
				capture = true
			} else {
				buf.WriteRune(rune)
			}
		case true:
			if unicode.IsLetter(rune) || unicode.IsDigit(rune) || rune == '_' {
				token.WriteRune(rune)
			} else {
				tok := token.String()
				val, ok := dict[tok]
				if rune == '}' && ok {
					buf.WriteString(val)
				} else {
					buf.WriteRune('{')
					buf.WriteString(tok)
					buf.WriteRune(rune)
				}
				token.Reset()
				capture = false // done with capture
			}
		}
	}
	if capture {
		buf.WriteRune('{')
		buf.Write(token.Bytes())
	}
	return buf.String()
}
