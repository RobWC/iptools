package martiancheck

import (
  "net"
  "github.com/robwc/iptools/easymask"
)
// MartianChecker to test if an address is within a martian range or not.
type MartianChecker struct {
  Martians []net.IPNet
  version uint
}

// NewMartianChecker returns a configured Checker struct.
// Version 4 or 6 can be specified to preconfigure the correct martian blocks
func NewMartianChecker(version uint) MartianChecker {
  ipch := MartianChecker{}
  ipch.version = version
  //create unused blocks
  if version == 4 {
      ipch.Martians = make([]net.IPNet,0)
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0,0,0,0},Mask:easymask.IPv4Mask8})          //This Net RFC1122
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{10,0,0,0},Mask:easymask.IPv4Mask8})         //Private IP RFC1918
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{127,0,0,0},Mask:easymask.IPv4Mask8})        //Loopback IP RFC1122
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{169,254,0,0},Mask:easymask.IPv4Mask16})     //Local link address
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{127,16,0,0},Mask:easymask.IPv4Mask12})      //Private IP RFC1918
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{192,0,0,0},Mask:easymask.IPv4Mask24})       //IETF Protocol Assignments RFC5736
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{192,0,2,0},Mask:easymask.IPv4Mask24})       //Test-Net-1 RFC5736
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{192,168,0,0},Mask:easymask.IPv4Mask16})     //Private IP RFC1918
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{192,18,0,0},Mask:easymask.IPv4Mask15})      //Network Testing
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{198,51,100,0},Mask:easymask.IPv4Mask24})    //Test-Net-2 RFC5736
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{203,0,113,0},Mask:easymask.IPv4Mask24})     //Test-Net-3 RFC5736
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{224,0,0,0},Mask:easymask.IPv4Mask4})        //Multicast RFC3171
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{240,0,0,0},Mask:easymask.IPv4Mask4})        //Future Use RFC1122
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{255,255,255,255},Mask:easymask.IPv4Mask32}) //Broadcast
  } else if version == 6 {
      ipch.Martians = make([]net.IPNet,0)
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},Mask:easymask.IPv6Mask128})               //Node-scoped unicast unspecified address RFC4291
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01},Mask:easymask.IPv6Mask128})            //Node-scoped unicast loopback address RFC4291
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 0, 0, 0, 0},Mask:easymask.IPv6Mask128})         //IPv4-mapped addresses RFC4291
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},Mask:easymask.IPv6Mask96})                //IPv4-compatible addresses RFC4291
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0, 0x64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},Mask:easymask.IPv6Mask64})             //Remotely triggered black hole addresses RFC6666
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0x07, 0xd1, 0, 0x0A, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 0, 0, 0, 0},Mask:easymask.IPv6Mask28}) //ORCHID RFC4843
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0x07, 0xd1, 0x0d, 0xb8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},Mask:easymask.IPv6Mask32})    //Documentation prefix RFC3849
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0xfc, 0x00, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},Mask:easymask.IPv6Mask7})           //ULA RFC4193
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},Mask:easymask.IPv6Mask10})          //Link-local unicast RFC4291
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0xfe, 0xc0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},Mask:easymask.IPv6Mask10})          //Site-local Unicast (depreciated) RFC3879
      ipch.Martians = append(ipch.Martians,net.IPNet{IP:[]byte{0xff, 0x00, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},Mask:easymask.IPv6Mask8})           //Multicast RFC4291
  }
  return ipch
}

func (ipch *MartianChecker) Contains(ip net.IP) bool {
  for item := range ipch.Martians {
    if ipch.Martians[item].Contains(ip) {
      return true
    }
  }
  return false
}

func (ipch *MartianChecker) AddMartianNet(net net.IPNet) {
  ipch.Martians = append(ipch.Martians,net)
}

func (ipch *MartianChecker) RemoveMartianNet(net net.IPNet) bool {
  for item := range ipch.Martians {
    if net.String() == ipch.Martians[item].String() {
      copy(ipch.Martians[item:], ipch.Martians[item+1:])
      ipch.Martians = ipch.Martians[:len(ipch.Martians)-1]
      return true
    }
  }
  return false
}

func (ipch *MartianChecker) RemoveAllMartianNet() {
  ipch.Martians = make([]net.IPNet,0)
}
