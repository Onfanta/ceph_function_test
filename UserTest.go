package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func UserTest() {
	cmds := [] string{
		"ceph auth list |grep client",            //枚举用户
		"ceph auth get-or-create client.george mon 'allow r' osd 'allow rw pool=liverpool' -o george.keyring",		//创建用户
		"ceph auth caps client.george mon 'allow rw' osd 'allow rwx pool=liverpool'",		//修改用户能力
		"ceph auth print-key client.george ",		//查看用户密钥
		"echo 'caps: [osd] allow *' >/etc/ceph/george.keyring ",		//密钥文件添加信息
		"ceph auth export client.george -o /etc/ceph/george.keyring",		//导入用户
		"ceph auth list |grep george",		//检验用户是否存在
		"ceph auth del client.george",}		//删除用户


		infom :=[]string{
			"枚举用户","创建用户","修改用户能力","查看用户密钥","删除用户","密钥文件添加信息","导入用户","检验用户是否存在",
		}
	for i:=0;i<len(cmds);i++  {
		cmd := exec.Command("/bin/bash","-c",cmds[i])
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			defer stdin.Close()
			io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
		}()

		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s:Success!\n%s\n", infom[i],out)
	}
}

