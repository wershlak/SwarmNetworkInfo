package main

import (
	"context"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/olekukonko/tablewriter"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	tasks, err := cli.TaskList(context.Background(), types.TaskListOptions{})
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NODE", "SERVICE", "STATUS", "IMAGE", "NETWORK", "ADDRESS"})

	for _, task := range tasks {
		if task.DesiredState == "running" {
			node, _, err := cli.NodeInspectWithRaw(context.Background(), task.NodeID)
			if err != nil {
				panic(err)
			}
			service, _, err := cli.ServiceInspectWithRaw(context.Background(), task.ServiceID, types.ServiceInspectOptions{})
			if err != nil {
				panic(err)
			}
			image := strings.Split(task.Spec.ContainerSpec.Image, "@")[0]

			for index, network := range task.NetworksAttachments {
				networkinfo, _, err := cli.NetworkInspectWithRaw(context.Background(), network.Network.ID, false)
				if err != nil {
					panic(err)
				}
				if index == 0 {
					data := []string{node.Description.Hostname, service.Spec.Name, string(task.Status.State), image, networkinfo.Name, network.Addresses[0]}
					table.Append(data)
				} else {
					data := []string{"", "", "", "", networkinfo.Name, network.Addresses[0]}
					table.Append(data)
				}
			}
		}
	}

	table.Render()
}
