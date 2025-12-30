package logger

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewCustomLogger(t *testing.T) {
	logger := NewCustomLogger()
	if logger == nil {
		t.Error("NewCustomLogger() returned nil")
	}
}

func TestAddLogFile(t *testing.T) {
	logFile := filepath.Join(t.TempDir(), "test.log")
	logger := NewCustomLogger()

	err := logger.AddLogFile(logFile)
	if err != nil {
		t.Errorf("AddLogFile() error = %v", err)
	}

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Error("Log file was not created")
	}

	// Test adding another log file (should replace)
	logFile2 := filepath.Join(t.TempDir(), "test2.log")
	err = logger.AddLogFile(logFile2)
	if err != nil {
		t.Errorf("AddLogFile() second call error = %v", err)
	}
}

func TestRemoveAnsiCodes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no codes",
			input: "plain text",
			want:  "plain text",
		},
		{
			name:  "red code",
			input: "\033[31mred text\033[0m",
			want:  "red text",
		},
		{
			name:  "multiple codes",
			input: "\033[31m\033[1mbold red\033[0m",
			want:  "bold red",
		},
		{
			name:  "mixed content",
			input: "start \033[32mgreen\033[0m middle \033[33myellow\033[0m end",
			want:  "start green middle yellow end",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeAnsiCodes(tt.input)
			if got != tt.want {
				t.Errorf("removeAnsiCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogMethods(t *testing.T) {
	logger := NewCustomLogger()

	tests := []struct {
		name   string
		method func(string, ...interface{})
	}{
		{"Info", logger.Info},
		{"Success", logger.Success},
		{"Error", logger.Error},
		{"Warning", logger.Warning},
		{"Time", logger.Time},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just verify that methods don't panic
			tt.method("test message %s", "argument")
		})
	}
}

func TestGetIconInput(t *testing.T) {
	icon := GetIconInput()
	if icon == "" {
		t.Error("GetIconInput() returned empty string")
	}
	if !strings.Contains(icon, "👉") {
		t.Error("GetIconInput() does not contain expected icon")
	}
}

func TestPrintRaw(t *testing.T) {
	// Just verify it doesn't panic
	PrintRaw("test message %s\n", "argument")
}
