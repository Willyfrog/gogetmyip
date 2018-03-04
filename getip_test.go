package gogetmyip

import (
	"testing"
)

// this is not a valid test, just for seeing how it worksó
func TestRun(t *testing.T) {
	GetIp()
}

func TestGetIpKey(t *testing.T) {
	cases := []struct {
		in   map[string]interface{}
		want string
	}{
		{map[string]interface{}{"ip": "1.1.1.1"}, "1.1.1.1"},
		{map[string]interface{}{"ip": "1.1.1.1", "apple": "red"}, "1.1.1.1"},
		{map[string]interface{}{"apple": "red", "ip": "1.1.1.1"}, "1.1.1.1"},
		{map[string]interface{}{"apple": "1.1.1.1"}, ""},
		{map[string]interface{}{"ip": 1}, ""},
	}
	for i, c := range cases {
		got, err := getIpKey(c.in)
		if got != c.want {
			t.Errorf("Case %v got \"%v\", expected %v: %v", i, got, c.want, err)
		}
	}
}
