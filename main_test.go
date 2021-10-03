package main

import (
	"testing"
)

var sortSliceTests = []struct {
	in         []string
	out        []string
	smallestId int
}{
	{
		[]string{
			"bob",
			"alice",
			"john",
			"kate",
			"sophia",
			"emma",
			"james",
			"robert",
		},
		[]string{
			"alice",
			"bob",
			"emma",
			"james",
			"john",
			"kate",
			"robert",
			"sophia",
		},
		1,
	},
	{
		[]string{
			"alice",
			"bob",
			"emma",
			"james",
			"john",
			"kate",
			"robert",
			"sophia",
		},
		[]string{
			"alice",
			"bob",
			"emma",
			"james",
			"john",
			"kate",
			"robert",
			"sophia",
		},
		0,
	},
}

func TestSortSlice(t *testing.T) {
	for i, tt := range sortSliceTests {
		t.Run("sort"+string(rune(i)), func(t *testing.T) {
			sortedSlice := selectionSort(tt.in)
			if !equalStrings(sortedSlice, tt.out) {
				t.Errorf("Test sortSlice failed")
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	slice, err := readFile("./test_data.txt")
	if err != nil {
		t.Errorf("Test readFile failed: %s", err)
		return
	}
	if !equalStrings(slice, sortSliceTests[0].in) {
		t.Errorf("Test readFile failed")
		return
	}
}

func TestFindSmallest(t *testing.T) {
	for i, tt := range sortSliceTests {
		t.Run("sort"+string(rune(i)), func(t *testing.T) {
			idx := findSmallest(tt.in)
			if !equalInt(idx, tt.smallestId) {
				t.Errorf("Test sortSlice failed")
			}
		})
	}
}

func TestRun(t *testing.T) {
	run("./test_data.txt")
}

func equalInt(a, b int) bool {
	if a != b {
		return false
	}

	return true
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
