package utils

import "net"

func GetFreePort() (int, error) {
	// 解析 "localhost:0" 地址，表示请求系统分配一个空闲的端口
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	// 监听指定的地址和端口，实际分配一个空闲的端口
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close() // 在函数结束时关闭监听器，释放端口

	// 返回实际分配的端口号
	return l.Addr().(*net.TCPAddr).Port, nil
}
