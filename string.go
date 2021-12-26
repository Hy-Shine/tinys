package go_utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// StrAddDel return string
// example: StrAddDel("1,2,3", "2", "4") --> "1,3,4"
func StrAddDel(old, add, del string) string {
	// add the `add` string
	if add != "" && !strings.Contains(old, add) {
		old += "," + add
	}
	// init a new slice
	s := make([]string, 0)
	for _, v := range strings.Split(old, ",") {
		if v != "" && v != del {
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
// example: IntStringSort("4,57,8") --> "4,8,57"
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
// example: Float64StringSort("4.2,5.7,8.1,2") --> "2.0,4.2,5.7,8.1"
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

func StrToDistinctStr(str string) string {
	strList := StrToList(str)
	strMap := make(map[string]int)
	for _, v := range strList {
		strMap[v] = 0
	}

	var builder strings.Builder
	for k := range strMap {
		builder.WriteString(k)
	}

	return builder.String()
}

func StrToDistinctList(str string) []string {
	list := StrToDistinctStr(str)
	return strings.Split(list, ",")
}
