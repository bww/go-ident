package ident

import (
	"bytes"
	"sort"
)

type sortable []Ident

func (s sortable) Len() int           { return len(s) }
func (s sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortable) Less(i, j int) bool { return bytes.Compare(s[i][:], s[j][:]) < 0 }

func Sort(v []Ident) {
	sort.Sort(sortable(v))
}
