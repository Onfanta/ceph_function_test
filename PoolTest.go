package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)


func main() {
	cmds := [] string{"ceph osd pool create func_test0 1",
		"ceph osd lspools | grep test0",
		"ceph osd dump | grep func_test0",
		"ceph osd pool rename func_test0 func_test1",
		"rados df",
		"ceph osd pool delete func_test1 func_test1 --yes-i-really-really-mean-it"}
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
		fmt.Printf("[Success]:\n%s\n", out)
	}
}
