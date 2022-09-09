package osadapter

import (
	"testing"
)

func Test_ToString(t *testing.T) {
	if "osadapter" != ToString() {
		t.Errorf("get pack name error.")
	}
}
