package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func EcTest() {
	cmds := [] string{
		"ceph osd erasure-code-profile set Ecprofile crush-failure-domain=osd k=3 m=2",
		"ceph osd erasure-code-profile ls",
		"echo test > test",
		"ceph osd pool create Ecpool 1 erasure Ecprofile",
		"ceph osd erasure-code-profile get Ecprofile",
		"ceph osd dump | grep Ecpool",
		"rados put -p Ecpool object1 test",
		"rados -p Ecpool ls",
		"ceph osd pool delete Ecpool Ecpool --yes-i-really-really-mean-it"}

	infom := []string{
		"创建ECProfile","查看ecprofile","写入测试文件","创建纠删码存储池","查看纠删码配置","检验纠删码池","上传文件到纠删码","检验文件是否存在","删除Ecpool",
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
