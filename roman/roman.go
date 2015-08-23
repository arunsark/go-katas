package roman

import "strings"

type RomanAtom struct {
    DecimalValue uint16
    Rune string
}

func DivMod(number uint16, divisor uint16) (quotient, reminder uint16) {
    quotient = number / divisor
    reminder = number % divisor
    return
}

func Roman(num uint16) (roman string) {
    var mapping = []RomanAtom {
        {DecimalValue: 1000, Rune: "M"},
        {DecimalValue: 900, Rune: "CM"},
        {DecimalValue: 500, Rune: "D"},
        {DecimalValue: 400, Rune: "CD"},
        {DecimalValue: 100, Rune: "C"},
        {DecimalValue: 90, Rune: "XC"},
        {DecimalValue: 50, Rune: "L"},
        {DecimalValue: 40, Rune: "XL"},
        {DecimalValue: 10, Rune: "X"},
        {DecimalValue: 9, Rune: "IX"},
        {DecimalValue: 5, Rune: "V"},
        {DecimalValue: 4, Rune: "IV"},
        {DecimalValue: 1, Rune: "I"},
    }


    for _, romanAtom := range mapping {
        amount_of_runes, remaining_decimal := DivMod(num, romanAtom.DecimalValue)
        num = remaining_decimal
        roman += strings.Repeat(romanAtom.Rune, int(amount_of_runes))
    }
    return roman
}
