package main

import (
	"testing"
)

func TestSortDirsBySize(t *testing.T) {
	dirs := map[string]uint64{
		"test":  1000,
		"test2": 2000,
	}
	sortedDirs := sortDirsBySize(dirs)
	if sortedDirs[0].Size != 2000 {
		t.Fatalf("Sort dirs failed!")
	}
}

func TestPercent(t *testing.T) {
	part, all := uint64(10), uint64(100)
	pcent := Percent(part, all)

	if pcent != float64(10) {
		t.Fatalf("Wrong percent!")
	}
}
