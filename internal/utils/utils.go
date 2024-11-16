package utils

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"regexp"
	"strings"
)

func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

type Number interface {
	constraints.Integer | constraints.Float
}

func SliceConversion[T1, T2 Number, S1 ~[]T1, S2 []T2](s1 S1) S2 {
	size := len(s1)
	s2 := make(S2, size)
	for i := 0; i < size; i++ {
		s2[i] = T2(s1[i])
	}
	return s2
}

func RemoveTitle(text string) string {
	res := ""
	for _, v := range strings.Split(text, "\n") {
		if !regexp.MustCompile("^[#]{1,6} .*").MatchString(v) {
			res = fmt.Sprintf("%s%s", res, v)
		}
	}
	return res
}
