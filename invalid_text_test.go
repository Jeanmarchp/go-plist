package plist

import (
	"strings"
	"testing"
)

var InvalidTextPlists = []struct {
	Name string
	Data string
}{
	{"Truncated array", "("},
	{"Truncated dictionary", "{a=b;"},
	{"Truncated dictionary 2", "{"},
	{"Unclosed nested array", "{0=(/"},
	{"Unclosed dictionary", "{0=/"},
	{"Broken GNUStep data", "(<*I5>,<*I5>,<*I5>,<*I5>,*I16777215>,<*I268435455>,<*I4294967295>,<*I18446744073709551615>,)"},
	{"Truncated nested array", "{0=(((/"},
	{"Truncated dictionary with comment-like", "{/"},
	{"Truncated array with comment-like", "(/"},
	{"Truncated array with empty data", "(<>"},
	{"Bad Extended Character", "{¬=A;}"},
	{"Missing Equals in Dictionary", `{"A"A;}`},
	{"Missing Semicolon in Dictionary", `{"A"=A}`},
	{"Invalid GNUStep type", "<*F33>"},
	{"Invalid GNUStep int", "(<*I>"},
	{"Invalid GNUStep date", "<*D5>"},
	{"Invalid data", "<EQ>"},
	{"Truncated unicode escape", `"\u231`},
	{"Truncated hex escape", `"\x2`},
	{"Truncated octal escape", `"\02`},
	{"Truncated data", `<33`},
	{"Truncated block comment", `/* hello`},
	{"Truncated quoted string", `"hi`},
}

func TestInvalidTextPlists(t *testing.T) {
	for _, test := range InvalidTextPlists {
		t.Run(test.Name, func(t *testing.T) {
			var obj interface{}
			buf := strings.NewReader(test.Data)
			err := NewDecoder(buf).Decode(&obj)
			if err == nil {
				t.Fatal("invalid plist failed to throw error")
			} else {
				t.Log(err)
			}
		})
	}
}
