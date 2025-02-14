package addr

import (
	"fmt"
	"testing"
)

func TestGetAddr(t *testing.T) {
	addr := GetAddr("")
	fmt.Println(addr)
}
