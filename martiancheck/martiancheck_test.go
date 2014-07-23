package martiancheck

import (
    "net"
    "fmt"
    "testing"
)

func TestNewMartianChecker(t *testing.T) {
  ipch := NewMartianChecker(4)
  if len(ipch.Martians) == 14{
      fmt.Println("MartianChecker IPv4 created correctly")
  } else {
      t.Fail()
  }
  ipch2 := NewMartianChecker(6)
  if len(ipch2.Martians) == 11 {
      fmt.Println("MartainChecker IPv6 created correctly")
  }
}

func TestContains(t *testing.T) {
    ipch := &MartianChecker{}
    ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0,0,0,0},Mask:[]byte{128,0,0,0}})
    result := ipch.Contains([]byte{0,0,0,1})
    if result {
        fmt.Println("IP contained within martians block")
    } else {
        t.Fail()
    }
}

func TestAddMartianNet(t *testing.T) {
    ipch := &MartianChecker{}
    ipch.AddMartianNet(net.IPNet{IP:[]byte{0,0,0,0},Mask:[]byte{128,0,0,0}})
    if len(ipch.Martians) == 1 {
        fmt.Println("Martian successfully created")
    } else {
        t.Fail()
    }
}

func TestRemoveMartianNet(t *testing.T) {
    ipch := &MartianChecker{}
    ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0,0,0,0},Mask:[]byte{128,0,0,0}})
    if len(ipch.Martians) == 1 {
        result := ipch.RemoveMartianNet(net.IPNet{IP:[]byte{0,0,0,0},Mask:[]byte{128,0,0,0}})
        if len(ipch.Martians) == 0 && result {
            fmt.Println("Martian net successfully added and removed")
        } else {
            t.Fail()
        }
    } else {
        t.Fail()
    }
}

func TestRemoveAllMartianNet(t *testing.T) {
    ipch := &MartianChecker{}
    ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0,0,0,0},Mask:[]byte{128,0,0,0}})
    ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{1,0,0,0},Mask:[]byte{128,0,0,0}})
    ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{2,0,0,0},Mask:[]byte{128,0,0,0}})
    if len(ipch.Martians) == 3 {
        ipch.RemoveAllMartianNet()
        if len(ipch.Martians) == 0 {
            fmt.Println("All martian nets removed")
        } else {
            t.Fail()
        }
    } else {
        t.Fail()
    }
}
