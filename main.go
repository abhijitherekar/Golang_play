package main

import (
	"fmt"
	"github.com/Golang_play/testpack"
	//	"github.com/stretchr/testify/assert"
)

type testp struct {
	podip *testpack.Ipcache
}

type testcombine struct {
	v4comb *testpack.IpAlloc
}

func main() {
	fmt.Printf("\n hi")
	var ab testp

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
}
