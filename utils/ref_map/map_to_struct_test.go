package ref_map

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapToStruct(t *testing.T) {
	a := "fsadfasdf"
	v := reflect.ValueOf(a)
	fmt.Println(v.Type())
}
