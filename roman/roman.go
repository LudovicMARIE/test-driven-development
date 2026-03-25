package roman

import "strings"

func ConvertIntToRomanNumeral(num int) string {

	romanValues := []struct {
		value  int
		symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var result strings.Builder

	for _, rv := range romanValues {
		for num >= rv.value {
			result.WriteString(rv.symbol)
			num -= rv.value
		}
	}

	return result.String()
}
