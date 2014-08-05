package hostmanager

import (
	"fmt"
	"net"
	"testing"
)

func TestAddHost(t *testing.T) {
	hm := &HostManager{}
	hm.AddHost("1.2.3.4")
	if hm.Hosts[0].String() == net.ParseIP("1.2.3.4").String() {
		fmt.Println("Host added successfully")
	} else {
		fmt.Println("Host not correctly added")
		t.Fail()
	}
}

func TestAddHostSep(t *testing.T) {
	hm := &HostManager{}
	count := hm.AddHostSep("1.2.3.4", ",")
	if len(hm.Hosts) == 1 && count == 1 {
		if hm.Hosts[0].String() == net.ParseIP("1.2.3.4").String() {
			fmt.Println("Single host successfully added")
		} else {
			fmt.Println("Incorrect host added")
			t.Fail()
		}
	}
	hm2 := &HostManager{}
	count2 := hm2.AddHostSep("1.2.3.4,2.3.4.5", ",")
	if len(hm2.Hosts) == 2 && count2 == 2 {
		if hm2.Hosts[0].String() == net.ParseIP("1.2.3.4").String() && hm2.Hosts[1].String() == net.ParseIP("2.3.4.5").String() {
			fmt.Println("Multiple hosts successfully added")
		} else {
			fmt.Println("Hosts incorrectly added")
			t.Fail()
		}
	} else {
		fmt.Println("Incorrect hosts added")
		t.Fail()
	}
}
