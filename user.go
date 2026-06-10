package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

//crest user api
func NerUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	//启动监听器
	go user.ListenMessage()

	return user
}

//用户上线功能
func (this *User) Online() {

	//用户上线，将用户加入onlinemap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播用户上线
	this.server.BroadCast(this, "has joined")
}

//用户下线功能
func (this *User) Offline() {

	//用户下线，将用户从onlinemap删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播用户上线
	this.server.BroadCast(this, "has left")
}

//用户处理消息的功能
func (this *User) DoMessage(msg string) {
	this.server.BroadCast(this, msg)
}

//监听器
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
