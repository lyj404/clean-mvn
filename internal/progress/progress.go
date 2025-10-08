package progress

import (
	"fmt"
	"strings"

	"clean-mvn/internal/logger"
)

// DrawProgressBar 绘制一个进度条
func DrawProgressBar(totalCount, currentCount int, activityText string, keepOnScreen, showCount bool) {
	const barWidth = 55 // 进度条内部字符长度

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
	bar := strings.Repeat("#", filledChars) + strings.Repeat(" ", emptyChars)

	// 构建最终输出
	output := fmt.Sprintf("\r%s: [%s%s%s]%s %3d%%", activityText, logger.ColorCyan, bar, logger.ColorReset, logger.ColorCyan, percentage)

	if showCount && totalCount > 0 {
		output += fmt.Sprintf(" (%d/%d)", currentCount, totalCount)
	}
	output += logger.ColorReset

	fmt.Print(output)

	// 如果达到100%且要求保留，则打印换行
	if keepOnScreen && (percentage == 100 || currentCount >= totalCount) {
		fmt.Print("\n")
	}
}
