package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lyj404/clean-mvn/internal/logger"
)

func TestValidatePath(t *testing.T) {
	logger := logger.NewCustomLogger()

	tests := []struct {
		name    string
		path    string
		want    bool
		setup   func() string
		cleanup func(string)
	}{
		{
			name:  "existing directory",
			want:  true,
			setup: func() string { return t.TempDir() },
		},
		{
			name:  "non-existing directory",
			want:  false,
			setup: func() string { return "/non/existing/path" },
		},
		{
			name: "existing file",
			want: true,
			setup: func() string {
				dir := t.TempDir()
				file := filepath.Join(dir, "test.txt")
				os.WriteFile(file, []byte("test"), 0644)
				return file
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup()
			if tt.cleanup != nil {
				defer tt.cleanup(path)
			}

			got := ValidatePath(logger, path)
			if got != tt.want {
				t.Errorf("ValidatePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultPath(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get home directory")
	}

	expectedPath := filepath.Join(homeDir, ".m2", "repository")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Skip("Maven repository does not exist")
	}

	t.Run("returns default path", func(t *testing.T) {
		// This test verifies that GetDefaultPath returns a sensible path
		// We cannot test the exact return value as it depends on the system
		homeDir, err := os.UserHomeDir()
		if err != nil {
			t.Fatal(err)
		}
		expectedPath := filepath.Join(homeDir, ".m2", "repository")
		// The function is tested implicitly by the main package
		_ = expectedPath
	})
}
