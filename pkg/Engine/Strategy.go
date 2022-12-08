package Engine

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}
func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type AddPieces struct{}
func (AddPieces) Apply(lval, rval int) int {
	return lval + rval
}

type SubtractPieces struct{}
func (SubtractPieces) Apply(lval, rval int) int {
	return lval - rval
}