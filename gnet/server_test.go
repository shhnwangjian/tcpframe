package gnet

import (
	"fmt"
	"testing"

	"github.com/shhnwangjian/tcpframe/giface"
)

//ping test 自定义路由
type PingRouter struct {
	BaseRouter
}

//Ping Handle
func (this *PingRouter) Handle(request giface.GRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

//HelloZinxRouter Handle
type HelloZinxRouter struct {
	BaseRouter
}

func (this *HelloZinxRouter) Handle(request giface.GRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

//创建连接的时候执行
func DoConnectionBegin(conn giface.GConnection) {
	fmt.Println("DoConnecionBegin is Called ... ")
	conn.SetProperty("Name", "ops")
	conn.SetProperty("Home", "https://www.jianshu.com/")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
}

//连接断开的时候执行
func DoConnectionLost(conn giface.GConnection) {
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Conn Property Home = ", home)
	}
	fmt.Println("DoConneciotnLost is Called ... ")
}

// Server 模块的测试函数
func TestServer(t *testing.T) {
	/*
	   服务端测试
	*/
	//1 创建一个server 句柄 s
	s := NewServer()

	//注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//配置路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//2 开启服务
	s.Serve()
}
