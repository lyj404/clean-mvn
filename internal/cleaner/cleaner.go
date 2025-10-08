package cleaner

import (
	"os"

	"clean-mvn/internal/logger"
	"clean-mvn/internal/progress"
	"clean-mvn/pkg/types"
)

// Cleaner 清理器
type Cleaner struct {
	logger *logger.CustomLogger
}

// NewCleaner 创建新的清理器
func NewCleaner(logger *logger.CustomLogger) *Cleaner {
	return &Cleaner{
		logger: logger,
	}
}

// CleanResult 清理结果
type CleanResult struct {
	DeletedCount int
	DeletedSize  int64
}

// CleanDirectories 删除指定的目录列表
func (c *Cleaner) CleanDirectories(results []types.Result) CleanResult {
	totalToDelete := len(results)
	var deletedCount int
	var deletedSize int64

	c.logger.Info("Starting file deletion...")

	for i := 0; i < totalToDelete; i++ {
		res := results[i]
		if err := os.RemoveAll(res.Path); err != nil {
			c.logger.Error("Failed to delete directory '%s': %v", res.Path, err)
		} else {
			deletedCount++
			deletedSize += res.Size
		}

		// 更新进度条
		progress.DrawProgressBar(totalToDelete, deletedCount, "Deleting", i == totalToDelete-1, true)
	}

	return CleanResult{
		DeletedCount: deletedCount,
		DeletedSize:  deletedSize,
	}
}
