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
	var buf, tok bytes.Buffer
	var capture bool
	for _, rune := range format {
		switch capture {
		case true:
			if unicode.IsDigit(rune) || unicode.IsLetter(rune) || rune == '_' {
				tok.WriteRune(rune)
			} else {
				token := tok.String()
				val, ok := dict[token]
				if rune == '}' && ok {
					buf.WriteString(val)
				} else {
					buf.WriteRune('{')
					buf.WriteString(token)
					buf.WriteRune(rune)
				}
				tok.Reset()
				capture = false // done with capture
			}
		case false:
			if rune == '{' {
				capture = true
			} else {
				buf.WriteRune(rune)
			}
		}
	}
	if capture {
		buf.WriteRune('{')
		buf.Write(tok.Bytes())
	}
	return buf.String()
}
