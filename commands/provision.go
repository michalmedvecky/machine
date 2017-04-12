package commands

import "github.com/michalmedvecky/machine/libmachine"

func cmdProvision(c CommandLine, api libmachine.API) error {
	return runAction("provision", c, api)
}
