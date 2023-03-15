package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"


	"git.fd.io/govpp.git"
	"git.fd.io/govpp.git/api"

	"github.com/jalapeno/go-vpp-sr/arangodb"

	interfaces "github.com/jalapeno/go-vpp-sr/vppbinapi/interface"
	"github.com/jalapeno/go-vpp-sr/vppbinapi/interface_types"
	"github.com/jalapeno/go-vpp-sr/vppbinapi/ip_types"
	sr "github.com/jalapeno/go-vpp-sr/vppbinapi/sr"
	"github.com/jalapeno/go-vpp-sr/vppbinapi/sr_types"
	"github.com/jalapeno/go-vpp-sr/vppbinapi/vpe"

)

func GetVPPVersion(ch api.Channel) error {
	fmt.Println("Get VPP Version")
	time.Sleep(1 * time.Second)
	request := &vpe.ShowVersion{}
	reply := &vpe.ShowVersionReply{}
	err := ch.SendRequest(request).ReceiveReply(reply)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("VPP Version: %q\n", reply.Version)
	fmt.Println()
	return nil
}

func SrPolicyDump(ch api.Channel) error {
	fmt.Println("Dumping SR Policies")
	time.Sleep(1 * time.Second)

	n := 0
	reqCtx := ch.SendMultiRequest(&sr.SrPoliciesDump{})

	for {
		msg := &sr.SrPoliciesDetails{}
		stop, err := reqCtx.ReceiveReply(msg)
		if stop {
			break
		}
		if err != nil {
			return err
		}
		n++
		fmt.Printf(" - SR Policy #%d: \n", n)
		fmt.Printf("    BSID:      %+v\n", msg.Bsid)
		fmt.Printf("    IsSpray:   %+v\n", msg.IsSpray)
		fmt.Printf("    IsEncap:   %+v\n", msg.IsEncap)
		fmt.Printf("    Fib Table: %+v\n", msg.FibTable)
		fmt.Printf("    SID List:  %+v\n", msg.SidLists[0].Sids)
		//		fmt.Printf("   SID List:  %+v\n", Sids)
	}
	if n == 0 {
		fmt.Println("No Srv6 Policies configured")
	}
	return nil
}

func InterfaceDump(ch api.Channel) error {
	fmt.Println("Dumping interfaces")
	time.Sleep(1 * time.Second)

	n := 0
	reqCtx := ch.SendMultiRequest(&interfaces.SwInterfaceDump{
		SwIfIndex: ^interface_types.InterfaceIndex(0),
	})
	for {
		msg := &interfaces.SwInterfaceDetails{}
		stop, err := reqCtx.ReceiveReply(msg)
		if stop {
			break
		}
		if err != nil {
			return err
		}
		n++
		fmt.Printf(" - interface #%d: %+v\n", n, msg)
		fmt.Printf(" - interface name: %s\n", msg.InterfaceName)
		fmt.Println()
	}
	return nil
}

func ToVppIP6Address(addr net.IP) ip_types.IP6Address {
	ip := [16]uint8{}
	copy(ip[:], addr)
	return ip
}

func ToVppAddress(addr net.IP) ip_types.Address {
	a := ip_types.Address{}
	if addr.To4() == nil {
		a.Af = ip_types.ADDRESS_IP6
		ip := [16]uint8{}
		copy(ip[:], addr)
		a.Un = ip_types.AddressUnionIP6(ip)
	} else {
		a.Af = ip_types.ADDRESS_IP4
		ip := [4]uint8{}
		copy(ip[:], addr.To4())
		a.Un = ip_types.AddressUnionIP4(ip)
	}
	return a
}

func ToVppPrefix(prefix *net.IPNet) ip_types.Prefix {
	len, _ := prefix.Mask.Size()
	r := ip_types.Prefix{
		Address: ToVppAddress(prefix.IP),
		Len:     uint8(len),
	}
	return r
}

func SrSteeringAddDel(ch api.Channel, Bsid ip_types.IP6Address, Traffic ip_types.Prefix) error {
	fmt.Println("Adding SR Steer policy")

	var traffic_type sr_types.SrSteer
	if Traffic.Address.Af == ip_types.ADDRESS_IP4 {
		traffic_type = 4
	} else {
		traffic_type = 6
	}

	request := &sr.SrSteeringAddDel{
		BsidAddr:    Bsid,
		TableID:     0,
		Prefix:      Traffic,
		SwIfIndex:   2,
		TrafficType: traffic_type,
	}

	response := &sr.SrSteeringAddDelReply{}
	err := ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	fmt.Println("SRv6 Steer Policy added!")
	return nil
}

func SrPolicyAdd(ch api.Channel, Bsid ip_types.IP6Address, Isspray bool, Isencap bool, Fibtable int, Sids [16]ip_types.IP6Address, Sidslen int) error {

	fmt.Println("Adding SRv6 Policy")

	BSID := (Bsid)
	PolicyBsid := ip_types.IP6Address{}
	PolicyBsid = BSID
	FwdTable := Fibtable
	FibTable := uint32(FwdTable)

	request := &sr.SrPolicyAdd{
		BsidAddr: PolicyBsid,
		IsSpray:  Isspray,
		IsEncap:  Isencap,
		FibTable: FibTable,
		Sids: sr.Srv6SidList{
			NumSids: uint8(Sidslen),
			Weight:  1,
			Sids:    Sids,
		},
	}
	response := &sr.SrPolicyAddReply{}
	err := ch.SendRequest(request).ReceiveReply(response)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	fmt.Println("SRv6 Policy added: ", Sids)
	return nil
}

