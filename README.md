# clean-mvn
Clean-MVN 是一个简单高效的工具，用于清理 Maven 仓库中损坏或不完整的下载文件。它会扫描您的`Maven`仓库目录，查找 .lastUpdated文件（表示下载失败的文件），并删除相关目录以帮助解决构建问题。

![效果图](/img/img1.png)

**功能特点:**
* 快速并发扫描 Maven 仓库
* 扫描和删除过程中提供可视化进度指示
* 提供详细的待删除文件大小信息
* 删除前需确认，防止意外操作
* 跨平台支持（Windows、macOS、Linux）
  
## 安装方式
**方法一：通过Go工具安装**
```shell
go install github.com/lyj404/clean-mvn@latest
```
安装完成后，确保`$GOPATH/bin`在系统PATH环境变量中，即可使用`iggen`命令

**方法二：下载预编译二进制文件**

从 [GitHub Releases](https://github.com/lyj404/clean-mvn/releases/latest) 下载适合您系统的最新版本

**使用方法:**
1. 运行应用程序
2. 输入 Maven 仓库路径（通常在 Unix 系统上为 ~/.m2/repository，Windows 上为C:\Users\YourUser\.m2\repository）
3. 查看扫描结果
4. 确认是否删除