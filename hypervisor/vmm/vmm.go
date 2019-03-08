/*
Package vmm provides support for OpenBSD's vmm hypervisor.
*/
package vmm

import (
	"errors"
	"os"

	"github.com/golang/glog"
	"github.com/hyperhq/runv/hypervisor"
	"github.com/hyperhq/runv/hypervisor/types"
)

const (
	vmctl = "/usr/sbin/vmctl"
)

// Driver implements the hypervisor.HypervisorDriver interface.
type Driver struct {
	executable string
}

// Context implements the hypervisor.DriverContext interface.
type Context struct {
	driver *Driver
}

func InitDriver() *Driver {
	_, err := os.Stat(vmctl)
	if err != nil {
		glog.Errorf("cannot initialize driver: %v", err)
		return nil
	}
	return &Driver{
		executable: vmctl,
	}
}

func (vd *Driver) Name() string {
	return "vmm"
}

func (vd *Driver) InitContext(homeDir string) hypervisor.DriverContext {
	glog.Errorf("InitContext(%s): Entering", homeDir)
	return &Context{
		driver: vd,
	}
}

func (vd *Driver) LoadContext(persisted map[string]interface{}) (hypervisor.DriverContext, error) {
	if t, ok := persisted["hypervisor"]; !ok || t != "empty" {
		return nil, errors.New("wrong driver type in persist info")
	}
	return nil, errors.New("XXX Implement me")
}

func (vd *Driver) SupportLazyMode() bool {
	return false
}

func (vd *Driver) SupportVmSocket() bool {
	return false
}

func (vc *Context) Launch(ctx *hypervisor.VmContext) {
	glog.Errorf("Launch(%v): Entering", ctx)
}

func (vc *Context) Associate(ctx *hypervisor.VmContext) {
	glog.Errorf("Associate(%v): Entering", ctx)
}

func (vc *Context) Dump() (map[string]interface{}, error) {
	glog.Errorf("Dump(%v): Entering")
	return nil, errors.New("XXX Implement me")
}

func (vc *Context) AddDisk(ctx *hypervisor.VmContext, sourceType string, blockInfo *hypervisor.DiskDescriptor, result chan<- hypervisor.VmEvent) {
}

func (vc *Context) RemoveDisk(ctx *hypervisor.VmContext, blockInfo *hypervisor.DiskDescriptor, callback hypervisor.VmEvent, result chan<- hypervisor.VmEvent) {
}

func (vc *Context) AddNic(ctx *hypervisor.VmContext, host *hypervisor.HostNicInfo, guest *hypervisor.GuestNicInfo, result chan<- hypervisor.VmEvent) {
}

func (vc *Context) RemoveNic(ctx *hypervisor.VmContext, n *hypervisor.InterfaceCreated, callback hypervisor.VmEvent, result chan<- hypervisor.VmEvent) {
}

func (vc *Context) SetCpus(ctx *hypervisor.VmContext, cpus int) error {
	return errors.New("XXX Implement me")
}

func (vc *Context) AddMem(ctx *hypervisor.VmContext, slot, size int) error {
	return errors.New("XXX Implement me")
}

func (vc *Context) Save(ctx *hypervisor.VmContext, path string) error {
	return errors.New("XXX Implement me")
}

func (vc *Context) Shutdown(ctx *hypervisor.VmContext) {
}

func (vc *Context) Kill(ctx *hypervisor.VmContext) {
}

func (vc *Context) Pause(ctx *hypervisor.VmContext, pause bool) error {
	return errors.New("XXX Implement me")
}

func (vc *Context) Stats(ctx *hypervisor.VmContext) (*types.PodStats, error) {
	return nil, errors.New("XXX Implement me")
}

func (vc *Context) Close() {
}
