package generator

import (
	"reflect"

	"github.com/bww/go-ident/v1"
)

func Distributed() reflect.Value {
	return reflect.ValueOf(ident.New())
}

func Ascending() reflect.Value {
	return reflect.ValueOf(ident.Asc())
}

func Descending() reflect.Value {
	return reflect.ValueOf(ident.Dsc())
}
