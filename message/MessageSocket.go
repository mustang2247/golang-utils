// socket相关工具
// 参考连接：http://studygolang.com/articles/8916
// 参考连接：http://studygolang.com/articles/805
// 这是一个高并发socket数据读写
package message

import (
	"net"
	"fmt"
	"utils/logutil"
	"time"
)

var client_num int = 0

/* 开始连接server */
func StartServer(add string, nettype string) bool {
	//tcpAddr, err := net.ResolveTCPAddr(nettype, add)
	//logutil.CheckErr(err)

	//listener, err := net.ListenTCP(nettype, tcpAddr)
	listener, err := net.Listen(nettype, add)
	logutil.CheckErr(err)

	for {
		conn, err := listener.Accept()
		logutil.CheckErr(err)
		if err != nil {
			continue
		}

		client_num++
		fmt.Printf("A new Connection %d.\n", client_num)
		go handlerConnection(conn)
	}

	fmt.Println("Server is Starting")
	return true
}

// 监控连接
func handlerConnection(conn net.Conn) {
	// conn.SetReadDeadline()设置超时 2分钟
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	//defer closeConnection(conn)

	readChannel := make(chan []byte, 1024)
	writeChannel := make(chan []byte, 1024)

	// 并行读写数据
	go readConnection(conn, readChannel)
	go WriteConnection(conn, writeChannel)

	for {
		select {
		case data := <-readChannel:
			if string(data) == "bye" {
				return
			}
			writeChannel <- append([]byte("Back"), data...)
		}
	}

}

// 写入数据
func WriteConnection(conn net.Conn, channel chan []byte) {
	for {
		select {
		case data := <-channel:
			println("Write:", conn.RemoteAddr().String(), string(data))
			_, err := conn.Write(data)
			logutil.CheckErr(err)
		}
	}

}

// 读取数据
func readConnection(conn net.Conn, channel chan []byte) {

	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		logutil.CheckErr(err)
		if err != nil {
			//log.CheckErr(err)
			channel <- []byte("bye") //这里须要进一步改进！
			break
		}
		println("Recei:", conn.RemoteAddr().String(), string(buffer[:n]))
		channel <- buffer[:n]
	}
}

func closeConnection(conn net.Conn) {
	conn.Close()
	client_num--
	fmt.Printf("Now, %d connections is alve.\n", client_num)
}
