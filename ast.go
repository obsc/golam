package lam

type Val interface {
	isVal()
}
type Expr interface {
	isExpr()
}

type Var string       // v
func (v Var) isExpr() {}

type Abs struct { // \x.e
	x Var
	e Expr
}

func (e Abs) isExpr() {}
func (e Abs) isVal()  {}

type App struct { // e1 e2
	e1 Expr
	e2 Expr
}

func (e App) isExpr() {}

type Env func(Var) Val
