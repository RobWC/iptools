package martiancheck

import (
    "net"
    "fmt"
    "testing"
)

func TestNewIPv4MartianChecker(t *testing.T) {
  ipch := NewMartianIPv4Checker(true)
  if len(ipch.Martians) == 14{
      fmt.Println("MartianIPv4Checker created correctly")
  } else {
      t.Fail()
  }
}

func TestContains(t *testing.T) {
    ipch := &MartianIPv4Checker{}
    ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0,0,0,0},Mask:[]byte{128,0,0,0}})
    result := ipch.Contains([]byte{0,0,0,1})
    if result {
        fmt.Println("IP contained within martians block")
    } else {
        t.Fail()
    }
}

func TestAddMartianNet(t *testing.T) {
    ipch := &MartianIPv4Checker{}
    ipch.AddMartianNet(net.IPNet{IP:[]byte{0,0,0,0},Mask:[]byte{128,0,0,0}})
    if len(ipch.Martians) == 1 {
        fmt.Println("Martian successfully created")
    } else {
        t.Fail()
    }
}

func TestRemoveMartianNet(t *testing.T) {
    ipch := &MartianIPv4Checker{}
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
    ipch := &MartianIPv4Checker{}
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

/*
func main() {
  ipch := NewIPv4MartianChecker(true)
  maskch := ezmv4.NewIPv4Masks()
  log.Println(len(ipch.Martians))
  log.Println(ipch.Contains([]byte{11,0,0,1}))
  ipch.AddMartianNet(net.IPNet{IP:[]byte{11,0,0,0},Mask:maskch.Masks["24"]})
  log.Println(ipch.Contains([]byte{11,0,0,1}))
  ipch.RemoveMartianNet(net.IPNet{IP:[]byte{11,0,0,0},Mask:maskch.Masks["24"]})
  log.Println(ipch.Contains([]byte{11,0,0,1}))
  ipch.RemoveAllMartianNet()
  log.Println(len(ipch.Martians))
}
*/
