package main

import(
	"fmt"
)

type Node interface {
	String() string
}

type Expression interface {
	Node
	isExpression() // dummy functions to force compiler to differentiate between E and S
}

type Statement interface {
	Node
	isStatement() // dummy functions to force compiler to differentiate between E and S
}

type Block []Statement
/*******/
type If struct {
	exp BoolExpression
	trueBlock Block
	falseBlock Block
}
func (i *If) isStatement() {}
func (i *If) String() string {
	return fmt.Sprintf("If:Expr: %v \nTrue block: %v \nFalse block: %v",
						i.exp
						i.trueBlock
						i.falseBlock)
}
/*******/
type Declaration struct {
	t token
	ttype string
	id Id
	value Expression// maybe not idk
}
func (d *Declaration) isStatement() {}
func (d *Declaration) String() string {
	return fmt.Sprintf("Declaration: %v %v = %v", d.ttype, d.id, d.value)
	//idk if this will work since Expression is an interface
}
/*****/
type InfixExpression struct {
	t token
	left Expression
	operator String
	right Expression
}
func (ie *InfixExpression) isExpression() {}
func (ie *InfixExpression) String() string {
	return fmt.Sprintf("InfixExpressions:(%v %v %v)", ie.left, ie.operator, ie.right)
	//idk if this will work since expression is an interface
}
/*****/
type IntExpression struct {
	t token
	value int
}
func (ie *IntExpression) isExpression() {}
func (ie *IntExpression) String() string {
	return fmt.Sprintf("%v", ie.value)
}
/**this would make sense right***/
type FloatExpression struct{
	t token
	value float64
}
func (fe *FloatExpression) isExpression() {}
func (fe *FloatExpression) String() string{
	return fmt.Sprintf("%v", fe.value)
}
/******this should make sense ****/
type StringExpression struct {
	t token
	value string
}
func (se *StringExpressiong) isExpression() {}
func (se *StringExpression) String() string{
	return fmt.Sprintf("%v", se.value)
}


/********/
//idk about this yet
type Function struct {
	t token
	parameters  // not sure
	returnType  // not sure
	name string
	block Block
}
















