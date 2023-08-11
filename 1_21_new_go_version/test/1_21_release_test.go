package test

import (
	"fmt"
	"slices"
	"testing"
)

func Test_Min(t *testing.T) {

	a := 1
	b := 10
	i := min(a, b)
	fmt.Println(i)
	fmt.Println(max(a, b))
	slices.Clip()
}
