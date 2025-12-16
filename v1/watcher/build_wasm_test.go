package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

// Sanity test: ensure the watcher-dta-v1 main package compiles for wasip1/wasm.
func TestWatcherV1_Build_WASM(t *testing.T) {
	_, thisFile, _, _ := runtime.Caller(0)
	dir := filepath.Dir(thisFile)

	out := filepath.Join(dir, "test-watcher-dta-v1.wasm")
	defer os.Remove(out)

	cmd := exec.Command("go", "build", "-o", out, ".")
	cmd.Dir = dir
	cmd.Env = append(os.Environ(),
		"GOOS=wasip1",
		"GOARCH=wasm",
		"CGO_ENABLED=0",
	)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build wasm for watcher-dta-v1: %v\n%s", err, string(output))
	}
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("expected wasm output missing: %v", err)
	}
}

