package caseinsensitivecmp

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func lessInternal(a string, b string, allowEqual bool) bool {
	for {
		if len(b) == 0 {
			return allowEqual && len(a) == 0
		}
		if len(a) == 0 {
			return true
		}

		aRune, aRuneSize := utf8.DecodeRuneInString(a)
		bRune, bRuneSize := utf8.DecodeRuneInString(b)

		lowerARune := unicode.ToLower(aRune)
		lowerBRune := unicode.ToLower(bRune)

		if lowerARune < lowerBRune {
			return true
		}
		if lowerARune > lowerBRune {
			return false
		}

		a = a[aRuneSize:]
		b = b[bRuneSize:]
	}
}

func Equal(a string, b string) bool {
	return strings.EqualFold(a, b)
}

func Less(a string, b string) bool {
	return lessInternal(a, b, false)
}

func LessOrEqual(a string, b string) bool {
	return lessInternal(a, b, true)
}

func Greater(a string, b string) bool {
	return !lessInternal(a, b, true)
}

func GreaterOrEqual(a string, b string) bool {
	return !lessInternal(a, b, false)
}
