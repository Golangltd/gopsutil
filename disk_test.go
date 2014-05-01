package gopsutil

import (
	"fmt"
	"runtime"
	"testing"
)

func TestDisk_usage(t *testing.T) {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	v, err := DiskUsage(path)
	if err != nil {
		t.Errorf("error %v", err)
	}
	fmt.Println(v)
}

func TestDisk_partitions(t *testing.T) {
	_, err := DiskPartitions(false)
	if err != nil {
		t.Errorf("error %v", err)
	}
}

func TestDisk_io_counters(t *testing.T) {
	ret, err := DiskIOCounters()
	if err != nil {
		t.Errorf("error %v", err)
	}
	for _, io := range ret {
		if io.Name == "" {
			t.Errorf("io_counter error %v", io)
		}
	}
}
