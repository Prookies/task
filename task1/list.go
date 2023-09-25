package main

import "fmt"

func Delete[T any](vals []T, idx int) []T {
	valsLen := len(vals)
	if idx < 0 || idx >= valsLen {
		panic("idx is invalid")
	}

	for i := idx; i < len(vals)-1; i++ {
		vals[i] = vals[i+1]
	}

	if valsLen-1 == cap(vals)/2 {
		newSlice := make([]T, valsLen - 1)
		copy(newSlice, vals[:valsLen-1])
		vals = newSlice
		return vals
	}

	return vals[:valsLen-1]
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s2 := Delete[int](s1, 1)
	fmt.Printf("s1: %p \n", s1)
	fmt.Printf("s1: %p \n", s2)

	fmt.Printf("s1: %v, len: %d, cpa: %d \n", s1, len(s1), cap(s1))

	fmt.Printf("s2: %v, len: %d, cpa: %d \n", s2, len(s2), cap(s2))

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("s3: %p \n", s3)
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	s3 = Delete[int](s3, 0)
	fmt.Printf("0 new s3: %p \n", s3)
	fmt.Printf("0 new s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	s3 = Delete[int](s3, 1)
	fmt.Printf("1 new s3: %p \n", s3)
	fmt.Printf("1 new s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	s3 = Delete[int](s3, 2)
	fmt.Printf("2 new s3: %p \n", s3)
	fmt.Printf("2 new s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	s3 = Delete[int](s3, 3)
	fmt.Printf("3 new s3: %p \n", s3)
	fmt.Printf("3 new s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))

}
