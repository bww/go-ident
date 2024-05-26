package ident

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeSet(t *testing.T) {
	d1, d2, d3 := New(), New(), New()
	tests := []struct {
		In    Set
		Enc   string
		Error error
	}{
		{
			nil,
			"",
			nil,
		},
		{
			Set{d1},
			d1.String(),
			nil,
		},
		{
			Set{d1, d2, d3},
			fmt.Sprintf("%v,%v,%v", d1, d2, d3),
			nil,
		},
	}
	for _, e := range tests {
		fmt.Println(">>>", e.In)
		enc, err := e.In.MarshalColumn()
		if e.Error != nil {
			fmt.Println("***", err)
			assert.Equal(t, e.Error, err)
		} else if assert.Nil(t, err, fmt.Sprint(err)) {
			fmt.Println("-->", string(enc))
			assert.Equal(t, e.Enc, string(enc))
		}
		if err == nil {
			var dec Set
			err = dec.UnmarshalColumn([]byte(enc))
			if assert.Nil(t, err, fmt.Sprint(err)) {
				fmt.Println("<--", dec)
				assert.Equal(t, e.In, dec)
			}
		}
	}
}
