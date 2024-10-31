package pwd

import "testing"

func TestHashPassword(t *testing.T) {
	res := HashPassword("meowrain")
	t.Log(res)
}

func TestComposePassword(t *testing.T) {
	res := HashPassword("meowrain")
	is_equal := ComparePassword(res, "meowrain")
	if is_equal {
		t.Log("相等")
	}
}
