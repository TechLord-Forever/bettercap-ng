package network

import (
	"fmt"
	"net"
	"regexp"

	"github.com/evilsocket/bettercap-ng/core"
)

// only matches gateway lines
var IPv4RouteParser = regexp.MustCompile("^(default|[0-9\\.]+)\\svia\\s([0-9\\.]+)\\sdev\\s(\\w+)\\s.*$")
var IPv4RouteTokens = 4
var IPv4RouteCmd = "ip"
var IPv4RouteCmdOpts = []string{"route"}

func IPv4RouteIsGateway(ifname string, tokens []string, f func(gateway string) (*Endpoint, error)) (*Endpoint, error) {
	ifname2 := tokens[3]

	if ifname == ifname2 {
		gateway := tokens[2]
		return f(gateway)
	}

	return nil, nil
}

// see Windows version to understand why ....
func getInterfaceName(iface net.Interface) string {
	return iface.Name
}

func SetInterfaceChannel(iface string, channel int) error {
	out, err := core.Exec("iwconfig", []string{iface, "channel", fmt.Sprintf("%d", channel)})
	if err != nil {
		return err
	} else if out != "" {
		return fmt.Errorf("Unexpected output while setting interface %s to channel %d: %s", iface, channel, out)
	}
	return nil
}
