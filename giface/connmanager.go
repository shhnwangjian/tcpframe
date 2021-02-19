package giface

/*
	连接管理抽象层
*/
type GConnManager interface {
	Add(conn GConnection)                   // 添加链接
	Remove(conn GConnection)                // 删除连接
	Get(connID uint32) (GConnection, error) // 通过ConnID获取链接
	Len() int                               // 获取当前连接长度
	ClearConn()                             // 删除并停止所有链接
}