func main() {
	// Connect to VPP
	conn, err := govpp.Connect("/var/run/vpp/vpp-api.sock")
	defer conn.Disconnect()
	if err != nil {
		fmt.Printf("Could not connect: %s\n", err)
		os.Exit(1)
	}

	// Open channel
	ch, err := conn.NewAPIChannel()
	defer ch.Close()
	if err != nil {
		fmt.Printf("Could not open API channel: %s\n", err)
		os.Exit(1)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("GoVPP Ready to Rock!")
	time.Sleep(500 * time.Millisecond)


	err = arangodb.Newclient()
	if err != nil {
    		fmt.Printf("New DB Client creation failed: %s\n", err)
		os.Exit(1)
	}
	
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println()
		fmt.Println("Please specify your desired action:")
		fmt.Println("If you want to get VPP detailes, type DET")
		fmt.Println("If you want to add SRv6 policy, type ADD")
		fmt.Println("If you want to add SR Steer policy, type STEER")
		fmt.Println("If you want to show SRv6 policy, type SHOW")
		fmt.Println("If you want to quit, type QUIT")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		if input == "DET" {
			fmt.Println("Getting VPP Details")
			time.Sleep(500 * time.Millisecond)
			err = GetVPPVersion(ch)
			if err != nil {
				fmt.Println("Cannot get the VPP Version: %s\n", err)
				break
			}
			time.Sleep(500 * time.Millisecond)
			err = InterfaceDump(ch)
			if err != nil {
				fmt.Printf("Could not dump interfaces: %s\n", err)
				break
			}
		} else if input == "ADD" {
			fmt.Println("Great! New SRv6 Policy on its way!")
			time.Sleep(1 * time.Second)
			fmt.Println("Please specify the BSID:")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
				continue
			}
			input = strings.TrimSuffix(input, "\n")
			policyBSID := ToVppIP6Address(net.ParseIP(input))
			fmt.Println("BSID: ", input)
			var Isspray bool
			for {
				fmt.Println("Is the policy SPRAY? [Y/N]")
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				input = strings.TrimSuffix(input, "\n")
				if input == "Y" {
					Isspray = true
					break
				} else if input == "N" {
					Isspray = false
					break
				} else {
					fmt.Println("Please type Y or N")
					continue
				}
			}
			var Isencap bool
			for {
				fmt.Println("Is the policy ENCAP? [Y/N]")
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				input = strings.TrimSuffix(input, "\n")
				if input == "Y" {
					Isencap = true
					break
				} else if input == "N" {
					Isencap = false
					break
				} else {
					fmt.Println("Please type Y or N")
					continue
				}
			}
			var Fibtable int
			for {
				fmt.Println("Please specify the FIB Table:")
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				input = strings.TrimSuffix(input, "\n")
				Fibtable, err = strconv.Atoi(input)
				if err != nil {
					fmt.Printf("Please try again. Cannot convert input to integer: %s\n", err)
					continue
				} else {
					break
				}
			}
			segments := [16]ip_types.IP6Address{}
			fmt.Println("Please insert the SID List [empty input will terminate the List]:")
			i := 0
			n := 1
			for {
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				if len(strings.TrimSpace(input)) == 0 {
					input = strings.TrimSuffix(input, "\n")
					break
				}
				input = strings.TrimSuffix(input, "\n")
				segments[i] = ToVppIP6Address(net.ParseIP(input))
				i++
				n++
			}
			fmt.Printf(" - SR Policy Ready to be added:\n")
			fmt.Printf("    BSID:      %+v\n", policyBSID)
			fmt.Printf("    IsSpray:   %+v\n", Isspray)
			fmt.Printf("    IsEncap:   %+v\n", Isencap)
			fmt.Printf("    Fib Table: %+v\n", Fibtable)
			fmt.Printf("    SID List:  %+v\n", segments)
			fmt.Printf("    SID List length:  %+v\n", i)
			fmt.Println()
			for {
				fmt.Println("Please confirm that this is the policy you want to add: [Y/N]")
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					continue
				}
				input = strings.TrimSuffix(input, "\n")
				if input == "Y" {
					//err = SrPolicyAdd(ch)
					err = SrPolicyAdd(ch, policyBSID, Isspray, Isencap, Fibtable, segments, i)
					//err = SrPolicyAdd(ch, policyBSID, Isspray, Isencap, Fibtable, segments, Sidsweight, n)
					if err != nil {
						fmt.Printf("Could not add SR Policy: %s\n", err)
						break
					}
					break
				}
				if input == "N" {
					break
				}
				fmt.Println("Please type Y or N")
				continue
			}
		} else if input == "STEER" {
			fmt.Println("Great! New SR Steer Policy on its way!")
			time.Sleep(1 * time.Second)
			fmt.Println("Please specify the BSID of the SR policy:")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
				continue
			}
			input = strings.TrimSuffix(input, "\n")
			BSID := ToVppIP6Address(net.ParseIP(input))
			fmt.Println("BSID: ", BSID)
			fmt.Println("Please specify the traffic to steer [Addr/Mask]:")
			reader = bufio.NewReader(os.Stdin)
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
				continue
			}
			input = strings.TrimSuffix(input, "\n")
			addr, network, err := net.ParseCIDR(input)
			_ = addr
			traffic := ToVppPrefix(network)
			fmt.Printf("Traffic: %s \n", traffic)
			err = SrSteeringAddDel(ch, BSID, traffic)
			if err != nil {
				fmt.Printf("Could not add SR Steer Policy: %s\n", err)
				continue
			}
			continue
		} else if input == "SHOW" {
			err = SrPolicyDump(ch)
			if err != nil {
				fmt.Printf("Could not dump SR Policies: %s\n", err)
				continue
			}
		} else if input == "QUIT" {
			fmt.Println("Exit the session")
			break
		} else {
			fmt.Println("Sorry, type again please")
			continue
		}
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Goodbye!")
	os.Exit(1)
}
