package lib

import "fmt"

var (
	errParseL0InvalidDestTypeIsArray error = fmt.Errorf("cannot parse dest data from array type")
	errParseL0InvalidSrcTypeIsNotMap error = fmt.Errorf("src type needs to be a map")
	errInvalidType                   error = fmt.Errorf("invalid type")
)
