package object

import (
	"bytes"
	"fmt"
	"strings"

	"example.com/ast"
)

type ObjectType string
type BuiltinFunction func(args ...Object) Object

const (
	INTEGEROBJ     = "INTEGER"
	BOOLEANOBJ     = "BOOLEAN"
	NULLOBJ        = "NULL"
	RETURNVALUEOBJ = "RETURN_VALUE"
	ERROROBJ       = "ERROR"
	FUNCTIONOBJ    = "FUNCTION"
	STRINGOBJ      = "STRING"
	BUILTINOBJ     = "BUILTIN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGEROBJ
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEANOBJ
}

type Null struct{}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() ObjectType {
	return NULLOBJ
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURNVALUEOBJ
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROROBJ
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType {
	return FUNCTIONOBJ
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return STRINGOBJ
}

func (s *String) Inspect() string {
	return s.Value
}

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTINOBJ }
func (b *Builtin) Inspect() string { return "builtin function" }
