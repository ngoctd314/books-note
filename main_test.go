//go:build !integration

package main

import (
	"testing"
	"time"
)

func TestFn1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test")
	}
	// t.Parallel()
	time.Sleep(time.Second)
	t.Log("run in parallel")
}

func TestFn2(t *testing.T) {
	// t.Parallel()
	time.Sleep(time.Second)
	t.Log("run in parallel")
}
