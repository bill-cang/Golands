package main

import "time"

type Conn interface {
	// Read从连接中读取数据
	// Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
	Read(b []byte) (n int, err error)
	// Write从连接中写入数据
	// Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
	Write(b []byte) (n int, err error)
	// Close方法关闭该连接
	// 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
	Close() error
	// 返回本地网络地址
	LocalAddr() Addr
	// 返回远端网络地址
	RemoteAddr() Addr
	// 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
	// deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
	// deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
	// 参数t为零值表示不设置期限
	SetDeadline(t time.Time) error
	// 设定该连接的读操作deadline，参数t为零值表示不设置期限
	SetReadDeadline(t time.Time) error
	// 设定该连接的写操作deadline，参数t为零值表示不设置期限
	// 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
	SetWriteDeadline(t time.Time) error
}

type Addr interface {
	Network() string // 网络名
	String() string  // 字符串格式的地址
}
