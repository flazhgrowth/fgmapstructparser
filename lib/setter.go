package lib

import (
	"reflect"
)

func (fgmsp *FGMapStructParser) getDestSeekerHandler() map[reflect.Kind]func(srcV interface{}, v reflect.Value) error {
	return map[reflect.Kind]func(srcV interface{}, v reflect.Value) error{
		reflect.String:  fgmsp.seekerString,
		reflect.Bool:    fgmsp.seekerBool,
		reflect.Float32: fgmsp.seekerFloat,
		reflect.Float64: fgmsp.seekerFloat,
		reflect.Int:     fgmsp.seekerInt,
		reflect.Int8:    fgmsp.seekerInt,
		reflect.Int16:   fgmsp.seekerInt,
		reflect.Int32:   fgmsp.seekerInt,
		reflect.Int64:   fgmsp.seekerInt,
		reflect.Uint:    fgmsp.seekerUint,
		reflect.Uint8:   fgmsp.seekerUint,
		reflect.Uint16:  fgmsp.seekerUint,
		reflect.Uint32:  fgmsp.seekerUint,
		reflect.Uint64:  fgmsp.seekerUint,
	}
}

func (fgmsp *FGMapStructParser) setter(f reflect.StructField, v reflect.Value, srcV interface{}) error {
	handler, ok := fgmsp.getDestSeekerHandler()[f.Type.Kind()]
	if !ok {
		v.Set(reflect.ValueOf(srcV))
	}
	if err := handler(srcV, v); err != nil {
		return err
	}

	return nil
}
