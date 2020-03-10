package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/mapadj/go-utils/src/logger"
)

// the net.Interface errors:
var (
	errInvalidInterface         = errors.New("invalid network interface")
	errInvalidInterfaceIndex    = errors.New("invalid network interface index")
	errInvalidInterfaceName     = errors.New("invalid network interface name")
	errNoSuchInterface          = errors.New("no such network interface")
	errNoSuchMulticastInterface = errors.New("no such multicast network interface")
)

//the Interface Struct:
/*
type Interface struct {
	Index        int          // positive integer that starts at one, zero is never used
	MTU          int          // maximum transmission unit
	Name         string       // e.g., "en0", "lo0", "eth0.100"
	HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast
}
*/
// the flags
const (
	FlagUp           Flags = 1 << iota // interface is up
	FlagBroadcast                      // interface supports broadcast access capability
	FlagLoopback                       // interface is a loopback interface
	FlagPointToPoint                   // interface belongs to a point-to-point link
	FlagMulticast                      // interface supports multicast access capability
)

// Type Flags:
type Flags uint

//Flag Names:
var flagNames = []string{
	"up",
	"broadcast",
	"loopback",
	"pointtopoint",
	"multicast",
}

// type Flags has a String() function.

// Addrs returns a list of unicast interface adresses for a specifc interface
//func (ifi *Interface) Addrs() ([]Addr, error)

// MulticastAddrs returns a list of multicast, joined group adresses for a
// specific interface.
//func (ifi *Interface) MulticastAddres() ([]Addr, error)

// Interface returns a list of System unicast interface adresses
//func Interfaces() ([]Interface, error)

// return a list of systems unicast interface adresses
// list does not identify interface. use Interfaces and Interface.Addrs fm detail
//func InterfaceAddrs()([]Addr, error)

//func InterfaceByIndex(index int)(*Interface, error)

// InterfaceByName returns the interface specified by name!
//func InterfaceByName(name string)(*Interface, error)

// plus some more ip6 stuff...

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		logger.Error("Ooops", err)
	}
	for _, iface := range ifaces {
		//fmt.Println(addr)
		fmt.Printf("   Index: %-8s", iface.Index)
		fmt.Printf("   Name: %-8s", iface.Name)
		fmt.Printf("   MTU: %-8s", iface.MTU)
		fmt.Printf("   Hardware Address: %-8s", iface.HardwareAddr)
		fmt.Println(iface.Flags.String())
	}
}
