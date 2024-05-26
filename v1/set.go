package ident

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

const (
	sep = ","
)

type Set []Ident

func (v Set) Any() (Ident, bool) {
	if len(v) > 0 {
		return v[0], true
	} else {
		return Zero, false
	}
}

func (v Set) String() string {
	enc, err := v.MarshalColumn()
	if err != nil {
		return fmt.Sprintf("%v (%v)", []Ident(v), err)
	}
	return string(enc)
}

func (v Set) MarshalColumn() ([]byte, error) {
	b := &strings.Builder{}
	for i, e := range v {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(e.String())
	}
	return []byte(b.String()), nil
}

func (v *Set) UnmarshalColumn(data []byte) error {
	s := string(data)
	var a []Ident
	for len(s) > 0 {
		var e string
		if x := strings.Index(s, sep); x >= 0 {
			e, s = strings.TrimSpace(s[:x]), s[x+1:]
		} else {
			e, s = strings.TrimSpace(s), ""
		}
		x, err := Parse(e)
		if err != nil {
			return err
		}
		a = append(a, x)
	}
	*v = a
	return nil
}

func (s Set) Value() (driver.Value, error) {
	var c = make([]string, len(s))
	for i, e := range s {
		c[i] = e.String()
	}
	return pq.Array(c).Value()
}

func (s *Set) Scan(src interface{}) error {
	a := pq.StringArray{}
	err := a.Scan(src)
	if err != nil {
		return err
	}
	r := make(Set, len(a))
	for i, e := range a {
		v, err := Parse(e)
		if err != nil {
			return err
		}
		r[i] = v
	}
	*s = r
	return nil
}
