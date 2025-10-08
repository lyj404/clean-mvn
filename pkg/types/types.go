package types

// Result 用于存储找到的需要删除的目录信息
type Result struct {
	Path string
	Size int64
}

// ScanConfig 扫描配置
type ScanConfig struct {
	InputPath               string
	MaxConcurrentGoRoutines int
}

// ScanResult 扫描结果
type ScanResult struct {
	Results   []Result
	TotalSize int64
	Duration  int64 // 毫秒
	Error     error
}
