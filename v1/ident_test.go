package ident

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdent(t *testing.T) {
	var id Ident
	fmt.Println(">>>", id)
	assert.Equal(t, true, id.IsZero())
}

func TestSorting(t *testing.T) {
	n := 1000
	check := make([]Ident, n)
	// asc
	for i := 0; i < n; i++ {
		check[i] = Asc()
	}
	for i, e := range check {
		if i > 0 {
			assert.Equal(t, true, bytes.Compare(check[i-1][:], e[:]) < 0, fmt.Sprintf("%v < %v", check, e))
		}
	}
	// dsc
	for i := 0; i < n; i++ {
		check[i] = Dsc()
	}
	for i, e := range check {
		if i > 0 {
			assert.Equal(t, true, bytes.Compare(check[i-1][:], e[:]) > 0, fmt.Sprintf("%v > %v", check, e))
		}
	}
}

func TestOrdering(t *testing.T) {
	a := MustParse("00000000000000000000")
	b := MustParse("A0000000000000000000")
	c := MustParse("00000000000000000000")
	assert.Equal(t, true, a.Before(b), fmt.Sprintf("%v < %v", a, b))
	assert.Equal(t, true, b.After(a), fmt.Sprintf("%v > %v", a, b))
	assert.Equal(t, false, a.Before(c)) // equal
}
