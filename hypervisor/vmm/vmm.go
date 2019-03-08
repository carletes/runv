package vmm

import (
	"errors"

	"github.com/hyperhq/runv/hypervisor"
	"github.com/hyperhq/runv/hypervisor/types"
)

type VmmDriver struct {
}

type VmmContext struct {
	driver *VmmDriver
}

func InitDriver() *VmmDriver {
	// Check that `vmctl` exists, etc.
	return &VmmDriver{}
}

func (vd *VmmDriver) Name() string {
	return "vmm"
}

func (vd *VmmDriver) InitContext(homeDir string) hypervisor.DriverContext {
	return &VmmContext{
		driver: vd,
	}
}

func (vd *VmmDriver) LoadContext(persisted map[string]interface{}) (hypervisor.DriverContext, error) {
	if t, ok := persisted["hypervisor"]; !ok || t != "empty" {
		return nil, errors.New("wrong driver type in persist info")
	}
	return nil, errors.New("XXX Implement me")
}

func (vc *VmmDriver) SupportLazyMode() bool {
	return false
}

func (vc *VmmDriver) SupportVmSocket() bool {
	return false
}

func (vc *VmmContext) Launch(ctx *hypervisor.VmContext) {
}

func (vc *VmmContext) Associate(ctx *hypervisor.VmContext) {
}

func (vc *VmmContext) Dump() (map[string]interface{}, error) {
	return nil, errors.New("XXX Implement me")
}

func (vc *VmmContext) AddDisk(ctx *hypervisor.VmContext, sourceType string, blockInfo *hypervisor.DiskDescriptor, result chan<- hypervisor.VmEvent) {
}

func (vc *VmmContext) RemoveDisk(ctx *hypervisor.VmContext, blockInfo *hypervisor.DiskDescriptor, callback hypervisor.VmEvent, result chan<- hypervisor.VmEvent) {
}

func (vc *VmmContext) AddNic(ctx *hypervisor.VmContext, host *hypervisor.HostNicInfo, guest *hypervisor.GuestNicInfo, result chan<- hypervisor.VmEvent) {
}

func (vc *VmmContext) RemoveNic(ctx *hypervisor.VmContext, n *hypervisor.InterfaceCreated, callback hypervisor.VmEvent, result chan<- hypervisor.VmEvent) {
}

func (vc *VmmContext) SetCpus(ctx *hypervisor.VmContext, cpus int) error {
	return errors.New("XXX Implement me")
}

func (vc *VmmContext) AddMem(ctx *hypervisor.VmContext, slot, size int) error {
	return errors.New("XXX Implement me")
}

func (vc *VmmContext) Save(ctx *hypervisor.VmContext, path string) error {
	return errors.New("XXX Implement me")
}

func (vc *VmmContext) Shutdown(ctx *hypervisor.VmContext) {
}

func (vc *VmmContext) Kill(ctx *hypervisor.VmContext) {
}

func (vc *VmmContext) Pause(ctx *hypervisor.VmContext, pause bool) error {
	return errors.New("XXX Implement me")
}

func (vc *VmmContext) Stats(ctx *hypervisor.VmContext) (*types.PodStats, error) {
	return nil, errors.New("XXX Implement me")
}

func (vc *VmmContext) Close() {
}
