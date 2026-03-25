package roman

import "testing"

func TestConvertIntToRomanNumeral(t *testing.T) {
	testCase := []struct {
		name   string
		value  int
		result string
	}{
		{"1", 1, "I"},
		{"4", 4, "IV"},
		{"7", 7, "VII"},
		{"10", 10, "X"},
		{"9", 9, "IX"},
		{"40", 40, "XL"},
		{"90", 90, "XC"},
		{"400", 400, "CD"},
		{"900", 900, "CM"},
		{"1879", 1879, "MDCCCLXXIX"},
		{"3999", 3999, "MMMCMXCIX"},
		{"368", 368, "CCCLXVIII"},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := ConvertIntToRomanNumeral(tc.value)
			if result != tc.result {
				t.Errorf("Result was incorrect, got: %s, want: %s.", result, tc.result)
			}
		})
	}

}
