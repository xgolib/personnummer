package personnummer

import (
	"testing"
)

func TestParse(t *testing.T) {
	for _, pn := range []string{"199405779910", "193408248981"} {
		ret, err := Parse(pn)
		if err != nil {
			t.Error(err)
		}
		t.Log(ret)
	}
}
