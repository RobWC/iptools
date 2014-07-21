package easymask

import (
  "net"
  "errors"
)

// IPv4Masks contains a map of all the available bitmasks
type IPv4Masks struct {
  Masks map[uint]net.IPMask
}

// AddMask adds a new mask to the IPv4Masks struct
func (ipv4m *IPv4Masks) AddMask(bits uint, mask net.IPMask) error {
  //check if mask exists
  //if so error
  //if not add
  if ipv4m.Masks[bits] == nil {
      ipv4m.Masks[bits] = mask
  } else {
      errors.New("Mask already exists")
  }
  return nil
}

// RemoveMask removes the mask from the IPv4Masks struct
func (ipv4m *IPv4Masks) RemoveMask(bits uint) error {
  //check if mask exists
  //if so remove
  if ipv4m.Masks[bits] != nil {
      delete(ipv4m.Masks,bits)
      return nil
  } else {
      return errors.New("Mask does not exist")
  }
  return nil
}

// String returns the specified mask as a string
func (ipv4m *IPv4Masks) String(bits uint) string {
  rMask := ipv4m.Masks[bits]
  if rMask == nil {
    return ""
  }
  return rMask.String()
}

// IPMask returns the specified mask as an net.IPMask
func (ipv4m *IPv4Masks) IPMask(bits uint) net.IPMask {
    return ipv4m.Masks[bits]
}

// NewIPv4Masks generates an IPv4Masks struct that contains all 33 possible masks
func NewIPv4Masks() IPv4Masks {
  ipv4m := IPv4Masks{}
  ipv4m.Masks = make(map[uint]net.IPMask)
  ipv4m.Masks[0] = []byte{0,0,0,0}
  ipv4m.Masks[1] = []byte{128,0,0,0}
  ipv4m.Masks[2] = []byte{192,0,0,0}
  ipv4m.Masks[3] = []byte{224,0,0,0}
  ipv4m.Masks[4] = []byte{240,0,0,0}
  ipv4m.Masks[5] = []byte{248,0,0,0}
  ipv4m.Masks[6] = []byte{252,0,0,0}
  ipv4m.Masks[7] = []byte{254,0,0,0}
  ipv4m.Masks[8] = []byte{255,0,0,0}
  ipv4m.Masks[9] = []byte{255,128,0,0}
  ipv4m.Masks[10] = []byte{255,192,0,0}
  ipv4m.Masks[11] = []byte{255,224,0,0}
  ipv4m.Masks[12] = []byte{255,240,0,0}
  ipv4m.Masks[13] = []byte{255,248,0,0}
  ipv4m.Masks[14] = []byte{255,252,0,0}
  ipv4m.Masks[15] = []byte{255,254,0,0}
  ipv4m.Masks[16] = []byte{255,255,0,0}
  ipv4m.Masks[17] = []byte{255,255,128,0}
  ipv4m.Masks[18] = []byte{255,255,192,0}
  ipv4m.Masks[19] = []byte{255,255,224,0}
  ipv4m.Masks[20] = []byte{255,255,240,0}
  ipv4m.Masks[21] = []byte{255,255,248,0}
  ipv4m.Masks[22] = []byte{255,255,252,0}
  ipv4m.Masks[23] = []byte{255,255,254,0}
  ipv4m.Masks[24] = []byte{255,255,255,0}
  ipv4m.Masks[25] = []byte{255,255,255,128}
  ipv4m.Masks[26] = []byte{255,255,255,192}
  ipv4m.Masks[27] = []byte{255,255,255,224}
  ipv4m.Masks[28] = []byte{255,255,255,240}
  ipv4m.Masks[29] = []byte{255,255,255,248}
  ipv4m.Masks[30] = []byte{255,255,255,252}
  ipv4m.Masks[31] = []byte{255,255,255,254}
  ipv4m.Masks[32] = []byte{255,255,255,255}
  return ipv4m
}
