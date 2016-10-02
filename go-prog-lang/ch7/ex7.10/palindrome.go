package palindrome

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Len() int           { return len(s) }

func equals(s sort.Interface, i, j int) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func IsPalindrome(s sort.Interface) bool {
	n := s.Len()
	for i := range make([]int, n/2) {
		if !equals(s, i, n-i-1) {
			return false
		}
	}

	return true
}
