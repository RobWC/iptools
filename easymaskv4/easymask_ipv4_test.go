package easymaskv4

import (
    "net"
    "testing"
    "fmt"
    "bytes"
)

// TestAddMask testing AddMask function
func TestAddMask(t *testing.T){
    newIPv4Masks := &IPv4Masks{}
    newIPv4Masks.Masks = make(map[uint]net.IPMask)
    err := newIPv4Masks.AddMask(8,[]byte{255,0,0,0})
    if err != nil {
        t.Fail()
    }
    if len(newIPv4Masks.Masks) == 1 {
        fmt.Println("Mask correctly added")
    } else {
        t.Fail()
    }
}

// TestRemoveMask testing RemoveMask function
func TestRemoveMask(t *testing.T) {
    newIPv4Masks := &IPv4Masks{}
    newIPv4Masks.Masks = make(map[uint]net.IPMask)
    newIPv4Masks.Masks[8] = []byte{255,0,0,0}
    err := newIPv4Masks.RemoveMask(8)
    if err != nil {
        t.Fatalf("Error removing mask")
    }
    if len(newIPv4Masks.Masks) == 0 {
        fmt.Println("Mask removed correctly")
    } else {
        t.Fatalf("Mask still present")
    }
}

func TestString(t *testing.T) {
    newIPv4Masks := &IPv4Masks{}
    newIPv4Masks.Masks = make(map[uint]net.IPMask)
    newIPv4Masks.Masks[8] = []byte{255,0,0,0}
    maskString := newIPv4Masks.String(8)
    testMask := net.IPMask([]byte{255,0,0,0})
    if testMask.String() == maskString {
        fmt.Println("Mask correctly converted to string")
    } else {
        t.Fatalf("Mask not correctly converted to string")
    }
}

func TestIPMask(t *testing.T) {
    newIPv4Masks := &IPv4Masks{}
    newIPv4Masks.Masks = make(map[uint]net.IPMask)
    newIPv4Masks.Masks[8] = []byte{255,0,0,0}
    maskIPMask := newIPv4Masks.IPMask(8)
    testMask := net.IPMask([]byte{255,0,0,0})
    if bytes.Compare(maskIPMask, testMask) == 0 {
        fmt.Println("Mask correctly returned as net.IPMask")
    } else {
        t.Fatalf("Mask not correctly generated")
    }
}

func TestNewIPv4Masks(t *testing.T) {
    newIPv4Masks := NewIPv4Masks()
    if len(newIPv4Masks.Masks) == 33 {
        fmt.Println("New IPv4 Masks correctly generated")
    } else {
        t.Fatalf("Masks not correctly initialized")
    }
}
