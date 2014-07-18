package ipcheck

import (
  "net"
  "log"
)

// IPChecker is an interface to check if an IP address is within a martian address range or not.
// This was designed as an interface to allow for both IPv4 and IPv6 to follow the same pattern.
type IPChecker interface {
  Contains(ip net.IP) bool
  AddMartianNet(net net.IPNet)
  RemoveMartianNet(net net.IPNet) bool
  RemoveAllMartianNet() bool
}

// IPv4Checker to test if an address is within a martian range or not.
type IPv4Checker struct {
  IPUnused []net.IPNet
  Masks map[string]net.IPMask
}

// NewIPv4MartianChecker returns a configured IPv4 Checker struct.
// Provides all default martian ranges
func NewIPv4MartianChecker(populate bool) IPv4Checker {
  ipch := IPv4Checker{}
  ipch.Masks = make(map[string]net.IPMask)
  ipch.Masks["Mask0Bit"] = []byte{0,0,0,0}
  ipch.Masks["Mask1Bit"] = []byte{128,0,0,0}
  ipch.Masks["Mask2Bit"] = []byte{192,0,0,0}
  ipch.Masks["Mask3Bit"] = []byte{224,0,0,0}
  ipch.Masks["Mask4Bit"] = []byte{240,0,0,0}
  ipch.Masks["Mask5Bit"] = []byte{248,0,0,0}
  ipch.Masks["Mask6Bit"] = []byte{252,0,0,0}
  ipch.Masks["Mask7Bit"] = []byte{254,0,0,0}
  ipch.Masks["Mask8Bit"] = []byte{255,0,0,0}
  ipch.Masks["Mask9Bit"] = []byte{255,128,0,0}
  ipch.Masks["Mask10Bit"] = []byte{255,192,0,0}
  ipch.Masks["Mask11Bit"] = []byte{255,224,0,0}
  ipch.Masks["Mask12Bit"] = []byte{255,240,0,0}
  ipch.Masks["Mask13Bit"] = []byte{255,248,0,0}
  ipch.Masks["Mask14Bit"] = []byte{255,252,0,0}
  ipch.Masks["Mask15Bit"] = []byte{255,254,0,0}
  ipch.Masks["Mask16Bit"] = []byte{255,255,0,0}
  ipch.Masks["Mask17Bit"] = []byte{255,255,128,0}
  ipch.Masks["Mask18Bit"] = []byte{255,255,192,0}
  ipch.Masks["Mask19Bit"] = []byte{255,255,224,0}
  ipch.Masks["Mask20Bit"] = []byte{255,255,240,0}
  ipch.Masks["Mask21Bit"] = []byte{255,255,248,0}
  ipch.Masks["Mask22Bit"] = []byte{255,255,252,0}
  ipch.Masks["Mask23Bit"] = []byte{255,255,254,0}
  ipch.Masks["Mask24Bit"] = []byte{255,255,255,0}
  ipch.Masks["Mask25Bit"] = []byte{255,255,255,128}
  ipch.Masks["Mask26Bit"] = []byte{255,255,255,192}
  ipch.Masks["Mask27Bit"] = []byte{255,255,255,224}
  ipch.Masks["Mask28Bit"] = []byte{255,255,255,240}
  ipch.Masks["Mask29Bit"] = []byte{255,255,255,248}
  ipch.Masks["Mask30Bit"] = []byte{255,255,255,252}
  ipch.Masks["Mask31Bit"] = []byte{255,255,255,254}
  ipch.Masks["Mask32Bit"] = []byte{255,255,255,255}
  //create unused blocks
  if populate {
    ipch.IPUnused = make([]net.IPNet,0)
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{0,0,0,0},Mask:ipch.Masks["Mask8Bit"]})          //This Net RFC1122
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{10,0,0,0},Mask:ipch.Masks["Mask8Bit"]})         //Private IP RFC1918
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{127,0,0,0},Mask:ipch.Masks["Mask8Bit"]})        //Loopback IP RFC1122
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{169,254,0,0},Mask:ipch.Masks["Mask16Bit"]})     //Local link address
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{127,16,0,0},Mask:ipch.Masks["Mask12Bit"]})      //Private IP RFC1918
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{192,0,0,0},Mask:ipch.Masks["Mask24Bit"]})       //IETF Protocol Assignments RFC5736
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{192,0,2,0},Mask:ipch.Masks["Mask24Bit"]})       //Test-Net-1 RFC5736
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{192,168,0,0},Mask:ipch.Masks["Mask16Bit"]})     //Private IP RFC1918
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{192,18,0,0},Mask:ipch.Masks["Mask15Bit"]})      //Network Testing
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{198,51,100,0},Mask:ipch.Masks["Mask24Bit"]})    //Test-Net-2 RFC5736
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{203,0,113,0},Mask:ipch.Masks["Mask24Bit"]})     //Test-Net-3 RFC5736
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{224,0,0,0},Mask:ipch.Masks["Mask4Bit"]})        //Multicast RFC3171
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{240,0,0,0},Mask:ipch.Masks["Mask4Bit"]})        //Future Use RFC1122
    ipch.IPUnused = append(ipch.IPUnused,net.IPNet{IP:[]byte{255,255,255,255},Mask:ipch.Masks["Mask32Bit"]}) //Broadcast
  }
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
  log.Println(len(ipch.IPUnused))
  log.Println(ipch.Contains([]byte{11,0,0,1}))
  ipch.AddMartianNet(net.IPNet{IP:[]byte{11,0,0,0},Mask:ipch.Masks["Mask24Bit"]})
  log.Println(ipch.Contains([]byte{11,0,0,1}))
  ipch.RemoveMartianNet(net.IPNet{IP:[]byte{11,0,0,0},Mask:ipch.Masks["Mask24Bit"]})
  log.Println(ipch.Contains([]byte{11,0,0,1}))
  ipch.RemoveAllMartianNet()
  log.Println(len(ipch.IPUnused))
}
