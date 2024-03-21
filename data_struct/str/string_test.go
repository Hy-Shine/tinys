package str

import (
	"testing"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Empty slice",
			input:    []string{},
			expected: "",
		},
		{
			name:     "Single string",
			input:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "Multiple strings",
			input:    []string{"Hello", "world", "!", "!", "!"},
			expected: "Helloworld!!!",
		},
		{
			name:     "Strings with different lengths",
			input:    []string{"hello", "world", "!", "this", "is", "a", "test"},
			expected: "helloworld!thisisatest",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Concat(test.input, "")
			if result != test.expected {
				t.Errorf("Expected %q, but got %q", test.expected, result)
			}
		})
	}
}

func BenchmarkStrConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Concat([]string{"hello", "world", "!"}, "")
	}
}

func TestToNumber(t *testing.T) {
	// Testing parsing of a valid integer string
	intResult, err := ToNumber[int]("123")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if intResult != 123 {
		t.Errorf("Expected result to be 123, but got %v", intResult)
	}

	// Testing parsing of a valid float string
	floatResult, err := ToNumber[float64]("12.34")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if floatResult != 12.34 {
		t.Errorf("Expected result to be 12.34, but got %v", floatResult)
	}

	// Testing parsing of an invalid string
	_, err = ToNumber[int]("abc")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestStrListAddDel(t *testing.T) {
	str := "1,2,3,4"
	newStr := StrAddDel(str, "5", "3")
	if newStr != "1,2,4,5" {
		t.Fatalf("meet %s, but expect %s", newStr, "1,2,4,5")
	}
	newStr = StrAddDel(str, "2", "3")
	if newStr != "1,2,4" {
		t.Fatalf("meet %s, but hope %s", newStr, "1,2,4")
	}
}

func TestIntStringSort(t *testing.T) {
	sortInt := IntStringSort("1,2,3,3,4,5,6,7,2,4,0,1,3")
	expect := "0,1,1,2,2,3,3,3,4,4,5,6,7"
	if sortInt != expect {
		t.Fatalf("meet %s, but expect %s", sortInt, expect)
	}
	sortInt = IntStringSort(",,0,23,-1,4,7,9")
	expect = "-1,0,4,7,9,23"
	if sortInt != expect {
		t.Fatalf("meet %s, but expect %s", sortInt, expect)
	}
}

func TestToUpper(t *testing.T) {
	s := "test lower"
	ToUpper(&s)
	if s != "TEST LOWER" {
		t.Fatalf("meet %s, but expect %s", s, "TEST LOWER")
	}
}

func BenchmarkToUpper(b *testing.B) {
	s := "Hello World"
	for i := 0; i < b.N; i++ {
		ToUpper(&s)
	}
}

func TestToLower(t *testing.T) {
	s := "TEST lower"
	ToLower(&s)
	if s != "test lower" {
		t.Fatalf("meet %s, but expect %s", s, "test lower")
	}
}

func BenchmarkStrAddDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StrAddDel("1,2,3,4,5,6,7,8,9,3,4,5,6,,34,3,4,4,5,2,3,45,,5,2,3", "10", "0")
	}
}
