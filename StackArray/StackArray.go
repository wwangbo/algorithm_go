package StackArray

type StackArray interface {
	Clear() //清空
	Size() int //大小
	Pop() interface{} //弹出
	Push(data interface{}) //押入
	IsFull() bool //是否满了
	IsEmpty() bool //是否为空
}

type Stack struct {
	dataSource [] interface{}
	capSize int //最大范围
	currentSize int //指针，指向当前数组坐标
}

func NewStack() *Stack{
	stack := new(Stack)
	stack.dataSource=make([]interface{},0,10)
	stack.capSize=10
	stack.currentSize=0
	return stack
}

func (stack *Stack) Clear() {
	stack.dataSource=make([]interface{},0,10)
	stack.capSize=10
	stack.currentSize=0
}

func (stack *Stack) Size() interface{}{
	return stack.currentSize
}
func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty(){
		last := stack.dataSource[stack.currentSize-1]
		stack.dataSource=stack.dataSource[:stack.currentSize-1]
		stack.currentSize--
		return last
	}
	return nil
}
func (stack *Stack) Push(data interface{}) {
	if !stack.IsFull() {
		stack.dataSource=append(stack.dataSource,data)
		stack.currentSize++
	}
}
func (stack *Stack) IsFull() bool {
	return stack.currentSize == stack.capSize
}
func (stack *Stack) IsEmpty() bool {
	return stack.currentSize == 0
}