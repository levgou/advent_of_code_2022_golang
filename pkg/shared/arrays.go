package shared

import (
	"errors"
	"golang.org/x/exp/constraints"
	"sort"
)

func GroupLines(input []string) [][]string {
	result := [][]string{}
	accumulator := []string{}

	for _, line := range input {
		if line == "" {
			if len(accumulator) > 0 {
				result = append(result, accumulator)
				accumulator = []string{}
			}
		} else {
			accumulator = append(accumulator, line)
		}
	}

	if len(accumulator) > 0 {
		result = append(result, accumulator)
	}

	return result
}

func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func Max[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func MaxN[T constraints.Ordered](slice []T, n int) ([]T, error) {
	if len(slice) == 0 {
		return make([]T, 0), errors.New("empty slice")
	}

	if len(slice) < n {
		return make([]T, 0), errors.New("array too small")
	}

	res := make([]T, len(slice))
	copy(res, slice)
	sortSlice(res)

	return res[len(res)-n : len(res)], nil
}

func Min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

func MinN[T constraints.Ordered](slice []T, n int) ([]T, error) {
	if len(slice) == 0 {
		return make([]T, 0), errors.New("empty slice")
	}

	if len(slice) < n {
		return make([]T, 0), errors.New("array too small")
	}

	res := make([]T, len(slice))
	copy(res, slice)
	sortSlice(res)

	return res[:n], nil
}

func Sum[T int | float64](i []T) (o T) {
	for _, v := range i {
		o += v
	}
	return
}
