package scanner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lyj404/clean-mvn/internal/logger"
	"github.com/lyj404/clean-mvn/pkg/types"
)

func TestNewScanner(t *testing.T) {
	logger := logger.NewCustomLogger()
	scanner := NewScanner(logger)
	if scanner == nil {
		t.Error("NewScanner() returned nil")
	}
}

func TestScanRepository(t *testing.T) {
	logger := logger.NewCustomLogger()

	tests := []struct {
		name    string
		setup   func() string
		want    int
		wantErr bool
	}{
		{
			name: "no lastUpdated files",
			setup: func() string {
				dir := t.TempDir()
				os.WriteFile(filepath.Join(dir, "normal.txt"), []byte("test"), 0644)
				return dir
			},
			want: 0,
		},
		{
			name: "single lastUpdated file",
			setup: func() string {
				dir := t.TempDir()
				subdir := filepath.Join(dir, "artifact")
				os.Mkdir(subdir, 0755)
				os.WriteFile(filepath.Join(subdir, "file.txt"), []byte("test"), 0644)
				os.WriteFile(filepath.Join(subdir, "file.lastUpdated"), []byte("test"), 0644)
				return dir
			},
			want: 1,
		},
		{
			name: "multiple lastUpdated files",
			setup: func() string {
				dir := t.TempDir()
				for i := 0; i < 3; i++ {
					subdir := filepath.Join(dir, "artifact"+string(rune('0'+i)))
					os.Mkdir(subdir, 0755)
					os.WriteFile(filepath.Join(subdir, "file.txt"), []byte("test"), 0644)
					os.WriteFile(filepath.Join(subdir, "file.lastUpdated"), []byte("test"), 0644)
				}
				return dir
			},
			want: 3,
		},
		{
			name: "nested directories",
			setup: func() string {
				dir := t.TempDir()
				subdir1 := filepath.Join(dir, "level1")
				os.Mkdir(subdir1, 0755)
				subdir2 := filepath.Join(subdir1, "level2")
				os.Mkdir(subdir2, 0755)
				os.WriteFile(filepath.Join(subdir2, "file.txt"), []byte("test"), 0644)
				os.WriteFile(filepath.Join(subdir2, "file.lastUpdated"), []byte("test"), 0644)
				return dir
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewScanner(logger)
			inputPath := tt.setup()
			scanConfig := types.ScanConfig{
				InputPath:               inputPath,
				MaxConcurrentGoRoutines: 2,
			}

			result := s.ScanRepository(scanConfig)

			if (result.Error != nil) != tt.wantErr {
				t.Errorf("ScanRepository() error = %v, wantErr %v", result.Error, tt.wantErr)
			}

			if len(result.Results) != tt.want {
				t.Errorf("ScanRepository() found %d directories, want %d", len(result.Results), tt.want)
			}

			if len(result.Results) > 0 && result.TotalSize == 0 {
				t.Error("ScanRepository() TotalSize should be > 0 when results found")
			}
		})
	}
}

func TestGetDirSize(t *testing.T) {
	logger := logger.NewCustomLogger()
	s := NewScanner(logger)

	tests := []struct {
		name  string
		setup func() string
		want  int64
	}{
		{
			name: "empty directory",
			setup: func() string {
				return t.TempDir()
			},
			want: 0,
		},
		{
			name: "single file",
			setup: func() string {
				dir := t.TempDir()
				content := "test content"
				os.WriteFile(filepath.Join(dir, "test.txt"), []byte(content), 0644)
				return dir
			},
			want: int64(len("test content")),
		},
		{
			name: "multiple files",
			setup: func() string {
				dir := t.TempDir()
				os.WriteFile(filepath.Join(dir, "test1.txt"), []byte("content1"), 0644)
				os.WriteFile(filepath.Join(dir, "test2.txt"), []byte("content2"), 0644)
				return dir
			},
			want: int64(len("content1") + len("content2")),
		},
		{
			name: "nested files",
			setup: func() string {
				dir := t.TempDir()
				subdir := filepath.Join(dir, "subdir")
				os.Mkdir(subdir, 0755)
				os.WriteFile(filepath.Join(subdir, "nested.txt"), []byte("nested content"), 0644)
				return dir
			},
			want: int64(len("nested content")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup()
			got, err := s.getDirSize(path)
			if err != nil {
				t.Errorf("getDirSize() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("getDirSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
