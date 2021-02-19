package gnet

import (
	"github.com/shhnwangjian/tcpframe/giface"
)

type Request struct {
	conn giface.GConnection // 已经和客户端建立好的 链接
	msg  giface.GMessage    // 客户端请求的数据
}

// 获取请求连接信息
func (r *Request) GetConnection() giface.GConnection {
	return r.conn
}

// 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// 获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
