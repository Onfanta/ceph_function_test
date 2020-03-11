package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func RgwTest() {
	cmds := [] string{
		"ceph-deploy rgw create ",		//部署rgw
		"radosgw-admin user create --uid fanta --display-name 'Infanta'",		//创建用户:s3
		"radosgw-admin subuser create --uid=fanta --subuser=fanta:swift --access=full",		//创建子用户：swift
		"radosgw-admin user list",		//查看用户信息
		"radosgw-admin user modify --uid fanta --display-name 'Onfanta'  --max_buckets 2000",  //修改用户（s3）信息
		"radosgw-admin key create --uid fanta --display-name 'Onfanta' --key-type=s3  --access-key=123456 --secret=123456",		//新建ak
		"radosgw-admin caps add --uid=fanta --caps=\"users=*\"",		//添加用户管理权限
		"radosgw-admin quota enable --quota-scope=bucket --uid=fanta",		//启用bucket配额
		"radosgw-admin quota enable --quota-scope=user --uid fanta",		//启用用户配额
		"radosgw-admin quota set --quota-scope=user --uid=fanta --max-objects=1024 --max-size=1024",		//设置用户配额
		"radosgw-admin user info --uid fanta",		//获取配额信息
		"radosgw-admin user stats --uid=fanta --sync-stats",		//更新配额统计信息
		"radosgw-admin usage show --uid fanta",		//获取用户用量统计信息
		"ceph osd erasure-code-profile set rgw_ec_profile k=3 m=2 crush-root=default plugin=isa crush-failure-domain=host",		//创建基于纠删码的对象存储池
		"radosgw-admin quota disable --quota-scope=user --uid fanta",		//禁用用户配额
		"radosgw-admin quota disable --quota-scope=bucket --uid=fanta",			//禁用bucket配额
		"radosgw-admin caps rm --uid=fanta --caps=\"users=*\"",			//删除用户管理权限
		"radosgw-admin key rm --uid fanta --display-name 'Onfanta' --key-type=s3 --access-key=123456",		//删除ak
		"radosgw-admin subuser rm --subuser=fanta:swift ",			//删除子用户
		"radosgw-admin user rm --uid fanta ",			//删除用户
	}

	infom :=[] string{
		"部署rgw","创建用户:s3",
		"创建子用户：swift",
		"查看用户信息",
		"修改用户（s3）信息",
		"新建ak","添加用户管理权限",
		"启用bucket配额",
		"启用用户配额",
		"设置用户配额",
		"获取配额信息",
		"更新配额统计信息",
		"获取用户用量统计信息",
		"创建基于纠删码的对象存储池",
		"禁用用户配额",
		"禁用bucket配额","删除用户管理权限","删除子用户","删除用户",
		//"删除ak",
	}

	var hostname string

	for i:=0;i<len(cmds);i++ {
		if i == 0  {
			fmt.Printf("plz input ur <hostname> which the host u want 2 deploy rgw:")
			fmt.Scanln(&hostname)
			cmd := exec.Command("/bin/bash", "-c", cmds[0]+hostname)
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
		} else {
			cmd := exec.Command("/bin/bash", "-c", cmds[i])
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
			fmt.Printf("%sSuccess!\n%s\n", infom[i],out)
		}
	}
}
