# clean-mvn
Clean-MVN 是一个简单高效的工具，用于清理 Maven 仓库中损坏或不完整的下载文件。它会扫描您的 Maven 仓库目录，查找 .lastUpdated 文件（表示下载失败的文件），并删除相关目录以帮助解决构建问题。

![效果图](/img/img1.png)

**功能特点:**
* 快速并发扫描 Maven 仓库
* 扫描和删除过程中提供可视化进度指示
* 提供详细的待删除文件大小信息
* 删除前需确认，防止意外操作
* 跨平台支持（Windows、macOS、Linux）
* 支持命令行参数，便于自动化
* 支持日志文件记录

## 安装方式

**方法一：通过 Go 工具安装**
```shell
go install github.com/lyj404/clean-mvn@latest
```
安装完成后，确保 `$GOPATH/bin` 在系统 PATH 环境变量中，即可使用 `clean-mvn` 命令。

**方法二：下载预编译二进制文件**

从 [GitHub Releases](https://github.com/lyj404/clean-mvn/releases/latest) 下载适合您系统的最新版本

## 使用方法

### 交互模式
直接运行应用程序（不带参数）：
```shell
clean-mvn
```
然后：
1. 输入您的 Maven 仓库路径（通常在 Unix 系统上为 `~/.m2/repository`，Windows 上为 `C:\Users\YourUser\.m2\repository`）
2. 查看扫描结果
3. 确认是否删除

### 命令行模式
使用命令行参数进行自动化操作：

```shell
# 指定路径
clean-mvn --path ~/.m2/repository

# 跳过确认
clean-mvn --path ~/.m2/repository --force

# 预览模式（只显示将要删除的内容而不实际删除）
clean-mvn --path ~/.m2/repository --dry-run

# 指定并发工作数
clean-mvn --path ~/.m2/repository --workers 4

# 保存日志到文件
clean-mvn --path ~/.m2/repository --log cleanup.log

# 组合多个选项
clean-mvn -p ~/.m2/repository -f -w 4 -l cleanup.log
```

### 命令行选项

| 简写 | 完整 | 说明 |
|------|------|------|
| `-p` | `--path` | Maven 仓库路径 |
| `-f` | `--force` | 跳过确认提示 |
| `-d` | `--dry-run` | 预览模式，只显示将要删除的内容而不实际删除 |
| `-w` | `--workers` | 并发工作数（默认：CPU 核心数） |
| `-l` | `--log` | 日志文件路径 |
| `-h` | `--help` | 显示帮助信息 |

### 环境变量

* `MAVEN_REPO_PATH` - 默认 Maven 仓库路径
* `CLEAN_MVN_WORKERS` - 默认并发工作数

### 使用示例

**扫描并清理（需要确认）：**
```shell
clean-mvn --path ~/.m2/repository
```

**自动清理（跳过确认）：**
```shell
clean-mvn --path ~/.m2/repository --force
```

**预览将要删除的内容：**
```shell
clean-mvn --path ~/.m2/repository --dry-run
```

**使用 4 个并发工作数进行清理：**
```shell
clean-mvn --path ~/.m2/repository --workers 4
```

**清理并保存日志：**
```shell
clean-mvn --path ~/.m2/repository --log cleanup.log
```
After installation, ensure `$GOPATH/bin` is in your system PATH environment variable to use the `clean-mvn` command.

**Method 2: Download pre-compiled binaries**

Download the latest version for your system from [GitHub Releases](https://github.com/lyj404/clean-mvn/releases/latest)

## Usage

### Interactive Mode
Run the application without arguments:
```shell
clean-mvn
```
Then:
1. Enter your Maven repository path (typically `~/.m2/repository` on Unix or `C:\Users\YourUser\.m2\repository` on Windows)
2. View scan results
3. Confirm whether to delete

### Command-Line Mode
Use command-line arguments for automation:

```shell
# Specify path
clean-mvn --path ~/.m2/repository

# Skip confirmation
clean-mvn --path ~/.m2/repository --force

# Dry run (show what would be deleted without actually deleting)
clean-mvn --path ~/.m2/repository --dry-run

# Specify number of concurrent workers
clean-mvn --path ~/.m2/repository --workers 4

# Save logs to file
clean-mvn --path ~/.m2/repository --log cleanup.log

# Combine options
clean-mvn -p ~/.m2/repository -f -w 4 -l cleanup.log
```

### Options

| Short | Long | Description |
|-------|------|-------------|
| `-p` | `--path` | Path to Maven repository |
| `-f` | `--force` | Skip confirmation prompt |
| `-d` | `--dry-run` | Show what would be deleted without actually deleting |
| `-w` | `--workers` | Number of concurrent workers (default: number of CPUs) |
| `-l` | `--log` | Log file path |
| `-h` | `--help` | Show help message |

### Environment Variables

* `MAVEN_REPO_PATH` - Default Maven repository path
* `CLEAN_MVN_WORKERS` - Default number of concurrent workers

### Examples

**Scan and clean with confirmation:**
```shell
clean-mvn --path ~/.m2/repository
```

**Automatic cleanup without confirmation:**
```shell
clean-mvn --path ~/.m2/repository --force
```

**Preview what would be deleted:**
```shell
clean-mvn --path ~/.m2/repository --dry-run
```

**Clean with 4 concurrent workers:**
```shell
clean-mvn --path ~/.m2/repository --workers 4
```

**Clean and save logs:**
```shell
clean-mvn --path ~/.m2/repository --log cleanup.log
```