package palindrome

import (
	"sort"
	"testing"
)

func TestPalindrome_1(t *testing.T) {
	ints := []int{7, 7, 0, 8, 0, 7, 7}
	if !IsPalindrome(sort.IntSlice(ints)) {
		t.Fail()
	}
}

func TestPalindrome_2(t *testing.T) {
	ints := []int{7, 7, 0, 8}
	if IsPalindrome(sort.IntSlice(ints)) {
		t.Fail()
	}
}

func TestPalindrome_3(t *testing.T) {
	r := []rune("level")
	if !IsPalindrome(sortRunes(r)) {
		t.Fail()
	}
}

func TestPalindrome_4(t *testing.T) {
	r := []rune("hanna")
	if IsPalindrome(sortRunes(r)) {
		t.Fail()
	}
}
