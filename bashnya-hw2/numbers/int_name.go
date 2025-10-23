package numbers

import (
	"slices"
	"strings"
)

var partWords = map[int]string{
	0:   "",
	1:   "один",
	2:   "два",
	3:   "три",
	4:   "четыре",
	5:   "пять",
	6:   "шесть",
	7:   "семь",
	8:   "восемь",
	9:   "девять",
	10:  "десять",
	11:  "одиннадцать",
	12:  "двенадцать",
	13:  "тринадцать",
	14:  "четырнадцать",
	15:  "пятнадцать",
	16:  "шестнадцать",
	17:  "семнадцать",
	18:  "восемнадцать",
	19:  "девятнадцать",
	20:  "двадцать",
	30:  "тридцать",
	40:  "сорок",
	50:  "пятьдесят",
	60:  "шестьдесят",
	70:  "семьдесят",
	80:  "восемьдесят",
	90:  "девяносто",
	100: "сто",
	200: "двести",
	300: "триста",
	400: "четыреста",
	500: "пятьсот",
	600: "шестьсот",
	700: "семьсот",
	800: "восемьсот",
	900: "девятьсот",
}

var intWords = [...]string{
	"",
	"тысяч",
	"миллион",
	"миллиард",
	"триллион",
	"квадриллион",
	"квинтиллион",
}

func getLastWord(part int, isThousands bool) string {
	var lastWord string
	if !isThousands {
		lastWord = partWords[part]
	} else {
		switch part {
		case 1:
			lastWord = "одна"
		case 2:
			lastWord = "две"
		default:
			lastWord = partWords[part]
		}
	}
	return lastWord
}

func numPartToWords(part int, isThousands bool) string {
	var wordRepr strings.Builder
	hundreds := (part / 100) * 100
	wordRepr.WriteString(partWords[hundreds])
	part -= hundreds
	if part == 0 {
		return wordRepr.String()
	}
	if wordRepr.Len() > 0 {
		wordRepr.WriteByte(' ')
	}
	if part <= 19 {
		lastWord := getLastWord(part, isThousands)
		wordRepr.WriteString(lastWord)
	} else {
		tens := (part / 10) * 10
		wordRepr.WriteString(partWords[tens])
		part -= tens
		if part != 0 {
			wordRepr.WriteByte(' ')
			lastWord := getLastWord(part, isThousands)
			wordRepr.WriteString(lastWord)
		}
	}
	return wordRepr.String()
}

func getDefaultSuffix(part int) string {
	tens := part - (part/100)*100
	if tens/10 == 1 {
		return "ов"
	}
	lastDigit := part % 10
	switch lastDigit {
	case 2, 3, 4:
		return "а"
	case 1:
		return ""
	default:
		return "ов"
	}
}

func getThousandsSuffix(part int) string {
	tens := part - (part/100)*100
	if tens/10 == 1 {
		return ""
	}
	lastDigit := part % 10
	switch lastDigit {
	case 2, 3, 4:
		return "и"
	case 1:
		return "а"
	default:
		return ""
	}
}

func IntToWords(num int) string {
	if num == 0 {
		return "ноль"
	}
	var wordRepr strings.Builder
	if num < 0 {
		wordRepr.WriteString("минус ")
		num = -num
	}
	parts := make([]int, 0)
	for num > 0 {
		parts = append(parts, num%1000)
		num /= 1000
	}
	slices.Reverse(parts)
	partsCount := len(parts)
	var isThousands bool
	for i, part := range parts {
		if part == 0 {
			continue
		}
		if i == len(parts)-2 {
			isThousands = true
		} else {
			isThousands = false
		}
		wordRepr.WriteString(numPartToWords(part, isThousands))
		if i == len(parts)-1 {
			break
		}
		wordRepr.WriteByte(' ')
		wordRepr.WriteString(intWords[partsCount-i-1])
		if !isThousands {
			wordRepr.WriteString(getDefaultSuffix(part))
		} else {
			wordRepr.WriteString(getThousandsSuffix(part))
		}
		wordRepr.WriteByte(' ')
	}
	return strings.TrimSpace(wordRepr.String())
}
