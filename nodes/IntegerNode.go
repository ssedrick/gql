package nodes

type IntegerNode struct {
	value int
}

func (in *IntegerNode) GetValue() int {
	return in.value
}

func (in *IntegerNode) SetValue(i int) {
	in.value = i
}
