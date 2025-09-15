package dimcli

import (
	"fmt"
	"errors"
	"reflect"
)

type dummy struct {}
var thisPackage = reflect.TypeFor[dummy]().PkgPath()
func PackagePath() string {
	return thisPackage
}

var supportedTypesList = []reflect.Type{
	// special. this is always defaulting to false, and becomes true if
	// supplied.
	reflect.TypeFor[bool](),
	

	// TODO: maybe don't put the optional types here. I think this mapping only
	// has to exist precicely because of this type, so that you can find the
	// reflect.Type based on the string within the fully qualified type name.
	// Not sure though. Will think more about this.

	reflect.TypeFor[int64](),
	reflect.TypeFor[Optional[int64]](),

	reflect.TypeFor[float64](),
	reflect.TypeFor[Optional[float64]](),

	reflect.TypeFor[string](),
	reflect.TypeFor[Optional[string]](),
}
var supportedTypesMap = map[string]reflect.Type{}
func init() {
	for _, t := range supportedTypesList {
		supportedTypesMap[FullTypeNameOf(t)] = t
	}
}

func FullTypeNameOf(t reflect.Type) string {
	return fmt.Sprintf("%s.%s", t.PkgPath(), t.Name())
}

func FullTypeNameFor[T any]() string {
	return FullTypeNameOf(reflect.TypeFor[T]())
}

func Parse[T any]() func() error {

	return func() error {
		return errors.New("TODO")
	}
}
