package main

import (
	"fmt"

	"github.com/Golang_play/stack"
	"github.com/Golang_play/testpack"
	//"github.com/stretchr/testify/assert"
	"github.com/yl2chen/cidranger"
	"net"
)

type testp struct {
	podip *testpack.Ipcache
}

type testcombine struct {
	v4comb *testpack.IpAlloc
}

func parseCIDR(sub string) *net.IPNet {
	_, netw, err := net.ParseCIDR(sub)
	if err == nil {
		return netw
	}
	ip := net.ParseIP(sub)
	if ip == nil {
		return nil
	}
	var mask net.IPMask
	if ip.To4() != nil {
		mask = net.CIDRMask(32, 32)
	} else if ip.To16() != nil {
		mask = net.CIDRMask(128, 128)
	} else {
		return nil
	}
	return &net.IPNet{
		IP:   ip,
		Mask: mask,
	}
}

func main() {
	s := stack.New()
	s.Push(1)
	fmt.Println(s.Peek().(int))
	//fmt.Println(s.Len())
	var ab testp
	ranger := cidranger.NewPCTrieRanger()
	if ranger == nil {
		return
	}
	_, network1, _ := net.ParseCIDR("2001::4/128")
	_, network2, _ := net.ParseCIDR("2001::5/128")
	_, network3, _ := net.ParseCIDR("2001::6/128")
	//_, net4, _ := net.ParseCIDR("192.168.1.1/24")
	ranger.Insert(cidranger.NewBasicRangerEntry(*network1))
	ranger.Insert(cidranger.NewBasicRangerEntry(*network2))
	ranger.Insert(cidranger.NewBasicRangerEntry(*network3))

	net := parseCIDR("2001::/64")
	entries, err := ranger.CoveredNetworks(*net)
	if err != nil {
		fmt.Println("\n ERROR")
	}
	fmt.Println("\n ALFA1 entries are: ", entries)
	for _, entry := range entries {
		fmt.Println("\n the entry for sub is :", entry)
	}

	ab.podip = testpack.New_Ipcache()
	ab.podip.Load_ip(true)
	//var podip testgo.Ipcache
	fmt.Println(ab.podip.GetV4()[0])
	fmt.Println(ab.podip)
	//V4 := ab.podip.Combine()
	//fmt.Println(V4)
	ab.podip.IsPresent(8)
	//combtest := &testcombine{
	//        v4comb : ab.podip.Combine(),
	//    }

	//fmt.Println(combtest.v4comb)

	ch := make(chan string, 2) //buffered channel of size 2
	ch <- "Hello"
	ch <- "World!"
	fmt.Println(<-ch, <-ch)
}
