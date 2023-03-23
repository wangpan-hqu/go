package ssh

import (
	"golang.org/x/crypto/ssh"
	"os"

	//	"golang.org/x/crypto/ssh/terminal"
	"log"
	//	"os"
	"time"
)

func Ssh() {
	sshConfig := &ssh.ClientConfig{
		User: "wangpan",
		Auth: []ssh.AuthMethod{
			ssh.Password("201314"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		ClientVersion:   "",
		Timeout:         10 * time.Second,
	}
	//建立与SSH服务器的连接
	sshClient, err := ssh.Dial("tcp", "192.168.199.200:22", sshConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer sshClient.Close()
	log.Println("sessionId: ", sshClient.SessionID())
	log.Println("user: ", sshClient.User())
	log.Println("ssh server version: ", string(sshClient.ServerVersion()))
	log.Println("ssh client version: ", string(sshClient.ClientVersion()))

	//打开交互式会话(A session is a remote execution of a program.)
	//https://tools.ietf.org/html/rfc4254#page-10
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalln("Failed to create ssh session", err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     //打开回显
		ssh.TTY_OP_ISPEED: 14400, //输入速率 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //输出速率 14.4kbaud
		ssh.VSTATUS:       1,
	}
	/*
		//使用VT100终端来实现tab键提示，上下键查看历史命令，clear键清屏等操作
		//VT100 start
		//windows下不支持VT100
		fd := int(os.Stdin.Fd())
		oldState, err := terminal.MakeRaw(fd)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer terminal.Restore(fd, oldState)
		//VT100 end

		termWidth, termHeight, err := terminal.GetSize(fd)
	*/
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	//打开伪终端
	//https://tools.ietf.org/html/rfc4254#page-11
	//	err = session.RequestPty("xterm", termHeight, termWidth, modes)
	err = session.RequestPty("linux", 32, 160, modes)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//session.Run("ls")

	session.Run("pwd")
	//	session.Output("")
	//	session.CombinedOutput("")
	/*	//启动一个远程shell
		//https://tools.ietf.org/html/rfc4254#page-13
		err = session.Shell()
		if err != nil {
			log.Fatalln(err.Error())
		}

		//等待远程命令结束或远程shell退出
		err = session.Wait()
		if err != nil {
			log.Fatalln(err.Error())
		}

	*/
}
