package go_utils

import "testing"

func TestIsAlphabet(t *testing.T) {
	cases := []struct {
		in  byte
		exp bool
	}{
		{in: '0', exp: false},
		{in: '1', exp: false},
		{in: 'A', exp: true},
		{in: 'Z', exp: true},
		{in: 'a', exp: true},
		{in: 'z', exp: true},
		{in: ',', exp: false},
		{in: '?', exp: false},
		{in: ';', exp: false},
	}

	for _, v := range cases {
		if b := IsAlphabet(v.in); b != v.exp {
			t.Errorf("case %s not pass!", string(v.in))
		}
	}
}

func TestIsNumber(t *testing.T) {
	cases := []struct {
		in  byte
		exp bool
	}{
		{in: '1', exp: true},
		{in: '2', exp: true},
		{in: '0', exp: true},
		{in: 'a', exp: false},
		{in: 'Z', exp: false},
		{in: ',', exp: false},
		{in: '?', exp: false},
	}

	for _, v := range cases {
		if b := IsNumber(v.in); b != v.exp {
			t.Errorf("case %s not pass!", string(v.in))
		}
	}
}
