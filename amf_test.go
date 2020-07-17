package amf

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestAlpha(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	enc := NewEncoder(buf, true)
	dcd := NewDecoder(buf)

	checkErr(t, enc.Encode(map[string]interface{}{
		"abc": "test",
		"ddd": 23,
	}))

	checkErr(t, enc.Encode(uint32(34)))

	vs := []interface{}{}
	for {
		p := new(interface{})
		err := dcd.Decode(p)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
		}
		vs = append(vs, *p)
	}
	fmt.Println(vs)
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
