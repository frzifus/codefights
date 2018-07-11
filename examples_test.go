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

package codefights_test

import (
	"fmt"
	"git.klimlive.de/frzifus/codefights"
)

func ExampleContainsDuplicates() {
	input := []int{1, 2, 3, 1}
	res := codefights.ContainsDuplicates(input)
	fmt.Printf("codefights.ContainsDuplicates(input) = %v\n", res)
	// Output:
	// codefights.ContainsDuplicates(input) = true
}

func ExampleSumOfTwo() {
	a := []int{1, 2, 3}
	b := []int{10, 20, 30, 40}
	v := 42
	res := codefights.SumOfTwo(a, b, v)
	fmt.Printf("codefights.SumOfTwo(a, b, v) = %v\n", res)
	// Output:
	// codefights.SumOfTwo(a, b, v) = true
}

func ExampleSumInRange() {
	nums := []int{3, 0, -2, 6, -3, 2}
	queries := [][]int{
		{0, 2},
		{2, 5},
		{0, 5},
	}
	res := codefights.SumInRange(nums, queries)
	fmt.Printf("codefights.SumInRange(nums, queries) = %v\n", res)
	// Output:
	// codefights.SumInRange(nums, queries) = 10
}

func ExampleArrayMaxConsecutiveSum2() {
	input := []int{-2, 2, 5, -11, 6}
	res := codefights.ArrayMaxConsecutiveSum2(input)
	fmt.Printf("codefights.ArrayMaxConsecutiveSum2(input) = %v\n", res)
	// Output:
	// codefights.ArrayMaxConsecutiveSum2(input) = 7
}

func ExampleFindLongestSubarrayBySum() {
	s := 12
	arr := []int{1, 2, 3, 7, 5}
	res := codefights.FindLongestSubarrayBySum(s, arr)
	fmt.Printf("codefights.FindLongestSubarrayBySum(s, arr) = %v\n", res)
	// Output:
	// codefights.FindLongestSubarrayBySum(s, arr) = [2 4]

}

func ExampleProductExceptSelf() {
	// coming sooon
}

func ExampleMinSubstringWithAllChars() {
	// coming sooon
}
