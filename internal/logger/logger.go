package logger

import (
	"fmt"
	"log"
	"os"
)

// ANSI 颜色码
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
)

// Unicode 图标
const (
	IconSuccess  = "✅"  // 绿色勾号
	IconError    = "❌"  // 红色叉号
	IconInfo     = "💡"  // 提示信息
	IconWarning  = "⚠️" // 警告
	IconTime     = "⏳"  // 时间
	IconInput    = "👉"  // 输入提示
	IconScanning = "🔍"  // 扫描中
)

// CustomLogger 自定义日志器，用于封装带颜色和图标的输出
type CustomLogger struct {
	logger *log.Logger
}

// NewCustomLogger 创建新的自定义日志器
func NewCustomLogger() *CustomLogger {
	return &CustomLogger{
		logger: log.New(os.Stdout, "", 0), // 不带前缀和时间戳，更灵活
	}
}

func (cl *CustomLogger) Info(format string, a ...interface{}) {
	cl.logger.Printf("%s", fmt.Sprintf("%s %s %s\n", IconInfo, fmt.Sprintf(format, a...), ColorReset))
}

func (cl *CustomLogger) Success(format string, a ...interface{}) {
	cl.logger.Printf("%s", fmt.Sprintf("%s %s%s %s\n", IconSuccess, ColorGreen, fmt.Sprintf(format, a...), ColorReset))
}

func (cl *CustomLogger) Error(format string, a ...interface{}) {
	cl.logger.Printf("%s", fmt.Sprintf("%s %s%s %s\n", IconError, ColorRed, fmt.Sprintf(format, a...), ColorReset))
}

func (cl *CustomLogger) Warning(format string, a ...interface{}) {
	cl.logger.Printf("%s", fmt.Sprintf("%s %s%s %s\n", IconWarning, ColorYellow, fmt.Sprintf(format, a...), ColorReset))
}

func (cl *CustomLogger) Time(format string, a ...interface{}) {
	cl.logger.Printf("%s", fmt.Sprintf("%s %s%s %s\n", IconTime, ColorGreen, fmt.Sprintf(format, a...), ColorReset))
}

// PrintRaw 打印原始内容
func PrintRaw(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// GetIconInput 获取输入提示图标
func GetIconInput() string {
	return IconInput
}
