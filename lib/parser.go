package lib

import (
	"errors"
	"reflect"
)

type FGMapStructParser struct {
	tagger string
}

// New instantiate new *FGMapStructParser

// Args: identifier string. Identifier used as the tag value identifier
// from the struct annotation
//
// ex:
//
//	type HolyStruct struct {
//			HolyField string `holyidentifier:"holy_field"`
//	}
//
// As per the example above, the argument when
// calling New would be New("holyidentifier")
func New(identifier string) *FGMapStructParser {
	return &FGMapStructParser{
		tagger: identifier,
	}
}

// Parse
func (fgmsp *FGMapStructParser) Parse(src interface{}, dest any) error {
	return fgmsp.parse(src, dest)
}

func (fgmsp *FGMapStructParser) parse(src interface{}, dest any) error {
	fields := reflect.TypeOf(dest)
	values := reflect.ValueOf(dest)
	fKind := fields.Kind()
	vKind := values.Kind()
	if fKind == vKind && fKind == reflect.Pointer {
		fields = fields.Elem()
		values = values.Elem()
	}

	if fields.Kind() == reflect.Slice {
		for i := 0; i < values.Len(); i++ {
			if err := fgmsp.parse(values.Index(i), dest); err != nil {
				return err
			}
		}
		return nil
	}

	if reflect.TypeOf(src).Kind() != reflect.Map {
		return errParseL0InvalidSrcTypeIsNotMap
	}

	for i := 0; i < fields.NumField(); i += 1 {
		f := fields.Field(i)
		if !f.IsExported() {
			continue
		}
		v := values.Field(i)
		if !v.CanSet() {
			continue
		}

		pursuer := f.Tag.Get(fgmsp.tagger)
		srcVal, ok := src.(map[string]interface{})[pursuer]
		if !ok {
			continue
		}

		if err := fgmsp.setter(f, v, srcVal); err != nil {
			if errors.Is(err, errInvalidType) {
				continue
			}
			return err
		}
	}

	return nil
}
