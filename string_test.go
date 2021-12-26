package go_utils

import "testing"

func TestStrListAddDel(t *testing.T) {
	str := "1,2,3,4"
	newStr := StrAddDel(str, "5", "3")
	if newStr != "1,2,4,5" {
		t.Fatalf("meet %s, but hope %s", newStr, "1,2,4,5")
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

func BenchmarkStrAddDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StrAddDel("1,2,3,4,5,6,7,8,9,3,4,5,6,,34,3,4,4,5,2,3,45,,5,2,3", "10", "0")
	}
}
