package martiancheck

import (
  "net"
  "github.com/robwc/easymask/ezmv4"
)

// IPChecker is an interface to check if an IP address is within a martian address range or not.
// This was designed as an interface to allow for both IPv4 and IPv6 to follow the same pattern.
type MartianChecker interface {
  Contains(ip net.IP) bool
  AddMartianNet(net net.IPNet)
  RemoveMartianNet(net net.IPNet) bool
  RemoveAllMartianNet() bool
}

// IPv4Checker to test if an address is within a martian range or not.
type MartianIPv4Checker struct {
  Martians []net.IPNet
}

// NewIPv4MartianChecker returns a configured IPv4 Checker struct.
// Provides all default martian ranges
func NewMartianIPv4Checker(populate bool) MartianIPv4Checker {
  ipch := MartianChecker{}
  maskch := ezmv4.NewIPv4Masks()
  //create unused blocks
  ipch.IPUnused = make([]net.IPNet,0)
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{0,0,0,0},Mask:maskch.Masks["8"]})          //This Net RFC1122
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{10,0,0,0},Mask:maskch.Masks["8"]})         //Private IP RFC1918
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{127,0,0,0},Mask:maskch.Masks["8"]})        //Loopback IP RFC1122
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{169,254,0,0},Mask:maskch.Masks["16"]})     //Local link address
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{127,16,0,0},Mask:maskch.Masks["12"]})      //Private IP RFC1918
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{192,0,0,0},Mask:maskch.Masks["24"]})       //IETF Protocol Assignments RFC5736
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{192,0,2,0},Mask:maskch.Masks["24"]})       //Test-Net-1 RFC5736
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{192,168,0,0},Mask:maskch.Masks["16"]})     //Private IP RFC1918
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{192,18,0,0},Mask:maskch.Masks["15"]})      //Network Testing
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{198,51,100,0},Mask:maskch.Masks["24"]})    //Test-Net-2 RFC5736
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{203,0,113,0},Mask:maskch.Masks["24"]})     //Test-Net-3 RFC5736
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{224,0,0,0},Mask:maskch.Masks["4"]})        //Multicast RFC3171
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{240,0,0,0},Mask:maskch.Masks["4"]})        //Future Use RFC1122
  ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{255,255,255,255},Mask:maskch.Masks["32"]}) //Broadcast
  return ipch
}

func (ipch *IPv4Checker) Contains(ip net.IP) bool {
  for item := range ipch.IPUnused {
    if ipch.IPUnused[item].Contains(ip) {
      return true
    }
  }
  return false
}

func (ipch *IPv4Checker) AddMartianNet(net net.IPNet) {
  ipch.IPUnused = append(ipch.IPUnused,net)
}

func (ipch *IPv4Checker) RemoveMartianNet(net net.IPNet) bool {
  for item := range ipch.IPUnused {
    if net.String() == ipch.IPUnused[item].String() {
      copy(ipch.IPUnused[item:], ipch.IPUnused[item+1:])
      ipch.IPUnused = ipch.IPUnused[:len(ipch.IPUnused)-1]
      return true
    }
  }
  return false
}

func (ipch *IPv4Checker) RemoveAllMartianNet() {
  ipch.IPUnused = make([]net.IPNet,0)
}

func main() {
  ipch := NewIPv4MartianChecker(true)
  maskch := ezmv4.NewIPv4Masks()
  log.Println(len(ipch.IPUnused))
  log.Println(ipch.Contains([]byte{11,0,0,1}))
  ipch.AddMartianNet(net.IPNet{IP:[]byte{11,0,0,0},Mask:maskch.Masks["24"]})
  log.Println(ipch.Contains([]byte{11,0,0,1}))
  ipch.RemoveMartianNet(net.IPNet{IP:[]byte{11,0,0,0},Mask:maskch.Masks["24"]})
  log.Println(ipch.Contains([]byte{11,0,0,1}))
  ipch.RemoveAllMartianNet()
  log.Println(len(ipch.IPUnused))
}
