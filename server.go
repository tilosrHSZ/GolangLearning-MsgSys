package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播channel
	Message chan string
}

// 创建server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + " : " + msg

	this.Message <- sendMsg
}

// 监听Message广播消息channel，一旦有消息全体广播
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将msg发送给全体在线user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// handle 业务
func (this *Server) Handler(conn net.Conn) {
	//当前链接的业务
	fmt.Println("链接建立成功")

	user := NerUser(conn, this)

	user.Online()

	//接受客户端传递发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			//TCP设计规定客户端关闭时的结果：
			//正常下线的结果：n == 0, err == io.EOF
			//某些系统实现：n == 0, err == nil
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}

			//提取用户的消息，要除去\n
			msg := string(buf[:n-1])

			//用户针对msg进行消息处理
			user.DoMessage(msg)
		}
	}()

	//当前handler阻塞
	select {}
}

// 启动server的接口
func (this *Server) StartServer() {
	//socket listten
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}

	//close listen socket
	defer listener.Close()

	//启动监听Message的进程
	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error:", err)
			continue
		}
		//do handler
		go this.Handler(conn)
	}

}
