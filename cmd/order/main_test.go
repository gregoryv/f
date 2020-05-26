package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFilterLines(t *testing.T) {
	lines := []string{"a", "b", "a", "c", "d"}
	sort.Sort(ByPattern{
		lines:    lines,
		patterns: []string{"a", "c"},
	})
	exp := []string{"a", "a", "c", "b", "d"}
	if !reflect.DeepEqual(lines, exp) {
		t.Error(lines)
	}
}
