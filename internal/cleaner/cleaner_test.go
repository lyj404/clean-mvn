package cleaner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lyj404/clean-mvn/internal/logger"
	"github.com/lyj404/clean-mvn/pkg/types"
)

func TestCleanDirectories(t *testing.T) {
	logger := logger.NewCustomLogger()

	tests := []struct {
		name    string
		setup   func() []types.Result
		want    int
		wantErr bool
	}{
		{
			name: "empty list",
			setup: func() []types.Result {
				return []types.Result{}
			},
			want: 0,
		},
		{
			name: "single directory",
			setup: func() []types.Result {
				dir := t.TempDir()
				subdir := filepath.Join(dir, "test-subdir")
				os.Mkdir(subdir, 0755)
				os.WriteFile(filepath.Join(subdir, "test.txt"), []byte("test content"), 0644)
				return []types.Result{{Path: subdir, Size: int64(len("test content"))}}
			},
			want: 1,
		},
		{
			name: "multiple directories",
			setup: func() []types.Result {
				dir := t.TempDir()
				results := []types.Result{}
				for i := 0; i < 3; i++ {
					subdir := filepath.Join(dir, "test-subdir"+string(rune('0'+i)))
					os.Mkdir(subdir, 0755)
					os.WriteFile(filepath.Join(subdir, "test.txt"), []byte("test content"), 0644)
					results = append(results, types.Result{Path: subdir, Size: int64(len("test content"))})
				}
				return results
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCleaner(logger)
			results := tt.setup()
			result := c.CleanDirectories(results)

			if result.DeletedCount != tt.want {
				t.Errorf("CleanDirectories() DeletedCount = %v, want %v", result.DeletedCount, tt.want)
			}

			// Verify directories are deleted
			for _, r := range results {
				if _, err := os.Stat(r.Path); !os.IsNotExist(err) {
					t.Errorf("Directory %s was not deleted", r.Path)
				}
			}
		})
	}
}

func TestCleanResult(t *testing.T) {
	tests := []struct {
		name string
		CleanResult
	}{
		{"zero", CleanResult{DeletedCount: 0, DeletedSize: 0}},
		{"non-zero", CleanResult{DeletedCount: 5, DeletedSize: 1024}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just test that the struct can be created
			_ = tt.CleanResult
		})
	}
}
