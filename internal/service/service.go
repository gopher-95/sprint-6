package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseOrWord(str string) string {
	// проверяем, содержит ли строка символы ".-", если да - конвертируем в слово
	if strings.ContainsAny(str, ".-") {
		word := morse.ToText(str)
		return word
	}

	// в обратном случае - конвертируем в Морзе
	morse := morse.ToMorse(str)
	return morse
}
