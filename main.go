package main

import (
	"runtime"

	"github.com/lyj404/clean-mvn/internal/cleaner"
	"github.com/lyj404/clean-mvn/internal/logger"
	"github.com/lyj404/clean-mvn/internal/scanner"
	"github.com/lyj404/clean-mvn/internal/util"
	"github.com/lyj404/clean-mvn/pkg/types"
)

func main() {
	// 初始化日志器
	cl := logger.NewCustomLogger()

	// 获取用户输入
	inputPath := util.GetUserInput()

	// 验证路径
	if !util.ValidatePath(cl, inputPath) {
		return
	}

	// 开始扫描
	cl.Info("Starting scan of Maven repository path: %s", inputPath)

	// 创建扫描器并执行扫描
	s := scanner.NewScanner(cl)
	scanConfig := types.ScanConfig{
		InputPath:               inputPath,
		MaxConcurrentGoRoutines: runtime.NumCPU(),
	}

	scanResult := s.ScanRepository(scanConfig)

	// 处理扫描结果
	if scanResult.Error != nil {
		cl.Error("An error occurred during file system scan: %v", scanResult.Error)
	}

	// 显示扫描结果
	util.DisplayScanResults(cl, scanResult)

	// 如果没有找到文件，退出
	if len(scanResult.Results) == 0 {
		cl.Success("Congratulations! No '.lastUpdated' related build directories found in your Maven repository.")
		return
	}

	// 询问用户确认
	if !util.GetUserConfirmation() {
		cl.Info("Operation cancelled.")
		return
	}

	// 执行清理
	cleanerInstance := cleaner.NewCleaner(cl)
	cleanResult := cleanerInstance.CleanDirectories(scanResult.Results)

	// 显示清理结果
	cl.Success("Cleanup complete! Deleted %d directories, freed %.2f MB space.",
		cleanResult.DeletedCount, float64(cleanResult.DeletedSize)/1024/1024)
}
