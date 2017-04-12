package commands

import "github.com/michalmedvecky/machine/libmachine"

func cmdIP(c CommandLine, api libmachine.API) error {
	return runAction("ip", c, api)
}
