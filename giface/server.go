package giface

// 定义服务器接口
type GServer interface {
	// 启动服务器方法
	Start()
	// 停止服务器方法
	Stop()
	// 开启业务服务方法
	Serve()
	// 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(msgId uint32, router GRouter)
	// 得到链接管理
	GetConnMgr() GConnManager
	// 设置该Server的连接创建时Hook函数
	SetOnConnStart(func(GConnection))
	// 设置该Server的连接断开时的Hook函数
	SetOnConnStop(func(GConnection))
	// 调用连接OnConnStart Hook函数
	CallOnConnStart(conn GConnection)
	// 调用连接OnConnStop Hook函数
	CallOnConnStop(conn GConnection)
}