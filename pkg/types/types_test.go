package types

import (
	"testing"
)

func TestResult(t *testing.T) {
	tests := []struct {
		name   string
		Result Result
	}{
		{"empty", Result{Path: "", Size: 0}},
		{"with path", Result{Path: "/test/path", Size: 0}},
		{"with size", Result{Path: "", Size: 1024}},
		{"full", Result{Path: "/test/path", Size: 1024}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just test that struct can be created
			_ = tt.Result
			if tt.Result.Path != "" && tt.Result.Size < 0 {
				t.Error("Size should not be negative")
			}
		})
	}
}

func TestScanConfig(t *testing.T) {
	tests := []struct {
		name       string
		ScanConfig ScanConfig
	}{
		{"default", ScanConfig{}},
		{"with path", ScanConfig{InputPath: "/test/path"}},
		{"with workers", ScanConfig{MaxConcurrentGoRoutines: 4}},
		{"full", ScanConfig{InputPath: "/test/path", MaxConcurrentGoRoutines: 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just test that struct can be created
			_ = tt.ScanConfig
			if tt.ScanConfig.MaxConcurrentGoRoutines < 0 {
				t.Error("MaxConcurrentGoRoutines should not be negative")
			}
		})
	}
}

func TestScanResult(t *testing.T) {
	tests := []struct {
		name       string
		ScanResult ScanResult
	}{
		{"empty", ScanResult{}},
		{"with results", ScanResult{Results: []Result{{Path: "/test", Size: 1024}}}},
		{"with total size", ScanResult{TotalSize: 2048}},
		{"with duration", ScanResult{Duration: 1000}},
		{"with error", ScanResult{Error: nil}},
		{"full", ScanResult{
			Results:   []Result{{Path: "/test", Size: 1024}},
			TotalSize: 1024,
			Duration:  1000,
			Error:     nil,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just test that struct can be created
			_ = tt.ScanResult
			if tt.ScanResult.TotalSize < 0 {
				t.Error("TotalSize should not be negative")
			}
			if tt.ScanResult.Duration < 0 {
				t.Error("Duration should not be negative")
			}
		})
	}
}
