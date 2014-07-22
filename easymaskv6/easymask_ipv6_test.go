package easymaskv6

import (
    "net"
    "testing"
    "fmt"
    "bytes"
)

// TestAddMask testing AddMask function
func TestAddMask(t *testing.T){
    newIPv6Masks := &IPv6Masks{}
    newIPv6Masks.Masks = make(map[uint]net.IPMask)
    err := newIPv6Masks.AddMask(8,[]byte{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
    if err != nil {
        t.Fail()
    }
    if len(newIPv6Masks.Masks) == 1 {
        fmt.Println("Mask correctly added")
    } else {
        t.Fail()
    }
}

// TestRemoveMask testing RemoveMask function
func TestRemoveMask(t *testing.T) {
    newIPv6Masks := &IPv6Masks{}
    newIPv6Masks.Masks = make(map[uint]net.IPMask)
    newIPv6Masks.Masks[8] = []byte{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    err := newIPv6Masks.RemoveMask(8)
    if err != nil {
        t.Fatalf("Error removing mask")
    }
    if len(newIPv6Masks.Masks) == 0 {
        fmt.Println("Mask removed correctly")
    } else {
        t.Fatalf("Mask still present")
    }
}

func TestString(t *testing.T) {
    newIPv6Masks := &IPv6Masks{}
    newIPv6Masks.Masks = make(map[uint]net.IPMask)
    newIPv6Masks.Masks[8] = []byte{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    maskString := newIPv6Masks.String(8)
    testMask := net.IPMask([]byte{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
    if testMask.String() == maskString {
        fmt.Println("Mask correctly converted to string")
    } else {
        t.Fatalf("Mask not correctly converted to string")
    }
}

func TestIPMask(t *testing.T) {
    newIPv6Masks := &IPv6Masks{}
    newIPv6Masks.Masks = make(map[uint]net.IPMask)
    newIPv6Masks.Masks[8] = []byte{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    maskIPMask := newIPv6Masks.IPMask(8)
    testMask := net.IPMask([]byte{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
    if bytes.Compare(maskIPMask, testMask) == 0 {
        fmt.Println("Mask correctly returned as net.IPMask")
    } else {
        t.Fatalf("Mask not correctly generated")
    }
}

func TestNewIPv6Masks(t *testing.T) {
    newIPv6Masks := NewIPv6Masks()
    if len(newIPv6Masks.Masks) == 129 {
        fmt.Println("New IPv6 Masks correctly generated")
    } else {
        t.Fatalf("Masks not correctly initialized")
    }
}
