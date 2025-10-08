package scanner

import (
	"io/fs"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"clean-mvn/internal/logger"
	"clean-mvn/internal/progress"
	"clean-mvn/pkg/types"
)

// Scanner Maven仓库扫描器
type Scanner struct {
	logger *logger.CustomLogger
}

// NewScanner 创建新的扫描器
func NewScanner(logger *logger.CustomLogger) *Scanner {
	return &Scanner{
		logger: logger,
	}
}

// ScanRepository 扫描Maven仓库，查找包含.lastUpdated文件的目录
func (s *Scanner) ScanRepository(config types.ScanConfig) types.ScanResult {
	startTime := time.Now()

	var (
		mu                      sync.Mutex
		results                 []types.Result
		totalSizeAtomic         atomic.Int64
		wg                      sync.WaitGroup
		maxConcurrentGoRoutines = config.MaxConcurrentGoRoutines
		sem                     = make(chan struct{}, maxConcurrentGoRoutines)
	)

	// 扫描进度条
	scanProgressCount := atomic.Int64{}
	scanStop := make(chan bool)
	scanDone := make(chan bool)

	// 启动进度条 Goroutine
	go s.runProgressBar(&scanProgressCount, scanStop, scanDone)

	// 使用 filepath.WalkDir 遍历文件系统
	err := filepath.WalkDir(config.InputPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			s.logger.Warning("Error accessing path %s: %v (skipped)", path, err)
			return nil
		}

		scanProgressCount.Add(1) // 每次处理一个文件/目录，递增计数

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(d.Name(), ".lastUpdated") {
			parentDir := filepath.Dir(path)

			sem <- struct{}{}
			wg.Add(1)
			go func(dirPath string) {
				defer wg.Done()
				defer func() { <-sem }()

				size, err := s.getDirSize(dirPath)
				if err != nil {
					s.logger.Warning("Error calculating directory %s size: %v (skipped)", dirPath, err)
					return
				}

				mu.Lock()
				results = append(results, types.Result{Path: dirPath, Size: size})
				mu.Unlock()
				totalSizeAtomic.Add(size)
			}(parentDir)

			return filepath.SkipDir
		}

		return nil
	})

	wg.Wait() // 等待所有协程完成目录大小计算

	// 停止进度条
	scanStop <- true
	<-scanDone

	endTime := time.Now()
	duration := endTime.Sub(startTime).Milliseconds()

	// 去重results
	uniqueResultsMap := make(map[string]types.Result)
	for _, res := range results {
		uniqueResultsMap[res.Path] = res
	}

	uniqueResults := make([]types.Result, 0, len(uniqueResultsMap))
	var actualTotalSize int64
	for _, res := range uniqueResultsMap {
		uniqueResults = append(uniqueResults, res)
		actualTotalSize += res.Size
	}

	return types.ScanResult{
		Results:   uniqueResults,
		TotalSize: actualTotalSize,
		Duration:  duration,
		Error:     err,
	}
}

// runProgressBar 运行扫描进度条
func (s *Scanner) runProgressBar(scanProgressCount *atomic.Int64, scanStop <-chan bool, scanDone chan<- bool) {
	defer close(scanDone)

	lastPrintTime := time.Now()
	for {
		select {
		case <-scanStop:
			// 扫描完成时，强制显示100%并换行
			currentScanned := int(scanProgressCount.Load())
			progress.DrawProgressBar(currentScanned, currentScanned, "Scanning", true, false)
			return
		default:
			if time.Since(lastPrintTime) >= 100*time.Millisecond {
				currentScanned := int(scanProgressCount.Load())
				// 使用一个动态的"虚拟总数"，让进度条看起来在前进
				virtualTotal := currentScanned + (currentScanned / 3) + 1000

				progress.DrawProgressBar(virtualTotal, currentScanned, "Scanning", false, false)
				lastPrintTime = time.Now()
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// getDirSize 递归计算目录的总大小
func (s *Scanner) getDirSize(path string) (int64, error) {
	var size int64
	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			size += info.Size()
		}
		return nil
	})
	return size, err
}
