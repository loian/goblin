package tables

import (
	"errors"
	"fmt"
)

const (
	BOOL    = uint16(iota)
	CHAR
	INT
	UINT
	BIGINT
	UBIGINT
	FLOAT
	DECIMAL
	STRING
)

var TypeLiterals = map[string]uint16{
	"bool":    BOOL,
	"char":    CHAR,
	"int":     INT,
	"uint":    UINT,
	"bigint":  BIGINT,
	"ubigint": UBIGINT,
	"float":   FLOAT,
	"decimal": DECIMAL,
	"string":  STRING,
}

type DataType struct {
	Input []uint16
	Output []uint16
	Literal string
}

var typeCount = uint16(len(TypeLiterals));

var TypesDefinition = map[uint16]*DataType{
	BOOL: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"bool"},
	CHAR: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"char"},
	INT: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"int"},
	UINT: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"uint"},
	BIGINT: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"bigint"},
	UBIGINT: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"ubigint"},
	FLOAT: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"float"},
	DECIMAL: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"decimal"},
	STRING: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"string"},
}


func AddType(dataType *DataType) error {
	if _, ok := TypeLiterals[dataType.Literal]; ok {
		return errors.New(fmt.Sprintf("Can not redefine type %s.", dataType.Literal))
	}
	typeCount ++;

	TypeLiterals[dataType.Literal] = typeCount
	TypesDefinition[typeCount] = dataType;
	return nil
}

func LookupType(ident string) (uint16, error) {
	if t,ok := TypeLiterals[ident];ok {
		return t,nil
	}
	return 0, errors.New(fmt.Sprintf("Uknown type %s", ident))
}