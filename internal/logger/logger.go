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
	IconSuccess  = "✅"
	IconError    = "❌"
	IconInfo     = "💡"
	IconWarning  = "⚠️"
	IconTime     = "⏳"
	IconInput    = "👉"
	IconScanning = "🔍"
)

// CustomLogger 自定义日志器，用于封装带颜色和图标的输出
type CustomLogger struct {
	consoleLogger *log.Logger
	fileLogger    *log.Logger
}

// NewCustomLogger 创建新的自定义日志器
func NewCustomLogger() *CustomLogger {
	return &CustomLogger{
		consoleLogger: log.New(os.Stdout, "", 0),
	}
}

// AddLogFile 添加日志文件
func (cl *CustomLogger) AddLogFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	cl.fileLogger = log.New(file, "", log.LstdFlags)
	return nil
}

func (cl *CustomLogger) logToAll(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	cl.consoleLogger.Printf("%s", message)
	if cl.fileLogger != nil {
		cl.fileLogger.Printf("%s", removeAnsiCodes(message))
	}
}

func (cl *CustomLogger) Info(format string, a ...interface{}) {
	cl.logToAll("%s %s %s\n", IconInfo, fmt.Sprintf(format, a...), ColorReset)
}

func (cl *CustomLogger) Success(format string, a ...interface{}) {
	cl.logToAll("%s %s%s %s\n", IconSuccess, ColorGreen, fmt.Sprintf(format, a...), ColorReset)
}

func (cl *CustomLogger) Error(format string, a ...interface{}) {
	cl.logToAll("%s %s%s %s\n", IconError, ColorRed, fmt.Sprintf(format, a...), ColorReset)
}

func (cl *CustomLogger) Warning(format string, a ...interface{}) {
	cl.logToAll("%s %s%s %s\n", IconWarning, ColorYellow, fmt.Sprintf(format, a...), ColorReset)
}

func (cl *CustomLogger) Time(format string, a ...interface{}) {
	cl.logToAll("%s %s%s %s\n", IconTime, ColorGreen, fmt.Sprintf(format, a...), ColorReset)
}

// PrintRaw 打印原始内容
func PrintRaw(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// GetIconInput 获取输入提示图标
func GetIconInput() string {
	return IconInput
}

// removeAnsiCodes 移除 ANSI 颜色代码
func removeAnsiCodes(s string) string {
	result := ""
	inEscape := false
	for _, c := range s {
		if c == '\033' {
			inEscape = true
		} else if inEscape && c == 'm' {
			inEscape = false
		} else if !inEscape {
			result += string(c)
		}
	}
	return result
}
