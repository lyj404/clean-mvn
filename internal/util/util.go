package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lyj404/clean-mvn/internal/logger"
	"github.com/lyj404/clean-mvn/pkg/types"
)

// GetUserInput 获取用户输入的路径
func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	logger.PrintRaw("%s", fmt.Sprintf("%s Please enter your Maven repository root path (e.g.: C:\\Users\\YourUser\\.m2\\repository): ", logger.GetIconInput()))
	inputPath, _ := reader.ReadString('\n')
	return strings.TrimSpace(inputPath)
}

// ValidatePath 验证路径是否存在
func ValidatePath(cl *logger.CustomLogger, path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		cl.Error("Path '%s' does not exist.", path)
		return false
	}
	return true
}

// DisplayScanResults 显示扫描结果
func DisplayScanResults(cl *logger.CustomLogger, result types.ScanResult) {
	cl.Time("Scan completed, took %s.", time.Duration(result.Duration*int64(time.Millisecond)).Round(time.Millisecond))

	if len(result.Results) > 0 {
		cl.Info("Found %d unique directories containing '.lastUpdated' files, total %.2f MB to be deleted.",
			len(result.Results), float64(result.TotalSize)/1024/1024)
	}
}

// GetUserConfirmation 获取用户确认
func GetUserConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)
	logger.PrintRaw("%s", fmt.Sprintf("%s Do you confirm deletion of these files? (y/n): ", logger.GetIconInput()))
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))
	return confirm == "y"
}
