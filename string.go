package go_utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// StrListAddDel return string
// example: StrListAddDel("1,2,3", "2", "4") --> "1,3,4"
func StrListAddDel(old, add, del string) string {
	// delete the `del` string
	if del != "" {
		old = strings.ReplaceAll(old, del, "")
	}
	// add the `add` string
	if add != "" && strings.Contains(old, add) {
		add += "," + add
	}
	// init a new slice
	s := make([]string, 0)
	for _, v := range strings.Split(old, ",") {
		if v != "" {
			s = append(s, v)
		}
	}
	return strings.Join(s, ",")
}

// StrToList return a list of string.
// example: 1,2,3 --> []string{"1", "2", "3"}
func StrToList(str string) []string {
	cnt := strings.Count(str, ",")
	s := make([]string, 0, cnt+1)
	for _, v := range strings.Split(str, ",") {
		if v != "" {
			s = append(s, v)
		}
	}
	return s
}

// IntStringSort return a sorted string
// example: "4,57,8" --> "4,8,57"
func IntStringSort(str string) string {
	strList := StrToList(str)
	intList := make([]int, 0)
	for _, v := range strList {
		n, _ := strconv.Atoi(v)
		intList = append(intList, n)
	}

	sort.Ints(intList)
	var builder strings.Builder
	for _, v := range intList {
		s := strconv.Itoa(v)
		builder.WriteString(s + ",")
	}
	strSorted := builder.String()
	length := len(strSorted)
	return strSorted[:length-1]
}

// Float64StringSort return a sorted string
// example: "4.2,5.7,8.1,2" --> "4,8,57"
func Float64StringSort(str string) string {
	strList := StrToList(str)
	float64List := make([]float64, 0)
	for _, v := range strList {
		f, _ := strconv.ParseFloat(v, 64)
		float64List = append(float64List, f)
	}

	sort.Float64sAreSorted(float64List)
	var builder strings.Builder
	for _, v := range float64List {
		s := fmt.Sprintf("%f,", v)
		builder.WriteString(s)
	}
	strSorted := builder.String()
	length := len(strSorted)
	return strSorted[:length-1]
}
