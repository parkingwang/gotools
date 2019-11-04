# clihelper 命令行帮助工具

## params
命令行参数定义及值获取工具

### 特性
* 默认提供`-d`守护进程参数,`-c`配置文件参数
* 提供的`Hook`方法追加自定义命令行参数定义
* 持续推进中

### 示例

```go
package main

var extra  = map[string]clihelper.Extra{
    "tablename": clihelper.Extra{
        Name:  "t",
        Value: "",
        Usage: "migrate table name",
        Data:  new(string),
    },
    "operatename": clihelper.Extra{
        Name:  "op",
        Value: "",
        Usage: "migrate operate name",
        Data:  new(string),
    },
}

func init(){
    params = clihelper.NewParams()
    params.Hook(extra)
    params.Parse()
    
    ...
}

```
