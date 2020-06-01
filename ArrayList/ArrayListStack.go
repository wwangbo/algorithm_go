package ArrayList

type StackArray interface {
	Clear() //清空
	Size() int //大小
	Pop() interface{} //弹出
	Push(data interface{}) //押入
	IsFull() bool //是否满了
	IsEmpty() bool //是否为空
}

type Stack struct {
	dataSource *ArrayList
	capSize int //最大范围
}

func NewArrayListStack() *Stack{
	stack := new(Stack)
	stack.dataSource=NewArrayList()
	stack.capSize=10
	return stack
}

func (stack *Stack) Clear() {
	stack.dataSource.Clear()
	stack.capSize=10
}

func (stack *Stack) Size() interface{}{
	return stack.dataSource.Size()
}
func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty(){
		last := stack.dataSource.dataStore[stack.dataSource.Size()-1]
		stack.dataSource.Delete(stack.dataSource.Size()-1)
		return last
	}
	return nil
}
func (stack *Stack) Push(data interface{}) {
	if !stack.IsFull() {
		stack.dataSource.Append(data)
	}
}
func (stack *Stack) IsFull() bool {
	return stack.Size() == stack.capSize
}
func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}