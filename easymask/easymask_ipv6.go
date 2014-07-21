package easymask

import (
    "net"
    "errors"
)

// IPv6Masks contains a map of all the available bitmasks
type IPv6Masks struct {
  Masks map[uint]net.IPMask
}

// AddMask adds a new mask to the IPv6Masks struct
func (ipv6m *IPv6Masks) AddMask(bits uint, mask net.IPMask) error {
  //check if mask exists
  //if so error
  //if not add
  if ipv6m.Masks[bits] == nil {
      ipv6m.Masks[bits] = mask
  } else {
      errors.New("Mask already exists")
  }
  return nil
}

// RemoveMask removes the mask from the IPv6Masks struct
func (ipv6m *IPv6Masks) RemoveMask(bits uint) error {
  //check if mask exists
  //if so remove
  if ipv6m.Masks[bits] != nil {
      delete(ipv6m.Masks,bits)
      return nil
  } else {
      return errors.New("Mask does not exist")
  }
  return nil
}

// String returns the specified mask as a string
func (ipv6m *IPv6Masks) String(bits uint) string {
  rMask := ipv6m.Masks[bits]
  if rMask == nil {
    return ""
  }
  return rMask.String()
}

func (ipv6m *IPv6Masks) IPMask(bits uint) net.IPMask {
    return ipv6m.Masks[bits]
}

// NewIPv6Masks generates an IPv6Masks struct that contains all 129 possible masks
func NewIPv6Masks() ipv6masks {
    ipv6m := IPv6Masks{}
    ipv6m.Masks = make(map[uint]net.IPMask)
    ipv6m.Masks[0] = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[1] = []byte{0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[2] = []byte{0xc0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[3] = []byte{0xe0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[4] = []byte{0xf0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[5] = []byte{0xf8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[6] = []byte{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[7] = []byte{0xfe, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[8] = []byte{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[9] = []byte{0xff, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[10] = []byte{0xff, 0xc0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[11] = []byte{0xff, 0xe0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[12] = []byte{0xff, 0xf0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[14] = []byte{0xff, 0xf8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[14] = []byte{0xff, 0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[15] = []byte{0xff, 0xfe, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[16] = []byte{0xff, 0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[24] = []byte{0xff, 0xff, 0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[32] = []byte{0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[40] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[48] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[56] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[64] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[72] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[80] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0, 0, 0}
    ipv6m.Masks[88] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0, 0}
    ipv6m.Masks[96] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0}
    ipv6m.Masks[104] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0}
    ipv6m.Masks[112] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0}
    ipv6m.Masks[120] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0}
    ipv6m.Masks[128] = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
    return ipv6m
}
