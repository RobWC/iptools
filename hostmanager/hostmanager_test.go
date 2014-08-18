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
	hm2 := &HostManager{}
	hm2.AddHost([]string{"1.2.3.4","2.2.3.4"}...)
	if hm2.Hosts[0].String() == net.ParseIP("1.2.3.4").String() {
		fmt.Println("Host1 added successfully")
	} else if hm2.Hosts[1].String() == net.ParseIP("2.2.3.4").String() {
		fmt.Println("Host2 added successfully")
	} else {
		fmt.Println("Hosts not correctly added")
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

func TestFindHost(t *testing.T){
	hm := &HostManager{}
	hm.AddHost("1.2.3.4")
	foundHost,err := hm.FindHost(net.ParseIP("1.2.3.4"))
	if foundHost == 0 && err == nil {
		fmt.Println("Host successfully found")
	} else {
		fmt.Println("Unable to find host")
		t.Fail()
	}
}

func TestAddHostIP(t *testing.T) {
	hm := &HostManager{}
	hm.AddHostIP(net.ParseIP("1.2.3.4"))
	if hm.Hosts[0].String() == net.ParseIP("1.2.3.4").String() {
		fmt.Println("Host added successfully")
	} else {
		fmt.Println("Host not correctly added")
		t.Fail()
	}
	hm2 := &HostManager{}
	hm2.AddHostIP([]net.IP{net.ParseIP("1.2.3.4"),net.ParseIP("2.2.3.4")}...)
	if hm2.Hosts[0].String() == net.ParseIP("1.2.3.4").String() {
		fmt.Println("Host1 added successfully")
	} else if hm2.Hosts[1].String() == net.ParseIP("2.2.3.4").String() {
		fmt.Println("Host2 added successfully")
	} else {
		fmt.Println("Hosts not correctly added")
		t.Fail()
	}
}

func TestRemoveHost(t *testing.T) {
	hm := &HostManager{}
	hm.AddHost("1.2.3.4")
	if hm.Hosts[0].String() == net.ParseIP("1.2.3.4").String() {
		fmt.Println("Host added successfully")
		hm.RemoveHost("1.2.3.4")
		if len(hm.Hosts) == 0 {
			fmt.Println("Host removed successfully")
		} else {
			fmt.Println("Host remove failed")
			t.Fail()
		}
	} else {
		fmt.Println("Host not correctly added")
		t.Fail()
	}
	hm2 := &HostManager{}
	hm2.AddHost([]string{"1.2.3.4","2.2.3.4"}...)
	if hm2.Hosts[0].String() == net.ParseIP("1.2.3.4").String() {
		fmt.Println("Host1 added successfully")
	} else if hm2.Hosts[1].String() == net.ParseIP("2.2.3.4").String() {
		fmt.Println("Host2 added successfully")
	} else {
		fmt.Println("Hosts not correctly added")
		t.Fail()
	}

	pos, removed := hm2.RemoveHost([]string{"1.2.3.4","2.2.3.4"}...)
	if removed == 2 {
		fmt.Println("Hosts removed successfully")
	} else if pos[0] == 0 && pos[1] == 1 {
		fmt.Println("Correct host positions removed")
	} else {
		fmt.Println("Hosts remove failed")
		t.Fail()
	}
}

func TestRemoveHostIP(t *testing.T) {
	hm := &HostManager{}
	hm.AddHostIP(net.ParseIP("1.2.3.4"))
	if hm.Hosts[0].String() == net.ParseIP("1.2.3.4").String() {
		fmt.Println("Host added successfully")
		hm.RemoveHost("1.2.3.4")
		if len(hm.Hosts) == 0 {
			fmt.Println("Host removed successfully")
		} else {
			fmt.Println("Host remove failed")
			t.Fail()
		}
	} else {
		fmt.Println("Host not correctly added")
		t.Fail()
	}
	hm2 := &HostManager{}
	hm2.AddHostIP([]net.IP{net.ParseIP("1.2.3.4"),net.ParseIP("2.2.3.4")}...)
	if hm2.Hosts[0].String() == net.ParseIP("1.2.3.4").String() {
		fmt.Println("Host1 added successfully")
	} else if hm2.Hosts[1].String() == net.ParseIP("2.2.3.4").String() {
		fmt.Println("Host2 added successfully")
	} else {
		fmt.Println("Hosts not correctly added")
		t.Fail()
	}

	pos, removed := hm2.RemoveHostIP([]net.IP{net.ParseIP("1.2.3.4"),net.ParseIP("2.2.3.4")}...)
	if removed == 2 {
		fmt.Println("Hosts removed successfully")
	} else if pos[0] == 0 && pos[1] == 1 {
		fmt.Println("Correct host positions removed")
	} else {
		fmt.Println("Hosts remove failed")
		t.Fail()
	}
}

func TestNextHost(t *testing.T) {
	hm := NewHostManager()
	hm.AddHost("1.2.3.4")
	hm.AddHost("2.2.3.4")
	host,curPos := hm.NextHost()
	if host.String() == "2.2.3.4" && curPos == 1 {
		fmt.Println("Next host successfully found")
	} else {
		fmt.Println("Unable to find next host")
		t.Fail()
	}
}

func TestPrevHost(t *testing.T) {
	hm := NewHostManager()
	hm.AddHost("1.2.3.4")
	hm.AddHost("2.2.3.4")
	host,curPos := hm.PrevHost()
	if host.String() == "2.2.3.4" && curPos == 1 {
		fmt.Println("Previous host successfully found")
	} else {
		fmt.Println("Unable to find previous host")
		t.Fail()
	}
}

func TestCurrentHost(t *testing.T) {
	hm := NewHostManager()
	hm.AddHost("1.2.3.4")
	host,curPos := hm.CurrentHost()
	if host.String() == "1.2.3.4" && curPos == 0 {
		fmt.Println("Current host successfully found")
	} else {
		fmt.Println("Unable to find current host")
		t.Fail()
	}
}
