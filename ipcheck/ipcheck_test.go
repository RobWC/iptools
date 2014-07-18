package ipcheck

import (
  "testing"
)

func TestNewIPv4MartianChecker(t *testing.T) {
  ipch := NewIPv4MartianChecker(true)
  if len(ipch.Masks) > 0 {
    if len(ipch.IPUnused) > 0 {
      t.Log("Success on true!")
    } else {
      t.Fatalf("IP Used Unset on true")
    }
  } else {
    t.Fatalf("IP Masks Unset on true")
  }
  ipch2 := NewIPv4MartianChecker(false)
  if len(ipch2.Masks) > 0 {
    if len(ipch2.IPUnused) == 0 {
      t.Log("Success on false!")
    } else {
      t.Fatalf("IP Used Set on false")
    }
  } else {
    t.Fatalf("IP Masks Set on false")
  }
}
