package main

import (
	"runtime"

	"github.com/lyj404/clean-mvn/internal/cleaner"
	"github.com/lyj404/clean-mvn/internal/cli"
	"github.com/lyj404/clean-mvn/internal/logger"
	"github.com/lyj404/clean-mvn/internal/scanner"
	"github.com/lyj404/clean-mvn/internal/util"
	"github.com/lyj404/clean-mvn/pkg/types"
)

func main() {
	// 解析命令行配置
	config := cli.ParseConfig()

	// 显示帮助信息
	if cli.IsHelpRequested() {
		cli.ShowUsage()
		return
	}

	// 初始化日志器
	loggerInstance := logger.NewCustomLogger()
	if config.LogFile != "" {
		if err := loggerInstance.AddLogFile(config.LogFile); err != nil {
			loggerInstance.Warning("Failed to create log file: %v", err)
		}
	}

	// 获取输入路径
	inputPath := config.Path
	if inputPath == "" {
		inputPath = util.GetUserInput()
		if inputPath == "" {
			inputPath = cli.GetDefaultPath()
		}
	}

	// 验证路径
	if inputPath == "" {
		loggerInstance.Error("No path specified. Use --path to specify the Maven repository path.")
		cli.ShowUsage()
		return
	}

	if !util.ValidatePath(loggerInstance, inputPath) {
		return
	}

	// 开始扫描
	loggerInstance.Info("Starting scan of Maven repository path: %s", inputPath)

	// 确定并发工作数
	workers := config.Workers
	if workers == 0 {
		workers = cli.GetWorkersFromEnv()
	}
	if workers == 0 {
		workers = runtime.NumCPU()
	}

	// 创建扫描器并执行扫描
	s := scanner.NewScanner(loggerInstance)
	scanConfig := types.ScanConfig{
		InputPath:               inputPath,
		MaxConcurrentGoRoutines: workers,
	}

	scanResult := s.ScanRepository(scanConfig)

	// 处理扫描结果
	if scanResult.Error != nil {
		loggerInstance.Error("An error occurred during file system scan: %v", scanResult.Error)
	}

	// 显示扫描结果
	util.DisplayScanResults(loggerInstance, scanResult)

	// 如果没有找到文件，退出
	if len(scanResult.Results) == 0 {
		loggerInstance.Success("Congratulations! No '.lastUpdated' related build directories found in your Maven repository.")
		return
	}

	// 预览模式
	if config.DryRun {
		loggerInstance.Info("Dry run mode: Would delete %d directories, freeing %.2f MB space.",
			len(scanResult.Results), float64(scanResult.TotalSize)/1024/1024)
		return
	}

	// 询问用户确认
	if !config.Force && !util.GetUserConfirmation() {
		loggerInstance.Info("Operation cancelled.")
		return
	}

	// 执行清理
	cleanerInstance := cleaner.NewCleaner(loggerInstance)
	cleanResult := cleanerInstance.CleanDirectories(scanResult.Results)

	// 显示清理结果
	loggerInstance.Success("Cleanup complete! Deleted %d directories, freed %.2f MB space.",
		cleanResult.DeletedCount, float64(cleanResult.DeletedSize)/1024/1024)
}
