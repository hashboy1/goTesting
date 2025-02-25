package main
import (
	"fmt"
	"net"
	"log"
	"os"
)


func main() {

//建立socket，监听端口
	netListen, err := net.Listen("tcp", "192.168.0.160:8091")
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}
//处理连接`
func handleConnection(conn net.Conn) {

	buffer := make([]byte, 2048)

	for {

		n, err := conn.Read(buffer)

		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}


		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

	}

}
//Log function
func Log(v ...interface{}) {
	log.Println(v...)
}
//CheckError function
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}