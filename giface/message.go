package giface

/*
   将请求的一个消息封装到message中，定义抽象层接口
*/
type GMessage interface {
	GetDataLen() uint32 // 获取消息数据段长度
	GetMsgId() uint32   // 获取消息ID
	GetData() []byte    // 获取消息内容
	SetMsgId(uint32)    // 设计消息ID
	SetData([]byte)     // 设计消息内容
	SetDataLen(uint32)  // 设置消息数据段长度
}

/*
	消息管理抽象层
*/
type GMsgHandle interface {
	DoMsgHandler(request GRequest)          // 以非阻塞方式处理消息
	AddRouter(msgId uint32, router GRouter) // 为消息添加具体的处理逻辑
	StartWorkerPool()                       // 启动worker工作池
	SendMsgToTaskQueue(request GRequest)    // 将消息交给TaskQueue,由worker进行处理
}
