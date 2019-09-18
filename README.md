# logging

#### 说明
logging 基于当前日志格式封装日志库 。

#### 主要功能
支持日志路径

支持日志命名

支持日志级别由低到高 Debug，Info，Warn, Error ,Fatal

支持日志格式

支持日志保留时间

### godoc手册
[logging](http://10.75.29.40:6060/pkg/git.intra.weibo.com/adx/logging/)

### 示例：
```
import (
    "fmt"
    "git.intra.weibo.com/adx/logging"
)
type Config struct {
	Log logging.Logger `toml:"logging" json:"logging"`
}
func main(){
	logTomlStr := `
[logging]
level = "info"
  [logging.file]
  filename = "log.log"
  path = "tmp/logs"
  file_name_date_format = "20060102.150405"
  file_name_date_align = true
  rotation_count = 3
`
}
	c := Config{}
	if _, err := toml.Decode(logTomlStr, &c); err != nil {
		panic(err)
	}
	if err := c.Log.Validate(); err != nil {
		panic(err)
	}
	c.Log.Info("test")
}

```
### 相关文档

