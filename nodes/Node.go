package nodes

type Node interface {
	GetValue() interface{}
	SetValue(interface{})
}
