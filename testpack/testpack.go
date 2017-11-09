package testpack

import (
	"fmt"
	"sort"
)

type IpAlloc struct { //IpAlloc struct
	FreeList []int
}

func new(size int) *IpAlloc {
	return &IpAlloc{
		FreeList: make([]int, size),
	}
}

type Ipcache struct {
	cacheIpV4 []*IpAlloc
	cacheIpV6 []*IpAlloc
}

func New_Ipcache() *Ipcache {
	return &Ipcache{
		cacheIpV4: []*IpAlloc{new(5), new(5)},
		cacheIpV6: []*IpAlloc{new(5), new(5)},
	}
}

func (iplists *Ipcache) Load_ip(ipv4 bool) {
	if ipv4 {
		if len(iplists.cacheIpV4[0].FreeList) == 0 {
			return
		}
		for i := 0; i < 5; i++ {
			iplists.cacheIpV4[0].FreeList[i] = i
		}
	} else {
		for i, j := 0, 5; i < 5; i, j = i+1, j+1 {
			iplists.cacheIpV6[0].FreeList[i] = j
		}
	}
}

func (iplists *Ipcache) GetV4() []*IpAlloc {
	return iplists.cacheIpV4
}

type Test struct {
	a int
}

//func (iplists *Ipcache) Combine() *IpAlloc {

//ranges := iplists.cacheIpV4
//result := new(10)
//for _, r := range ranges {
//fmt.Println(r)
//}
//return result
//}

func (iplists *Ipcache) IsPresent(a int) bool {
	i := sort.Search(len(iplists.cacheIpV4[0].FreeList), func(i int) bool {
		return iplists.cacheIpV4[0].FreeList[i] == a
	})
	if i >= len(iplists.cacheIpV4[0].FreeList) {
		fmt.Printf("\n **** not present")
		return false
	}
	return true
}

func hello() {
	fmt.Printf("hello, Abhi\n")
	//var ex test
	//var podip Ipcache
	//podip := new_Ipcache()

	//ex.a = 7
	//fmt.Println(ex)
	//podip.load_ip(true)
	//podip.load_ip(false)
	/*
	   if len(podip.cacheIpV4[0].FreeList) == 0 {
	       podip.cacheIpV4 = append(podip.cacheIpV4[1:], new(5))
	       fmt.Printf("\n HII I am NIL")
	   }

	   fmt.Println(podip.cacheIpV4[0])
	   fmt.Println(podip.cacheIpV4[1])
	   fmt.Println(podip.cacheIpV6[0])
	*/
}
