package randip

import (
  "testing"
)

func TestNewRandIPv4Mgr(t *testing.T) {
  ipv4mgr := NewRandIPv4Mgr(true,0)
  if ipv4mgr.ip == 0 {
    if ipv4mgr.Seq == true {
      t.Log("Success")
    } else {
      t.Fatal("Error setting sequence")
    }
  } else {
    t.Fatal("Error creating IPv4Mgr")
  }
}
