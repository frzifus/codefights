package codefights

import (
	"math/rand"
	"testing"
)

func TestContainsDuplicates(t *testing.T) {
	tt := []struct {
		name   string
		a      []int
		expect bool
	}{
		{"test1", []int{1, 2, 3, 1}, true},
		{"test2", []int{3, 1}, false},
		{"test3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{"test4", []int{}, false},
		{"test5", []int{1}, false},
		{"test6", []int{-1200000005, -1200000005}, true},
		{"test7", []int{0, 4, 5, 0, 3, 6}, true},
		{"test8", []int{1, 2, 3, 4}, false},
		{"test9", []int{0, 1, 0, -1}, true},
		{"test10", []int{2}, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsDuplicates(tc.a)
			if res != tc.expect {
				t.Errorf("%s", "")
			}
		})
	}
}

func BenchmarkContainsDuplicates(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rnd := rand.Perm(100)
		b.StartTimer()
		ContainsDuplicates(rnd)
		b.StopTimer()
	}
}

func TestSumOfTwo(t *testing.T) {
	tt := []struct {
		name   string
		a      []int
		b      []int
		v      int
		expect bool
	}{
		{"test1", []int{1, 2, 3}, []int{10, 20, 30, 40}, 42, true},
		{"test2", []int{1, 2, 3}, []int{10, 20, 30, 40}, 50, false},
		{"test3", []int{}, []int{1, 2, 3, 4}, 4, false},
		{"test4", []int{10, 1, 5, 3, 8}, []int{100, 6, 3, 1, 5}, 4, true},
		{"test5", []int{1, 4, 3, 6, 10, 1, 0, 1, 6, 5}, []int{9, 5, 6, 9, 0, 1, 2, 1, 6, 10}, 8, true},
		{"test6", []int{3, 2, 3, 7, 5, 0, 3, 0, 4, 2}, []int{6, 8, 2, 9, 7, 10, 3, 8, 6, 0}, 2, true},
		{"test7", []int{4, 6, 4, 2, 9, 6, 6, 2, 9, 2}, []int{3, 4, 5, 1, 4, 10, 9, 9, 6, 4}, 5, true},
		{"test8", []int{6, 10, 25, 13, 20, 21, 11, 10, 18, 21}, []int{21, 10, 6, 0, 29, 25, 1, 17, 19, 25}, 37, true},
		{"test9", []int{22, 26, 6, 22, 17, 11, 9, 22, 7, 12}, []int{14, 25, 22, 27, 22, 30, 6, 26, 30, 27}, 56, true},
		{"test10", []int{17, 72, 18, 72, 73, 15, 83, 90, 8, 18}, []int{100, 27, 33, 51, 2, 71, 76, 19, 16, 43}, 37, true},
		{"test11", []int{75, 38, 10, 57, 67, 39, 26, 14, 53, 80}, []int{3, 19, 28, 92, 92, 47, 98, 30, 71, 21}, 61, true},
		{"test12", []int{1, 2, 3}, []int{}, 1, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := SumOfTwo(tc.a, tc.b, tc.v)
			if res != tc.expect {
				t.Errorf("didn't find a pair like expected (%v) + (%v) >> %v",
					tc.a, tc.b, tc.v)
			}
		})
	}
}

func BenchmarkSumOfTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		x := rand.Perm(100)
		y := rand.Perm(100)
		z := 50
		b.StartTimer()
		SumOfTwo(x, y, z)
		b.StopTimer()
	}
}

func TestSumInRange(t *testing.T) {
	tt := []struct {
		name    string
		nums    []int
		queries [][]int
		expect  int
	}{
		{"test1",
			[]int{3, 0, -2, 6, -3, 2},
			[][]int{
				{0, 2},
				{2, 5},
				{0, 5},
			},
			10,
		},
		{"test2",
			[]int{-1000},
			[][]int{
				{0, 0},
			},
			999999007,
		},
		{"test3",
			[]int{34, 19, 21, 5, 1, 10, 26, 46, 33, 10},
			[][]int{
				{3, 7},
				{3, 4},
				{3, 7},
				{4, 5},
				{0, 5},
			},
			283,
		},
		{"test4",
			[]int{-4, -18, -22, -14, -33, -47, -29, -35, -50, -19},
			[][]int{
				{2, 9},
				{5, 6},
				{1, 2},
				{2, 2},
				{4, 5},
			},
			999999540,
		},
		{"test5",
			[]int{
				-23, -8, -52, -58, 93, -16, -26, 75, -77, 25, 90, -50, -31,
				70, 53, -68, 96, 100, 69, 13,
			},
			[][]int{
				{0, 4},
				{0, 8},
				{7, 7},
				{3, 4},
				{2, 3},
				{0, 3},
				{8, 8},
				{2, 2},
				{5, 7},
				{2, 2},
			},
			999999578,
		},
		{"test6",
			[]int{-77, 54, -59, -94, -13, -78, -81, -38, -26, 17, -73, -88, 90,
				-42, -63, -36, 37, 25, -22, 4, 25, -86, -44, 88, 2, -47, -29,
				71, 54, -42},
			[][]int{
				{2, 2},
				{4, 7},
				{2, 4},
				{0, 2},
				{3, 6},
				{6, 6},
				{3, 3},
				{2, 7},
				{3, 4},
				{3, 3},
				{2, 9},
				{0, 1},
				{4, 4},
				{2, 3},
				{0, 6},
				{4, 4},
				{2, 3},
				{0, 5},
				{2, 5},
				{4, 5},
			},
			999996808,
		},
		{"test7",
			[]int{
				1000,
			},
			[][]int{
				{0, 0},
			},
			1000,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := SumInRange(tc.nums, tc.queries)
			if res != tc.expect {
				t.Errorf("Want: %v, Got: %v", tc.expect, res)
			}
		})
	}
}

