package rut

import (
    "fmt"
    "strconv"
    "strings"
)

func Parse(rut string) (int, error) {
    if rut == "" || len(rut) < 1 || len(rut) > 12 {
        return 0, fmt.Errorf("invalid rut %s", rut)
    }

    rut = strings.Replace(rut, ".", "", -1)
    rut = strings.Replace(rut, ",", "", -1)
    rut = strings.Replace(rut, "-", "", -1)

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
