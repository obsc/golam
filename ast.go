package lam

type Val interface {
	isVal()
}
type Expr interface {
	isExpr()
}

type Var string   // v
type Abs struct { // \x.e
	x Var
	e Expr
}
type App struct { // e1 e2
	e1 Expr
	e2 Expr
}

type Env interface {
	lookup(Var) Val
	extend(Var, Val) Env
}

func (v Var) isExpr() {}
func (e Abs) isExpr() {}
func (e Abs) isVal()  {}
func (e App) isExpr() {}
