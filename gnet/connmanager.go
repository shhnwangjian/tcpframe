package gnet

import (
	"errors"
	"fmt"
	"sync"

	"github.com/shhnwangjian/tcpframe/giface"
)

/*
ConnManager	连接管理模块
*/
type ConnManager struct {
	connections map[uint32]giface.GConnection // 管理的连接信息
	connLock    sync.RWMutex                  // 读写连接的读写锁
}

/*
NewConnManager	创建一个链接管理
*/
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]giface.GConnection),
	}
}

// Len 获取当前连接长度
func (m *ConnManager) Len() int {
	return len(m.connections)
}

// Add 添加链接
func (m *ConnManager) Add(conn giface.GConnection) {
	// 保护共享资源Map 加写锁
	m.connLock.Lock()
	defer m.connLock.Unlock()

	// 将conn连接添加到ConnMananger中
	m.connections[conn.GetConnID()] = conn

	fmt.Println("connection add to ConnManager successfully: conn num = ", m.Len())
}

// Remove 删除连接
func (m *ConnManager) Remove(conn giface.GConnection) {
	// 保护共享资源Map 加写锁
	m.connLock.Lock()
	defer m.connLock.Unlock()

	//删除连接信息
	delete(m.connections, conn.GetConnID())

	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", m.Len())
}

// Get 通过ConnID获取链接
func (m *ConnManager) Get(connID uint32) (giface.GConnection, error) {
	// 保护共享资源Map 加读锁
	m.connLock.RLock()
	defer m.connLock.RUnlock()

	if conn, ok := m.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// ClearConn 清除并停止所有连接
func (m *ConnManager) ClearConn() {
	// 保护共享资源Map 加写锁
	m.connLock.Lock()
	defer m.connLock.Unlock()

	// 停止并删除全部的连接信息
	for connID, conn := range m.connections {
		// 停止
		conn.Stop()
		// 删除
		delete(m.connections, connID)
	}

	fmt.Println("Clear All Connections successfully: conn num = ", m.Len())
}
