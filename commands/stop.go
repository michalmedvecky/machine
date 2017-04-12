package commands

import "github.com/michalmedvecky/machine/libmachine"

func cmdStop(c CommandLine, api libmachine.API) error {
	return runAction("stop", c, api)
}
