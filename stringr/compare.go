/**
 * @Author: liubaoshuai3
 * @Date: 2020/5/13 10:47
 * @File: compare
 * @Description:
 */

package stringr

// compare two string
func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}