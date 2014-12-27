package lam

func Eval(env Env, e Expr) Val {
	switch exp := e.(type) {
	case Var:
		return env.lookup(exp)
	case Abs:
		return exp
	case App:
		switch left := Eval(env, exp.e1).(type) {
		case Abs:
			return Eval(env.extend(left.x, Eval(env, exp.e2)), left.e)
		default:
			panic("Illformed Expression")
		}
	default:
		panic("Unable to evaluate expression")
	}
}
