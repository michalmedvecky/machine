package provision

import (
	"github.com/michalmedvecky/machine/libmachine/auth"
	"github.com/michalmedvecky/machine/libmachine/engine"
)

type EngineConfigContext struct {
	DockerPort       int
	AuthOptions      auth.Options
	EngineOptions    engine.Options
	DockerOptionsDir string
}
