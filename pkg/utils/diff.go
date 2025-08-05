package utils

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type DiffResult[T comparable] struct {
	Added   []T
	Removed []T
}

func DiffCalculate[T constraints.Ordered](old []T, new []T) DiffResult[T] {
	var result DiffResult[T]

	slices.Sort(old)
	slices.Sort(new)

	i := 0
	j := 0
	for i < len(old) && j < len(new) {
		if old[i] < new[j] {
			result.Removed = append(result.Removed, old[i])
			i++
		} else if old[i] > new[j] {
			result.Added = append(result.Added, new[j])
			j++
		} else {
			i++
			j++
		}
	}

	for i < len(old) {
		result.Removed = append(result.Removed, old[i])
		i++
	}

	for j < len(new) {
		result.Added = append(result.Added, new[j])
		j++
	}

	return result
}

func SubsetOf[T comparable](superSet []T, subSet []T) bool {
	if len(subSet) > len(superSet) {
		return false
	}

	var superHM map[T]int = make(map[T]int)
	for _, v := range superSet {
		superHM[v]++
	}
	for _, v := range subSet {
		if superHM[v] == 0 {
			return false
		}
		superHM[v] -= 1
	}
	return true
}
