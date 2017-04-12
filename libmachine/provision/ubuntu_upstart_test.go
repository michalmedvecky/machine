package provision

import (
	"testing"

	"github.com/michalmedvecky/machine/drivers/fakedriver"
	"github.com/michalmedvecky/machine/libmachine/auth"
	"github.com/michalmedvecky/machine/libmachine/engine"
	"github.com/michalmedvecky/machine/libmachine/provision/provisiontest"
	"github.com/michalmedvecky/machine/libmachine/swarm"
)

func TestUbuntuCompatibleWithHost(t *testing.T) {
	info := &OsRelease{
		ID:        "ubuntu",
		VersionID: "14.04",
	}
	p := NewUbuntuProvisioner(nil)
	p.SetOsReleaseInfo(info)

	compatible := p.CompatibleWithHost()

	if !compatible {
		t.Fatalf("expected to be compatible with ubuntu 14.04")
	}

	info.VersionID = "15.04"

	compatible = p.CompatibleWithHost()

	if compatible {
		t.Fatalf("expected to NOT be compatible with ubuntu 15.04")
	}

}

func TestUbuntuDefaultStorageDriver(t *testing.T) {
	p := NewUbuntuProvisioner(&fakedriver.Driver{}).(*UbuntuProvisioner)
	p.SSHCommander = provisiontest.NewFakeSSHCommander(provisiontest.FakeSSHCommanderOptions{})
	p.Provision(swarm.Options{}, auth.Options{}, engine.Options{})
	if p.EngineOptions.StorageDriver != "aufs" {
		t.Fatal("Default storage driver should be aufs")
	}
}
