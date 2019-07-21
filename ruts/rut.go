package ruts

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	ErrInvalidRut = errors.New("invalid rut")
)

func Clean(rut string) string {
	rut = strings.Replace(rut, ".", "", -1)
	rut = strings.Replace(rut, ",", "", -1)
	rut = strings.Replace(rut, "-", "", -1)
	return rut
}

func Parse(rut string) (int, error) {
	if rut == "" || len(rut) < 1 || len(rut) > 12 {
		return 0, ErrInvalidRut
	}

	rut = Clean(rut)

	s := []rune(rut)
	s = s[:len(s)-1]
	r, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, err
	}
	return r, nil
}

func GetDV(rut int) rune {
	t := rut
	count := 0
	index := 2
	for ; t != 0; t /= 10 {
		m := t % 10
		count += m * index
		index++
		if index > 7 {
			index = 2
		}
	}
	r := count % 11
	d := 11 - r
	if d == 10 {
		return 'K'
	}
	return rune(d + 48)
}

func Validate(rut string) bool {
	parsed, err := Parse(rut)
	if err != nil {
		return false
	}

	calculateDV := GetDV(parsed)
	cleaned := Clean(rut)
	cleanedRune := []rune(cleaned)
	providedDV := cleanedRune[len(cleanedRune)-1]

	if providedDV == 'k' {
		providedDV = 'K'
	}

	return calculateDV == providedDV
}

func Format(rut int) string {
	formatted := message.NewPrinter(language.Spanish).Sprintf("%d", rut)
	dv := GetDV(rut)
	return fmt.Sprintf("%s-%s", formatted, string(dv))
}
