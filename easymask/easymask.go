package easymask

import (
  "net"
)

// EasyMask is an interface to provide simple access to multiple types of IPMasks
// This was designed as an interface for both IPv4 and IPv6 support
type EasyMask interface {
  AddMask() error
  RemoveMask() error
  String() unit
  IPMask() net.IPMask
}
