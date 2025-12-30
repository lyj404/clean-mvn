package cli

import (
	"flag"
	"os"
	"strconv"
)

// Config CLI 配置
type Config struct {
	Path    string // Maven 仓库路径
	Force   bool   // 是否跳过确认
	DryRun  bool   // 是否只预览不删除
	Workers int    // 并发工作数
	LogFile string // 日志文件路径
}

// ParseConfig 解析命令行参数
func ParseConfig() Config {
	config := Config{}

	flag.StringVar(&config.Path, "path", "", "Maven 仓库路径")
	flag.StringVar(&config.Path, "p", "", "Maven 仓库路径（简写）")
	flag.BoolVar(&config.Force, "force", false, "跳过确认提示")
	flag.BoolVar(&config.Force, "f", false, "跳过确认提示（简写）")
	flag.BoolVar(&config.DryRun, "dry-run", false, "预览模式，只显示将要删除的内容而不实际删除")
	flag.BoolVar(&config.DryRun, "d", false, "预览模式（简写）")
	flag.IntVar(&config.Workers, "workers", 0, "并发工作数（默认：CPU 核心数）")
	flag.IntVar(&config.Workers, "w", 0, "并发工作数（简写）")
	flag.StringVar(&config.LogFile, "log", "", "日志文件路径")
	flag.StringVar(&config.LogFile, "l", "", "日志文件路径（简写）")

	flag.Parse()

	return config
}

// IsHelpRequested 检查是否请求帮助信息
func IsHelpRequested() bool {
	for _, arg := range os.Args {
		if arg == "-h" || arg == "--help" || arg == "help" {
			return true
		}
	}
	return false
}

// ShowUsage 显示使用帮助
func ShowUsage() {
	println("Usage: clean-mvn [options]")
	println()
	println("Options:")
	println("  -p, --path <path>      Maven 仓库路径")
	println("  -f, --force            跳过确认提示")
	println("  -d, --dry-run          预览模式，只显示将要删除的内容而不实际删除")
	println("  -w, --workers <n>     并发工作数（默认：CPU 核心数）")
	println("  -l, --log <file>       日志文件路径")
	println("  -h, --help             显示此帮助信息")
	println()
	println("Environment Variables:")
	println("  MAVEN_REPO_PATH        默认 Maven 仓库路径")
	println()
	println("Examples:")
	println("  clean-mvn --path ~/.m2/repository")
	println("  clean-mvn -p ~/.m2/repository --force")
	println("  clean-mvn -p ~/.m2/repository --dry-run")
	println("  clean-mvn -p ~/.m2/repository --workers 4")
}

// GetDefaultPath 获取默认的 Maven 仓库路径
func GetDefaultPath() string {
	if path := os.Getenv("MAVEN_REPO_PATH"); path != "" {
		return path
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return homeDir + "/.m2/repository"
}

// GetWorkersFromEnv 从环境变量获取并发工作数
func GetWorkersFromEnv() int {
	if workersStr := os.Getenv("CLEAN_MVN_WORKERS"); workersStr != "" {
		if workers, err := strconv.Atoi(workersStr); err == nil && workers > 0 {
			return workers
		}
	}
	return 0
}
