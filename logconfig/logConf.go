package logconfig

import (
    "fmt"
    "github.com/sirupsen/logrus"
    "io"
    "mygin/settings"
    "os"
    "path/filepath"
    "runtime"
    "strings"
)

/*
   日志配置（待完善）
*/

// CustomFormatter 自定义日志格式
type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
    line := entry.Caller.Line

    // 获取当前时间，格式为 yyyy-MM-dd HH:mm:ss,SSS
    timestamp := entry.Time.Format("2006-01-02 15:04:05.000")

    msg := entry.Message

    // 获取调用栈（仅 error 级别及以上）
    stackTrace := ""
    if entry.Level <= logrus.ErrorLevel {
        stackBuf := make([]byte, 1024)
        n := runtime.Stack(stackBuf, false)
        stackTrace = string(stackBuf[:n])
        msg = fmt.Sprintf("%s\n%s", msg, stackTrace)
    }

    // 提取文件路径
    fullPath := entry.Caller.File

    // 生成详细的模块路径，如 api.a.b
    modulePath := getModulePath(fullPath)

    entryData := ""
    if len(entry.Data) > 0 {
        for key, value := range entry.Data {
            entryData += fmt.Sprintf("[%s=%v]", key, value)
        }
    }

    // 构造日志格式
    log := fmt.Sprintf("[%-5s][%s][module:%s][line:%d]%s %s\n",
        entry.Level.String(), // 日志级别
        timestamp,            // 时间戳
        // entry.Data["module"], // 自定义字段：模块
        modulePath, // 自定义字段：模块
        line,       // 行号
        entryData,  // 用户自定义信息
        msg,        // 日志消息
    )

    return []byte(log), nil
}

// getModulePath 解析文件路径为模块路径
func getModulePath(filePath string) string {
    // 获取相对路径
    relativePath, err := filepath.Rel(settings.BaseDir, filePath)
    if err != nil {
        return filePath
    }

    // 将路径转换为模块路径 (替换 / 为 .)
    modulePath := strings.ReplaceAll(relativePath, string(filepath.Separator), ".")

    // 移除文件后缀
    modulePath = strings.TrimSuffix(modulePath, filepath.Ext(modulePath))

    // 例如：api.a.b
    return modulePath
}

// InitLogger 初始化日志配置
func InitLogger() *logrus.Logger {
    logger := logrus.New()

    // 设置日志级别
    logger.SetLevel(logrus.DebugLevel)

    // 设置开启打印文件名及行号
    logger.SetReportCaller(true)

    // 设置自定义格式
    logger.SetFormatter(&CustomFormatter{})

    // 创建 logs 目录
    logDir := "./logs"
    if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
        panic(err)
    }

    // 配置日志文件
    logFilePath := filepath.Join(logDir, "service.log")
    logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }

    // 设置多目标输出（控制台 + 文件）
    logger.SetOutput(io.MultiWriter(os.Stdout, logFile))

    return logger
}

// func main() {
//     // 初始化 logger
//     logger := InitLogger()
//
//     fmt.Println("Project Path:", baseDir)
//
//     // 示例日志
//     logger.WithFields(logrus.Fields{
//         "module": "main",
//     }).Info("Service is starting...")
//     logger.WithFields(logrus.Fields{
//         "module": "auth",
//     }).Warn("Authentication failed")
//     logger.WithFields(logrus.Fields{
//         "module": "database",
//     }).Error("Database connection error")
//
//     log := logger.WithFields(logrus.Fields{
//         "module": "main",
//     })
//     a := 100
//     log.Infof("Info 日志 --- %d", a)
//     log.Errorf("Error 日志 --- %d", a)
//
//     log.Infof("Info 日志2 --- %d", a)
//     log.Infof("Info 日志3 --- %d", a)
//     log.Infof("Info 日志4 --- %d", a)
// }
