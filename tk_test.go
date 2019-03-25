package tk

import (
	"fmt"
	"testing"
)

func TestTK(t *testing.T) {
	tkk, _ := GetTKK()
	tk := GetTK("hello", tkk)
	fmt.Println(tk)
}
