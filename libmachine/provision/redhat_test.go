package provision

import (
	"testing"

	"github.com/michalmedvecky/machine/drivers/fakedriver"
	"github.com/michalmedvecky/machine/libmachine/auth"
	"github.com/michalmedvecky/machine/libmachine/engine"
	"github.com/michalmedvecky/machine/libmachine/provision/provisiontest"
	"github.com/michalmedvecky/machine/libmachine/swarm"
)

func TestRedHatDefaultStorageDriver(t *testing.T) {
	p := NewRedHatProvisioner("", &fakedriver.Driver{})
	p.SSHCommander = provisiontest.NewFakeSSHCommander(provisiontest.FakeSSHCommanderOptions{})
	p.Provision(swarm.Options{}, auth.Options{}, engine.Options{})
	if p.EngineOptions.StorageDriver != "devicemapper" {
		t.Fatal("Default storage driver should be devicemapper")
	}
}
