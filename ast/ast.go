package ast

type Program struct {
	FileName string
	Stmts    []Stmt
}

type Stmt interface{ stmt() }

type Expr interface{ expr() }

type FunDeclStmt struct {
	Name  string
	Proto *Proto
	Body  []Stmt
}
type PackageDecl struct {
	Expr Expr
}
type ImportDecl struct {
	Expr Expr
}
type LetStmt struct {
	Name      Expr
	ValueType TypeInfo
	Value     Expr
}
type WhileStmt struct {
	Cond Expr
	Body Expr
}
type IfStmt struct {
	Cond Expr
	Then []Stmt
	Else []Stmt
}
type ReturnStmt struct {
	Value Expr
}
type ExprStmt struct {
	Expr Expr
}

func (FunDeclStmt) stmt() {}
func (PackageDecl) stmt() {}
func (ImportDecl) stmt()  {}
func (LetStmt) stmt()     {}
func (WhileStmt) stmt()   {}
func (IfStmt) stmt()      {}
func (ReturnStmt) stmt()  {}
func (ExprStmt) stmt()    {}

type BinaryExpr struct {
	Rhs, Lhs Expr
	Op       BinaryOp
}
type UnaryExpr struct {
	Value Expr
	Op    UnaryOp
}
type IdLit struct {
	ID string
}
type NumLit struct {
	Num float32
}
type StrLit struct {
	Str float32
}

func (BinaryExpr) expr() {}
func (UnaryExpr) expr()  {}
func (IdLit) expr()      {}
func (StrLit) expr()     {}
func (NumLit) expr()     {}

type UnaryOp int

const (
	UNARY_NEG UnaryOp = iota
	UNARY_NOT
)

type BinaryOp int

const (
	BINARY_ADD BinaryOp = iota
	BINARY_SUB
	BINARY_MUL
	BINARY_DIV
	BINARY_IDIV
	BINARY_MOD
	BINARY_POW
	BINARY_LT
	BINARY_LE
	BINARY_GT
	BINARY_GE
	BINARY_EQ
	BINARY_NE
	BINARY_AND
	BINARY_OR
	BINARY_NOT
	BINARY_ASSIGN
)

type KTypeKinds int

const (
	TYPE_BUILTIN_NUMBER KTypeKinds = iota
	TYPE_BUILTIN_STRING
	TYPE_BUILTIN_FUNCTION
	TYPE_BUILTIN_ANY
	TYPE_USERTYPE_NORMAL
	TYPE_USERTYPE_GENERIC
)

// semantic analyzer helpers
type Proto struct {
	Base   *FunDeclStmt
	Params []Param
	Return TypeInfo
}
type Param struct {
	Name string
	Type TypeInfo
}
type Attr struct {
	Name string
	Type TypeInfo
}
type TypeInfo struct {
	TypeName    string
	Kind        KTypeKinds
	Attrs       []Attr
	MemberFuns  []Proto
	MemberTypes []TypeInfo
	// TODO: utils functions
}
