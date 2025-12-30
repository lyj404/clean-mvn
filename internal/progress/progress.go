package progress

import (
	"fmt"
	"strings"
	"time"

	"github.com/lyj404/clean-mvn/internal/logger"
)

// 进度条样式配置
const (
	barWidth    = 40 // 进度条内部字符长度
	refreshRate = 100 * time.Millisecond
	filledChar  = "█" // 填充字符
	emptyChar   = "░" // 空白字符
)

// 上次更新时间和速率计算
var (
	lastUpdateTime time.Time
	lastCount      int
)

// DrawProgressBar 绘制一个进度条
func DrawProgressBar(totalCount, currentCount int, activityText string, keepOnScreen, showCount bool) {
	// 计算进度百分比
	progress := 0.0
	if totalCount > 0 {
		progress = float64(currentCount) / float64(totalCount)
	}
	percentage := int(progress * 100)

	// 确保进度和百分比不会超过100%
	if percentage > 100 {
		percentage = 100
	}
	if progress > 1.0 {
		progress = 1.0
	}

	// 计算填充字符和空白字符数量
	filledChars := int(progress * float64(barWidth))
	emptyChars := barWidth - filledChars

	// 构建进度条字符串
	bar := strings.Repeat(filledChar, filledChars) + strings.Repeat(emptyChar, emptyChars)

	// 计算速率(items/sec)
	now := time.Now()
	rate := 0.0
	elapsed := now.Sub(lastUpdateTime)
	if elapsed > refreshRate && lastUpdateTime != (time.Time{}) {
		countDiff := currentCount - lastCount
		rate = float64(countDiff) / elapsed.Seconds()
	}

	// 更新上次计数和时间
	if elapsed > refreshRate || lastUpdateTime.Equal((time.Time{})) {
		lastCount = currentCount
		lastUpdateTime = now
	}

	// 预计剩余时间
	remaining := ""
	if rate > 0 && currentCount < totalCount {
		remainingSecs := float64(totalCount-currentCount) / rate
		if remainingSecs < 60 {
			remaining = fmt.Sprintf("%.1fs", remainingSecs)
		} else if remainingSecs < 3600 {
			remaining = fmt.Sprintf("%.1fm", remainingSecs/60)
		} else {
			remaining = fmt.Sprintf("%.1fh", remainingSecs/3600)
		}
	}

	// 构建最终输出
	output := fmt.Sprintf("\r%s: %s|%s%s%s|%s %3d%%",
		activityText,
		logger.ColorCyan,
		bar,
		logger.ColorCyan,
		logger.ColorReset,
		logger.ColorCyan,
		percentage)

	if showCount && totalCount > 0 {
		output += fmt.Sprintf(" %d/%d", currentCount, totalCount)
	}

	if rate > 0 {
		output += fmt.Sprintf(" [%.1f it/s", rate)
		if remaining != "" {
			output += fmt.Sprintf(", %s left", remaining)
		}
		output += "]"
	}

	output += logger.ColorReset
	fmt.Print(output)

	// 如果达到100%且要求保留，则打印换行
	if keepOnScreen && (percentage == 100 || currentCount >= totalCount) {
		fmt.Print("\n")
		// 重置状态变量
		lastUpdateTime = time.Time{}
		lastCount = 0
	}
}
