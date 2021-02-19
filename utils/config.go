package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/shhnwangjian/tcpframe/giface"
)

/*
   定义一个全局的对象
*/
var GlobalConfig *GConfig

/*
   存储一切有关Zinx框架的全局参数，供其他模块使用
   一些参数也可以通过 用户根据 json文件来配置
*/
type GConfig struct {
	TcpServer        giface.GServer // 全局Server对象
	Host             string         // 服务器主机IP
	TcpPort          int            // 服务器主机监听端口号
	Name             string         // 服务器名称
	Version          string         // 版本号
	MaxPacketSize    uint32         // 数据包的最大值
	MaxConn          int            // 服务器主机允许的最大链接个数
	WorkerPoolSize   uint32         // 业务工作Worker池的数量
	MaxWorkerTaskLen uint32         // 业务工作Worker对应负责的任务队列最大任务存储数量
	ConfFilePath     string
	MaxMsgChanLen    uint32
}

/*
   提供init方法，默认加载
*/
func init() {
	//初始化GlobalObject变量，设置一些默认值
	GlobalConfig = &GConfig{
		Name:             "tcpframe",
		Version:          "V0.1",
		TcpPort:          7777,
		Host:             "0.0.0.0",
		MaxConn:          12000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	//从配置文件中加载一些用户配置的参数
	GlobalConfig.Reload()
}

//读取用户的配置文件
func (g *GConfig) Reload() {
	//data, err := ioutil.ReadFile(path.Join(SelfDir(), "config.json.sample"))
	data, err := ioutil.ReadFile("/Users/wangjian/go/src/github.com/shhnwangjian/tcpframe/utils/config.json.sample")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalConfig)
	if err != nil {
		panic(err)
	}
}

func SelfPath() string {
	p, _ := filepath.Abs(os.Args[0])
	return p
}

func SelfDir() string {
	return filepath.Dir(SelfPath())
}
