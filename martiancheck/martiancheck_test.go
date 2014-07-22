package martiancheck

import (
    "net"
    "testing"
)

func TestNewIPv4MartianChecker(t *testing.T) {
  ipch := NewIPv4MartianChecker(true)
}

func TestContains(t *T.testing) {
    ipch := &MartianIPv4Checker{}
    ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0,0,0,0},Mask:maskch.Masks["8"]})
    result := ipch.Contains([]byte{0.0.0.1})
    if result {
        fmt.Println("IP contained within martians block")
    } else {
        t.Fail()
    }
}
