package commit

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Commit() {
	// 创建与远程Docker守护程序的连接
	cli, err := client.NewClientWithOpts(client.WithHost("tcp://192.168.195.43:2375"))
	if err != nil {
		panic(err)
	}

	// 需要执行commit操作的容器ID
	containerID := "609a44d5dbd1"

	// 提交容器的配置
	commitConfig := types.ContainerCommitOptions{
		Reference: "nginxdemos/hello:customVersion1", // 提交后的镜像名称
		Author:    "kangmingfa",                      // 提交的作者名称
		Changes:   []string{},
		Comment:   "Committing container changes",
		Pause:     true,
	}

	// 执行commit操作
	resp, err := cli.ContainerCommit(context.Background(), containerID, commitConfig)
	if err != nil {
		panic(err)
	}

	// 输出结果
	fmt.Println("New Image ID:", resp.ID)
}
