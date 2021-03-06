/*
给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：
	A.length >= 3
	在 0 < i < A.length - 1 条件下，存在 i 使得：

		A[0] < A[1] < ... A[i-1] < A[i]
		A[i] > A[i+1] > ... > A[A.length - 1]
。
*/
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	s, e := 0, len(A)-1
	for s < e && A[s] < A[s+1] {
		s++
	}
	for e > s && A[e] < A[e-1] {
		e--
	}
	return s == e && s != 0 && e != len(A)-1
}