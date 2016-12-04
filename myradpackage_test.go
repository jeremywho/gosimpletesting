package myradpackage_test

import (
	"testing"

	"github.com/jeremywho/gosimpletesting"
)

func TestMyRadPackage(t *testing.T) {
	if myradpackage.GetProtocol("http://jeremywho.com") != "http" {
		t.Error("Expected http")
	}
	if myradpackage.GetProtocol("ws://jeremywho.com") != "ws" {
		t.Error("Expected ws")
	}
}
