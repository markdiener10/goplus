package main

//Provide some of the same code from the unit tests so that a potential user
//Can easily set a break point and step through the code to see how it works

//TODO: Modify the go language server so that VS Code does not report an IDENT or ILLEGAL warnings/errors
//in the IDE UI.  The editor is being given incorrect syntax information from the golang language server

//The AtSign @ is used as our identifier for special operator syntax declarations

type ginterface interface {
	@(parm int) *gstruct
	@[parms ...any] int
	@{parm int} int 
	@+(parm int) int 		
	@^^() bool
	fn(parm int) bool	
}

type gstruct struct {
	fd int
}

//C++ Constructor
//Only can define a single constructor.  Not multiple declarations with different parameter combinations
func (g *gstruct)@(parm int) *gstruct {
	return g
}

//Array Operator Overloading
func (g *gstruct)@[parms ...any] int {
	return 1
}

//If multiple versions of any operator syntax declarations are needed, use the
// ...any variadic form and handle the different types passed in the implementation
// You would lose compiler strict type checking, but you can implement strict type checking
// in the implementation.  Life is full of tradeoffs.

//Bracket Operator Overloading
func (g *gstruct)@{parm int} int {
	return 2
}

//Define + operator
func (g *gstruct)@+(parm int) int {
	return 3
}

//Define ++ operator
func (g *gstruct)@++(parm1 int,parm2 int) int {
	return parm1 + parm2
}

//Define = operator
func (g *gstruct)@=(parm int) int {
	return 4
}

//Traditional function calling semantics
func (g *gstruct)fn(parm int) bool {
	return true
}

//Define ^^ Operator
func (g *gstruct)@^^() bool {
	return true
}

func testinterface(g ginterface) bool {
	return g.^^
}

func main () {

	//C++ Default constructor (But respecting Golang Idiomatic composition paradigm)
	g := gstruct@(3)	

	//Traditional functional calling
	g.fn(4)

	//Operator Overloading for array access
	g.[4]

	//Operator Overloading for bracket access
	g.{5}

	//Classic Operator Overloading Syntax
	g.+ 6

	//Improved Operator Overloading Syntax (Two plusses, Two parameters for example)
	x := g.++ 4,5
	_ = x

	//Now to verify that our C++ Operator overloading and default constructors are accepted by 
	//a defined interface so we can leverage dependency injection and begin to make golang start to have
	//some C++ like behaviors but keep the advantages of golang fully in view
	resb := testinterface(g)
	_ = resb

	
}

