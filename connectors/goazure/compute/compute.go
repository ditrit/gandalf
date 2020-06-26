package compute

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/Azure/go-autorest/autorest"
)

func getVMClient() compute.VirtualMachinesClient {
	vmClient := compute.NewVirtualMachinesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	vmClient.Authorizer = a
	vmClient.AddToUserAgent(config.UserAgent())
	return vmClient
}

func getVMExtensionsClient() compute.VirtualMachineExtensionsClient {
	extClient := compute.NewVirtualMachineExtensionsClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	extClient.Authorizer = a
	extClient.AddToUserAgent(config.UserAgent())
	return extClient
}

func CreateVM(ctx context.Context, vmName, nicName, username, password, sshPublicKeyPath string) (vm compute.VirtualMachine, err error) {
	nic, _ := network.GetNic(ctx, nicName)

	var sshKeyData string
	if _, err = os.Stat(sshPublicKeyPath); err == nil {
		sshBytes, err := ioutil.ReadFile(sshPublicKeyPath)
		if err != nil {
			log.Fatalf("failed to read SSH key data: %v", err)
		}
		sshKeyData = string(sshBytes)
	} else {
		sshKeyData = fakepubkey
	}

	vmClient := getVMClient()
	future, err := vmClient.CreateOrUpdate(
		ctx,
		config.GroupName(),
		vmName,
		compute.VirtualMachine{
			Location: to.StringPtr(config.Location()),
			VirtualMachineProperties: &compute.VirtualMachineProperties{
				HardwareProfile: &compute.HardwareProfile{
					VMSize: compute.VirtualMachineSizeTypesBasicA0,
				},
				StorageProfile: &compute.StorageProfile{
					ImageReference: &compute.ImageReference{
						Publisher: to.StringPtr(publisher),
						Offer:     to.StringPtr(offer),
						Sku:       to.StringPtr(sku),
						Version:   to.StringPtr("latest"),
					},
				},
				OsProfile: &compute.OSProfile{
					ComputerName:  to.StringPtr(vmName),
					AdminUsername: to.StringPtr(username),
					AdminPassword: to.StringPtr(password),
					LinuxConfiguration: &compute.LinuxConfiguration{
						SSH: &compute.SSHConfiguration{
							PublicKeys: &[]compute.SSHPublicKey{
								{
									Path: to.StringPtr(
										fmt.Sprintf("/home/%s/.ssh/authorized_keys",
											username)),
									KeyData: to.StringPtr(sshKeyData),
								},
							},
						},
					},
				},
				NetworkProfile: &compute.NetworkProfile{
					NetworkInterfaces: &[]compute.NetworkInterfaceReference{
						{
							ID: nic.ID,
							NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		return vm, fmt.Errorf("cannot create vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return vm, fmt.Errorf("cannot get the vm create or update future response: %v", err)
	}

	return future.Result(vmClient)
}

func GetVM(ctx context.Context, vmName string) (compute.VirtualMachine, error) {
	vmClient := getVMClient()
	return vmClient.Get(ctx, config.GroupName(), vmName, compute.InstanceView)
}

func UpdateVM(ctx context.Context, vmName string, tags map[string]*string) (vm compute.VirtualMachine, err error) {

	vm, err = GetVM(ctx, vmName)
	if err != nil {
		return
	}

	vm.Tags = tags

	vmClient := getVMClient()
	future, err := vmClient.CreateOrUpdate(ctx, config.GroupName(), vmName, vm)
	if err != nil {
		return vm, fmt.Errorf("cannot update vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return vm, fmt.Errorf("cannot get the vm create or update future response: %v", err)
	}

	return future.Result(vmClient)
}

func DeallocateVM(ctx context.Context, vmName string) (osr autorest.Response, err error) {
	vmClient := getVMClient()
	future, err := vmClient.Deallocate(ctx, config.GroupName(), vmName)
	if err != nil {
		return osr, fmt.Errorf("cannot deallocate vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return osr, fmt.Errorf("cannot get the vm deallocate future response: %v", err)
	}

	return future.Result(vmClient)
}

func StartVM(ctx context.Context, vmName string) (osr autorest.Response, err error) {
	vmClient := getVMClient()
	future, err := vmClient.Start(ctx, config.GroupName(), vmName)
	if err != nil {
		return osr, fmt.Errorf("cannot start vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return osr, fmt.Errorf("cannot get the vm start future response: %v", err)
	}

	return future.Result(vmClient)
}

func RestartVM(ctx context.Context, vmName string) (osr autorest.Response, err error) {
	vmClient := getVMClient()
	future, err := vmClient.Restart(ctx, config.GroupName(), vmName)
	if err != nil {
		return osr, fmt.Errorf("cannot restart vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return osr, fmt.Errorf("cannot get the vm restart future response: %v", err)
	}

	return future.Result(vmClient)
}

func StopVM(ctx context.Context, vmName string) (osr autorest.Response, err error) {
	vmClient := getVMClient()
	future, err := vmClient.PowerOff(ctx, config.GroupName(), vmName, nil)
	if err != nil {
		return osr, fmt.Errorf("cannot power off vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return osr, fmt.Errorf("cannot get the vm power off future response: %v", err)
	}

	return future.Result(vmClient)
}
