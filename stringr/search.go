/**
 * @Author: liubaoshuai3
 * @Date: 2020/5/13 11:06
 * @File: search
 * @Description:
 */

package stringr

// good suffix and bad what what...
type stringFinder struct {
	pattern string
	badCharSkip [256]int
	goodSuffixSkip []int
}

// TODO search is too hard for me to understand
func makeStringFinder(pattern string) *stringFinder {
	f := &stringFinder{
		pattern:        pattern,
		goodSuffixSkip: make([]int, 0, len(pattern)),
	}
	return f
}