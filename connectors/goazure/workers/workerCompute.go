package workers

import (
	"connectors/goazure/config"
	"connectors/goazure/network"
	"connectors/goazure/resource"
	"context"
	"fmt"
	"libraries/goclient"
	"libraries/goclient/models"
	"time"
)

type WorkerCompute struct {
	clientGandalf *goclient.ClientGandalf
}

func NewWorkerCompute(identity string, connections []string) *WorkerCompute {
	workerCompute := new(WorkerCompute)
	workerCompute.clientGandalf = goclient.NewClientGandalf(identity, connections)

	return workerCompute
}

func (r WorkerCompute) Run() {
	done := make(chan bool)
	go r.CreateVM()
	//go r.UpdateVM()
	go r.StartVM()
	go r.RestartVM()
	go r.StopVM()
	<-done
}

func (r WorkerCompute) CreateVM() {
	id := r.clientGandalf.CreateIteratorCommand()
	fmt.Println(id)
	command := r.clientGandalf.WaitCommand("CREATE_VM", id)
	fmt.Println(command)

	payload := models.Unmarshall(command.GetPayload())

	groupeName := payload.GetValues()["groupName"]
	virtualNetworkName := payload.GetValues()["virtualNetworkName"]
	subnet1Name := payload.GetValues()["subnet1Name"]
	subnet2Name := payload.GetValues()["subnet2Name"]
	nsgName := payload.GetValues()["nsgName"]
	ipName := payload.GetValues()["ipName"]
	nicName := payload.GetValues()["nicName"]
	vmName := payload.GetValues()["vmName"]
	username := payload.GetValues()["username"]
	password := payload.GetValues()["password"]
	sshPublicKeyPath := payload.GetValues()["sshPublicKeyPath"]

	var groupName = config.GenerateGroupName("VM")
	// TODO: remove and use local `groupName` only
	config.SetGroupName(groupName)

	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()
	defer resource.Cleanup(ctx)

	_, err := resource.CreateGroup(ctx, groupeName)
	if err != nil {
		util.LogAndPanic(err)
	}

	_, err = network.CreateVirtualNetworkAndSubnets(ctx, virtualNetworkName, subnet1Name, subnet2Name)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("created vnet and 2 subnets")

	_, err = network.CreateNetworkSecurityGroup(ctx, nsgName)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("created network security group")

	_, err = network.CreatePublicIP(ctx, ipName)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("created public IP")

	_, err = network.CreateNIC(ctx, virtualNetworkName, subnet1Name, nsgName, ipName, nicName)
	if err != nil {
		util.LogAndPanic(err)
	}
	util.PrintAndLog("created nic")

	_, err = CreateVM(ctx, vmName, nicName, username, password, sshPublicKeyPath)
	if err == nil {
		r.client.SendEvent(command.GetUUID(), "10000", "SUCCES", "test")
	}
	r.client.SendEvent(command.GetUUID(), "10000", "FAIL", "test")
}

/* func (r WorkerCompute) UpdateVM() {
	id := r.clientGandalf.CreateIteratorCommand()
	fmt.Println(id)
	command := r.clientGandalf.WaitCommand("UPDATE_VM", id)
	fmt.Println(command)

	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()
	defer resource.Cleanup(ctx)

	// set or change VM metadata
	_, err = UpdateVM(ctx, vmName, map[string]*string{
		"runtime": to.StringPtr("go"),
		"cloud":   to.StringPtr("azure"),
	})
	if err == nil {
		r.client.SendEvent(command.GetUUID(), "10000", "SUCCES", "test")
	}
	r.client.SendEvent(command.GetUUID(), "10000", "FAIL", "test")
} */

func (r WorkerCompute) StartVM() {
	id := r.clientGandalf.CreateIteratorCommand()
	fmt.Println(id)
	command := r.clientGandalf.StartVM("START_VM", id)
	fmt.Println(command)

	payload = models.Unmarshall(command.GetPayload())
	vmName := payload.GetValues()["vmName"]

	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()
	defer resource.Cleanup(ctx)

	// set or change system state
	_, err = StartVM(ctx, vmName)
	if err == nil {
		r.client.SendEvent(command.GetUUID(), "10000", "SUCCES", "test")
	}
	r.client.SendEvent(command.GetUUID(), "10000", "FAIL", "test")
}

func (r WorkerCompute) RestartVM() {
	id := r.clientGandalf.CreateIteratorCommand()
	fmt.Println(id)
	command := r.clientGandalf.WaitCommand("RESTART_VM", id)
	fmt.Println(command)

	payload = models.Unmarshall(command.GetPayload())
	vmName := payload.GetValues()["vmName"]

	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()
	defer resource.Cleanup(ctx)

	// set or change system state
	_, err = RestartVM(ctx, vmName)
	if err == nil {
		r.client.SendEvent(command.GetUUID(), "10000", "SUCCES", "test")
	}
	r.client.SendEvent(command.GetUUID(), "10000", "FAIL", "test")
}

func (r WorkerCompute) StopVM() {
	id := r.clientGandalf.CreateIteratorCommand()
	fmt.Println(id)
	command := r.clientGandalf.WaitCommand("STOP_VM", id)
	fmt.Println(command)

	payload = models.Unmarshall(command.GetPayload())
	vmName := payload.GetValues()["vmName"]

	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	defer cancel()
	defer resource.Cleanup(ctx)

	// set or change system state
	_, err = StopVM(ctx, vmName)
	if err == nil {
		r.client.SendEvent(command.GetUUID(), "10000", "SUCCES", "test")
	}
	r.client.SendEvent(command.GetUUID(), "10000", "FAIL", "test")
}
