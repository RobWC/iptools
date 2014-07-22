package martiancheck

import (
  "net"
  "github.com/robwc/iptools/easymaskv4"
)

// IPChecker is an interface to check if an IP address is within a martian address range or not.
// This was designed as an interface to allow for both IPv4 and IPv6 to follow the same pattern.
type MartianChecker interface {
  Contains(ip net.IP) bool
  AddMartianNet(net net.IPNet)
  RemoveMartianNet(net net.IPNet) bool
  RemoveAllMartianNet() bool
}

// MartianIPv4Checker to test if an address is within a martian range or not.
type MartianIPv4Checker struct {
  Martians []net.IPNet
}

// NewIPv4MartianChecker returns a configured IPv4 Checker struct.
// Provides all default martian ranges
func NewMartianIPv4Checker(populate bool) MartianIPv4Checker {
  ipch := MartianIPv4Checker{}
  maskch := easymaskv4.NewIPv4Masks()
  //create unused blocks
  ipch.Martians = make([]net.IPNet,0)
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0,0,0,0},Mask:maskch.Masks[8]})          //This Net RFC1122
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{10,0,0,0},Mask:maskch.Masks[8]})         //Private IP RFC1918
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{127,0,0,0},Mask:maskch.Masks[8]})        //Loopback IP RFC1122
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{169,254,0,0},Mask:maskch.Masks[16]})     //Local link address
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{127,16,0,0},Mask:maskch.Masks[12]})      //Private IP RFC1918
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{192,0,0,0},Mask:maskch.Masks[24]})       //IETF Protocol Assignments RFC5736
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{192,0,2,0},Mask:maskch.Masks[24]})       //Test-Net-1 RFC5736
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{192,168,0,0},Mask:maskch.Masks[16]})     //Private IP RFC1918
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{192,18,0,0},Mask:maskch.Masks[15]})      //Network Testing
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{198,51,100,0},Mask:maskch.Masks[24]})    //Test-Net-2 RFC5736
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{203,0,113,0},Mask:maskch.Masks[24]})     //Test-Net-3 RFC5736
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{224,0,0,0},Mask:maskch.Masks[4]})        //Multicast RFC3171
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{240,0,0,0},Mask:maskch.Masks[4]})        //Future Use RFC1122
  ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{255,255,255,255},Mask:maskch.Masks[32]}) //Broadcast
  return ipch
}

func (ipch *MartianIPv4Checker) Contains(ip net.IP) bool {
  for item := range ipch.Martians {
    if ipch.Martians[item].Contains(ip) {
      return true
    }
  }
  return false
}

func (ipch *MartianIPv4Checker) AddMartianNet(net net.IPNet) {
  ipch.Martians = append(ipch.Martians,net)
}

func (ipch *MartianIPv4Checker) RemoveMartianNet(net net.IPNet) bool {
  for item := range ipch.Martians {
    if net.String() == ipch.Martians[item].String() {
      copy(ipch.Martians[item:], ipch.Martians[item+1:])
      ipch.Martians = ipch.Martians[:len(ipch.Martians)-1]
      return true
    }
  }
  return false
}

func (ipch *MartianIPv4Checker) RemoveAllMartianNet() {
  ipch.Martians = make([]net.IPNet,0)
}
