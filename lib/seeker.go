package lib

import (
	"reflect"
	"strconv"
)

func (fgmsp *FGMapStructParser) seekerString(srcV any, v reflect.Value) error {
	x, ok := srcV.(string)
	if !ok {
		return errInvalidType
	}
	v.SetString(x)
	return nil
}

func (fgmsp *FGMapStructParser) seekerBool(srcV any, v reflect.Value) error {
	x, ok := srcV.(bool)
	if !ok {
		return errInvalidType
	}
	v.SetBool(x)
	return nil
}

func (fgmsp *FGMapStructParser) seekerBytes(srcV any, v reflect.Value) error {
	x, ok := srcV.([]byte)
	if !ok {
		return errInvalidType
	}
	v.SetBytes(x)
	return nil
}

func (fgmsp *FGMapStructParser) seekerMap(srcV any, v reflect.Value) error {
	return nil
}

func (fgmsp *FGMapStructParser) seekerFloat(srcV any, v reflect.Value) error {
	switch reflect.TypeOf(srcV).Kind() {
	case reflect.Float32:
		x := srcV.(float32)
		v.SetFloat(float64(x))
	case reflect.Float64:
		x := srcV.(float64)
		v.SetFloat(x)
	case reflect.String:
		x, err := strconv.ParseFloat(srcV.(string), 64)
		if err != nil {
			return err
		}
		v.SetFloat(x)
	default:
		return errInvalidType
	}
	return nil
}

func (fgmsp *FGMapStructParser) seekerInt(srcV any, v reflect.Value) error {
	switch reflect.TypeOf(srcV).Kind() {
	case reflect.Int:
		x := srcV.(int)
		v.SetInt(int64(x))
	case reflect.Int8:
		x := srcV.(int8)
		v.SetInt(int64(x))
	case reflect.Int16:
		x := srcV.(int16)
		v.SetInt(int64(x))
	case reflect.Int32:
		x := srcV.(int32)
		v.SetInt(int64(x))
	case reflect.Int64:
		x := srcV.(int64)
		v.SetInt(x)
	case reflect.String:
		x, err := strconv.Atoi(srcV.(string))
		if err != nil {
			return err
		}
		v.SetInt(int64(x))
	}
	return nil
}

func (fgmsp *FGMapStructParser) seekerUint(srcV any, v reflect.Value) error {
	switch reflect.TypeOf(srcV).Kind() {
	case reflect.Uint:
		x := srcV.(uint)
		v.SetUint(uint64(x))
	case reflect.Uint8:
		x := srcV.(uint8)
		v.SetUint(uint64(x))
	case reflect.Uint16:
		x := srcV.(uint16)
		v.SetUint(uint64(x))
	case reflect.Uint32:
		x := srcV.(uint32)
		v.SetUint(uint64(x))
	case reflect.Uint64:
		x := srcV.(uint64)
		v.SetUint(x)
	case reflect.String:
		x, err := strconv.Atoi(srcV.(string))
		if err != nil {
			return err
		}
		v.SetUint(uint64(x))
	}
	return nil
}
