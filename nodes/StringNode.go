package nodes

type StringNode struct {
	value string
}

func (sn *StringNode) GetValue() string {
	return sn.value
}

func (sn *StringNode) SetValue(s string) {
	sn.value = s
}
