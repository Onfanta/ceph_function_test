package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)


func PoolTest() {
	cmds := [] string{"ceph osd pool create func_test0 1",
		"ceph osd lspools | grep test0",
		"ceph osd dump | grep func_test0",
		"ceph osd pool rename func_test0 func_test1",
		"rados df",
		"ceph osd pool delete func_test1 func_test1 --yes-i-really-really-mean-it"}

	infom := []string{
		"创建存储池","查看存储池选项","查看存储池统计信息","重命名存储池","存储池配额","删除存储池",
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
