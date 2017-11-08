package tables

import (
	"errors"
	"fmt"
)

const (
	BOOL    = uint16(iota)
	CHAR
	INT
	FLOAT
	DECIMAL
	STRING
)

var TypeLiterals = map[string]uint16{
	"bool":    BOOL,
	"int":     INT,
	"float":   FLOAT,
	"decimal": DECIMAL,
	"string":  STRING,
}

type DataType struct {
	Input []uint16
	Output []uint16
	Literal string
}

//To increment when a new type appears
var typeCount = uint16(len(TypeLiterals));

//Table of type definition
var TypesDefinition = map[uint16]*DataType{
	BOOL: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"bool"},
	INT: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"int"},
	FLOAT: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"float"},
	DECIMAL: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"decimal"},
	STRING: &DataType{Input: []uint16{}, Output: []uint16{}, Literal:"string"},
}

//Table for unresolved types
var unresolvedTypes =  map[string]bool{}

//Add a new type
func AddType(dataType *DataType) error {
	if _, ok := TypeLiterals[dataType.Literal]; ok {
		return errors.New(fmt.Sprintf("Can not redefine type %s.", dataType.Literal))
	}
	typeCount ++;
	//Add a new entry in the type literal table
	TypeLiterals[dataType.Literal] = typeCount
	//Record the type spec
	TypesDefinition[typeCount] = dataType;
	return nil
}

//Get the  uint16 type identifier
func LookupTypeCode(ident string) (uint16, error) {
	if t,ok := TypeLiterals[ident];ok {
		return t,nil
	}
	return 0, errors.New(fmt.Sprintf("Uknown type %s", ident))
}

func LookupTypeName(t uint16) (string,error){
	if t,ok := TypesDefinition[t];ok {
		return t.Literal, nil
	}
	return "unknown", errors.New(fmt.Sprintf("Uknown type %d", t))
}

//Record a type to be resolved at the end of parsing
func LateResolveType(ident string) {
	unresolvedTypes[ident] = true;
}

//Check if there are any unresolved types
func LateTypeResolvingCheck() (string, error) {
	for k, _ := range unresolvedTypes {
		if _,ok := TypeLiterals[k]; !ok {
			return k,errors.New("unresolved type")
		}
	}
	return "",nil
}

//Cleanup the late resolved type map
func LateTypeMapCleanUp() {
	for k := range unresolvedTypes {
		delete(unresolvedTypes, k)
	}
}