package commands

import "github.com/michalmedvecky/machine/libmachine"

func cmdUpgrade(c CommandLine, api libmachine.API) error {
	return runAction("upgrade", c, api)
}
