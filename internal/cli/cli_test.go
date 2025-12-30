package cli

import (
	"flag"
	"os"
	"testing"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantPath    string
		wantForce   bool
		wantDryRun  bool
		wantWorkers int
	}{
		{
			name:        "default values",
			args:        []string{},
			wantForce:   false,
			wantDryRun:  false,
			wantWorkers: 0,
		},
		{
			name:     "with path",
			args:     []string{"--path", "/test/path"},
			wantPath: "/test/path",
		},
		{
			name:     "with shorthand path",
			args:     []string{"-p", "/test/path"},
			wantPath: "/test/path",
		},
		{
			name:      "with force",
			args:      []string{"--force"},
			wantForce: true,
		},
		{
			name:      "with shorthand force",
			args:      []string{"-f"},
			wantForce: true,
		},
		{
			name:       "with dry run",
			args:       []string{"--dry-run"},
			wantDryRun: true,
		},
		{
			name:       "with shorthand dry run",
			args:       []string{"-d"},
			wantDryRun: true,
		},
		{
			name:        "with workers",
			args:        []string{"--workers", "4"},
			wantWorkers: 4,
		},
		{
			name:        "with shorthand workers",
			args:        []string{"-w", "8"},
			wantWorkers: 8,
		},
		{
			name:        "all options",
			args:        []string{"-p", "/test/path", "-f", "-d", "-w", "4"},
			wantPath:    "/test/path",
			wantForce:   true,
			wantDryRun:  true,
			wantWorkers: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			// Set test args
			os.Args = append([]string{"clean-mvn"}, tt.args...)

			// Reset flag set
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

			config := ParseConfig()

			if config.Path != tt.wantPath {
				t.Errorf("ParseConfig().Path = %v, want %v", config.Path, tt.wantPath)
			}
			if config.Force != tt.wantForce {
				t.Errorf("ParseConfig().Force = %v, want %v", config.Force, tt.wantForce)
			}
			if config.DryRun != tt.wantDryRun {
				t.Errorf("ParseConfig().DryRun = %v, want %v", config.DryRun, tt.wantDryRun)
			}
			if config.Workers != tt.wantWorkers {
				t.Errorf("ParseConfig().Workers = %v, want %v", config.Workers, tt.wantWorkers)
			}
		})
	}
}

func TestIsHelpRequested(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{"no help", []string{"clean-mvn"}, false},
		{"short help", []string{"clean-mvn", "-h"}, true},
		{"long help", []string{"clean-mvn", "--help"}, true},
		{"plain help", []string{"clean-mvn", "help"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = tt.args

			got := IsHelpRequested()
			if got != tt.want {
				t.Errorf("IsHelpRequested() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultPath(t *testing.T) {
	// Just test that the function runs without panicking
	path := GetDefaultPath()
	if path == "" {
		t.Error("GetDefaultPath() returned empty string")
	}
}

func TestGetWorkersFromEnv(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want int
	}{
		{"no env", "", 0},
		{"valid env", "4", 4},
		{"invalid env", "invalid", 0},
		{"zero env", "0", 0},
		{"negative env", "-1", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldEnv := os.Getenv("CLEAN_MVN_WORKERS")
			defer os.Setenv("CLEAN_MVN_WORKERS", oldEnv)

			if tt.env != "" {
				os.Setenv("CLEAN_MVN_WORKERS", tt.env)
			} else {
				os.Unsetenv("CLEAN_MVN_WORKERS")
			}

			got := GetWorkersFromEnv()
			if got != tt.want {
				t.Errorf("GetWorkersFromEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowUsage(t *testing.T) {
	// Just test that the function runs without panicking
	ShowUsage()
}
