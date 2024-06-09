package main

//Provide some of the same code from the unit tests so that a potential user
//Can easily set a break point and step through the code to see how it works

//TODO: Modify the go language server so that VS Code does not report an IDENT or ILLEGAL warnings/errors
//in the IDE UI.  The editor is being given incorrect syntax information from the golang language server
	
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

func (g *gstruct)@(parm int) *gstruct {
	return g
}

func (g *gstruct)@[parms ...any] int {
	return 1
}

func (g *gstruct)@{parm int} int {
	return 2
}

func (g *gstruct)@+(parm int) int {
	return 3
}

func (g *gstruct)@++(parm1 int,parm2 int) int {
	return parm1 + parm2
}

func (g *gstruct)@=(parm int) int {
	return 4
}

func (g *gstruct)fn(parm int) bool {
	return true
}

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

