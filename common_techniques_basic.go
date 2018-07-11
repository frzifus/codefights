// This is free and unencumbered software released into the public domain.
//
// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.
//
// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// For more information, please refer to <http://unlicense.org/>

package codefights

// ContainsDuplicates returns true if the input slice contains two equal integers.
func ContainsDuplicates(a []int) bool {
	l := len(a)
	res := make(map[int]bool)
	for i := 0; i < l; i++ {
		res[a[i]] = true
	}
	return len(res) < l
}

// SumOfTwo returns true if the sum of the two slices can be formed.
func SumOfTwo(a []int, b []int, v int) bool {
	p := make(map[int]bool)
	for j := 0; j < len(b); j++ {
		p[b[j]] = true
	}

	for i := 0; i < len(a); i++ {
		c := v - a[i]
		if v, ok := p[c]; ok {
			return v
		}
	}
	return false
}

// SumInRange accepts an slice of integers and an slice of indice pairs.
// The total is offset against mudolo 10^9 + 7 and returned.
func SumInRange(nums []int, queries [][]int) int {
	const mod = 1000000007
	var sum int
	var part int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		part += nums[i]
		m[i] = part % mod
	}
	for i := 0; i < len(queries); i++ {
		if queries[i][0] == 0 {
			sum = sum%mod + m[queries[i][1]]%mod
		} else {
			sum = sum%mod + m[queries[i][1]] - m[queries[i][0]-1]%mod
		}
	}
	return (sum + mod) % mod
}

// ArrayMaxConsecutiveSum2 returns the maximum possible sum you can get from one
// of its contiguous subarrays.
func ArrayMaxConsecutiveSum2(inputArray []int) int {
	max := inputArray[0]
	end := inputArray[0]
	for _, x := range inputArray[1:] {
		if x < end+x {
			end = end + x
		} else {
			end = x
		}
		if max < end {
			max = end
		}
	}
	return max
}

// FindLongestSubarrayBySum returns two integers that represent its inclusive
// bounds. If there are several possible answers, return the one with the
// smallest left bound. If there are no answers, return [-1]
func FindLongestSubarrayBySum(s int, arr []int) []int {
	tmp := make(map[int]int)
	res := []int{-1, -1}
	var begin, sum, max int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		for sum > s {
			sum -= arr[begin]
			begin++
		}
		if sum == s {
			tmp[begin+1] = i + 1
		}
	}
	if len(tmp) == 0 {
		return []int{-1}
	}
	for k, v := range tmp {
		diff := v - k
		if max <= diff {
			max = diff
			res[0] = k
			res[1] = v
		}
	}
	return res
}

// ProductExceptSelf is not yet implemented.
func ProductExceptSelf(nums []int, m int) int {
	// TODO
	return 0
}

// MinSubstringWithAllChars is not yet implemented.
func MinSubstringWithAllChars(s string, t string) string {
	// TODO
	return ""
}
