package addr

import (
	"fmt"
	"testing"
)

func TestGetAddr(t *testing.T) {
	addr := GetAddr("120.208.241.144")
	fmt.Println(addr)
}
