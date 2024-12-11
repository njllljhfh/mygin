package settings

import "os"

// BaseDir 项目根目录
var BaseDir string

func init() {
    // 获取项目根目录
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    BaseDir = wd
}
