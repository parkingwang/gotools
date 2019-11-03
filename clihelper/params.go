package clihelper

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	daemon = false
	help   = false
	config = "/configs/conf.debug.ini"

	htext = "Show this help info"
	dtext = "Start Server as daemon"
	ctext = "set configuration file"
)

var text = `Usage : %s [-h,--help | -d,--daemon | -c filename]
-h,--help    Show this help info
-d,--daemon  Start Server as daemon
-c filename  set configuration file (default %s)`

// Params 命令执行参数
type Params interface {
	Parse()
	Hook(ext map[string]Extra)
	Daemon() bool
	Config() string
}

// Extra 扩展参数类型
type Extra struct {
	Name  string
	Value string
	Usage string
	Data  *string
}

type params struct {
	help   bool
	daemon bool
	config string
	extra  map[string]Extra
}

// NewParams 命令参数
func NewParams() Params {
	p := &params{}
	return p
}

func (p *params) Parse() {
	// 默认配置文件路径
	cfile := func(name string) string {
		file, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(file)
		index := strings.LastIndex(path, string(os.PathSeparator))
		return path[:index] + name
	}(config)

	for _, ext := range p.extra {
		text += "\n-" + ext.Name + ` ` + ext.Usage + "(default:" + ext.Value + ")"
		flag.StringVar(ext.Data, ext.Name, ext.Value, ext.Usage)
	}
	flag.BoolVar(&p.help, "h", help, "")
	flag.BoolVar(&p.help, "help", help, htext)
	flag.BoolVar(&p.daemon, "d", daemon, "")
	flag.BoolVar(&p.daemon, "daemon", daemon, dtext)
	flag.StringVar(&p.config, "c", cfile, ctext)
	flag.Parse()

	// 输出帮助信息
	if (*p).help {
		fmt.Printf(text, os.Args[0], cfile)
		os.Exit(0)
	}

	// 检查配置文件是否存在
	if _, e := os.Stat(p.config); e != nil {
		if os.IsNotExist(e) {
			fmt.Println("The configuration file:", e.Error())
			os.Exit(0)
		}
	}
}

func (p *params) Hook(ext map[string]Extra) {
	p.extra = ext
}

// Daemon 是否守护进程标识
func (p *params) Daemon() bool {
	return p.daemon
}

// Config 配置文件路径
func (p *params) Config() string {
	return p.config
}