func BenchmarkSumInRange(b *testing.B) {
	var queries [10][]int
	for n := 0; n < b.N; n++ {
		nums := rand.Perm(100)
		for i := 0; i < 10; i++ {
			queries[i] = rand.Perm(2)
			// queries = append(queries, append(queries[i], rand.Perm(2)...))
		}
		b.StartTimer()
		SumInRange(nums, queries[:])
		b.StopTimer()
	}
}

// func TestProductExceptSelf(t *testing.T) {
//	tt := []struct {
//		name   string
//		nums   []int
//		m      int
//		expect int
//	}{
//		{"test1", []int{1, 2, 3, 4}, 12, 2},
//		{"test2", []int{2, 100}, 24, 6},
//		{"test3", []int{5, 8, 6, 3, 2}, 8, 4},
//		{"test4", []int{3, 3, 9, 5, 5, 4, 2, 8, 5, 9}, 17, 16},
//		{"test5", []int{
//			27, 37, 47, 30, 17, 6, 20, 17, 21, 43, 5, 49, 49, 50, 20, 42, 45,
//			1, 22, 44,
//		},
//			12,
//			2,
//		},
//		{"test6", []int{28, 27, 11, 17, 19, 49, 19, 46, 41, 21, 1, 49, 18, 26,
//			16, 24, 16, 36, 19, 49, 31, 39, 11, 21, 29, 37, 34, 34, 6, 16, 26,
//			31, 6, 48, 38, 36, 26, 36, 38, 18,
//		},
//			5040,
//			0,
//		},
//		{"test7", []int{52, 40, 2, 78, 49, 70, 39, 26, 58, 58, 52, 93, 80, 64,
//			33, 72, 29, 17, 81, 83, 48, 9, 49, 82, 67, 76, 54, 64, 6, 48, 16,
//			82, 67, 56, 32, 98, 14, 47, 48, 26, 56, 54, 80, 13, 32, 18, 4, 73,
//			45, 65,
//		},
//			530,
//			220,
//		},
//		{"test8", []int{37, 50, 50, 6, 8, 54, 7, 64, 2, 64, 67, 59, 22, 73, 33,
//			53, 43, 77, 56, 76, 59, 96, 64, 100, 89, 38, 64, 73, 85, 96, 65, 50,
//			62, 4, 82, 57, 98, 61, 92, 55, 80, 53, 80, 55, 94, 9, 73, 89, 83,
//			80,
//		},
//			67,
//			55,
//		},
//	}

//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			res := ProductExceptSelf(tc.nums, tc.m)
//			if res != tc.expect {
//				t.Errorf("Want: %v, Got: %v", tc.expect, res)
//			}
//		})
//	}
// }

// func TestMinSubstringWithAllChars(t *testing.T) {
//	tt := []struct {
//		name   string
//		s      string
//		t      string
//		expect string
//	}{
//		{"test1", "adobecodebanc", "abc", "banc"},
//		{"test2", "", "", ""},
//		{"test3", "abz", "abz", "abz"},
//		{"test4", "zqyvbfeiee", "ze", "zqyvbfe"},
//		{"test5", "tvdsxcqsnoeccaurocnk", "acpt", "tvdsxcqsnoecca"},
//		{"test6", "xgajymplpvftjwjqomhbnutorgysaf", "j", "j"},
//	}

//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			res := MinSubstringWithAllChars(tc.s, tc.t)
//			if res != tc.expect {
//				t.Errorf("Want: %v, Got: %v", tc.expect, res)
//			}
//		})
//	}
// }
