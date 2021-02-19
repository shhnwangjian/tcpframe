package gnet

import (
	"fmt"
	"strconv"

	"github.com/shhnwangjian/tcpframe/giface"
	"github.com/shhnwangjian/tcpframe/utils"
)

type MsgHandle struct {
	Apis           map[uint32]giface.GRouter // 存放每个MsgId 所对应的处理方法的map属性
	WorkerPoolSize uint32                    // 业务工作Worker池的数量
	TaskQueue      []chan giface.GRequest    // Worker负责取任务的消息队列
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]giface.GRouter),
		WorkerPoolSize: utils.GlobalConfig.WorkerPoolSize,
		// 一个worker对应一个queue
		TaskQueue: make([]chan giface.GRequest, utils.GlobalConfig.WorkerPoolSize),
	}
}

// 以非阻塞方式处理消息
func (h *MsgHandle) DoMsgHandler(request giface.GRequest) {
	handler, ok := h.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgID(), " is not FOUND!")
		return
	}

	// 执行对应处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// 为消息添加具体的处理逻辑
func (h *MsgHandle) AddRouter(msgId uint32, router giface.GRouter) {
	// 判断当前msg绑定的API处理方法是否已经存在
	if _, ok := h.Apis[msgId]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}

	// 添加msg与api的绑定关系
	h.Apis[msgId] = router
	fmt.Println("Add api msgId = ", msgId)
}

// 启动一个Worker工作流程
func (h *MsgHandle) StartOneWorker(workerID int, taskQueue chan giface.GRequest) {
	fmt.Println("Worker ID = ", workerID, " is started.")
	// 不断的等待队列中的消息
	for {
		select {
		// 有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			h.DoMsgHandler(request)
		}
	}
}

// 启动worker工作池
func (h *MsgHandle) StartWorkerPool() {
	// 遍历需要启动worker的数量，依此启动
	for i := 0; i < int(h.WorkerPoolSize); i++ {
		// 一个worker被启动
		// 给当前worker对应的任务队列开辟空间
		h.TaskQueue[i] = make(chan giface.GRequest, utils.GlobalConfig.MaxWorkerTaskLen)
		// 启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go h.StartOneWorker(i, h.TaskQueue[i])
	}
}

// 将消息交给TaskQueue,由worker进行处理
func (h *MsgHandle) SendMsgToTaskQueue(request giface.GRequest) {
	// 根据ConnID来分配当前的连接应该由哪个worker负责处理
	// 轮询的平均分配法则

	// 得到需要处理此条连接的workerID
	workerID := request.GetConnection().GetConnID() % h.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(), " request msgID=", request.GetMsgID(), "to workerID=", workerID)
	// 将请求消息发送给任务队列
	h.TaskQueue[workerID] <- request
}
