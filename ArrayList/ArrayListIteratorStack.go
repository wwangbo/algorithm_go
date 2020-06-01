package ArrayList

type StackArrayX interface {
	Clear() //清空
	Size() int //大小
	Pop() interface{} //弹出
	Push(data interface{}) //押入
	IsFull() bool //是否满了
	IsEmpty() bool //是否为空
}

type StackX struct {
	dataSource *ArrayList
	pos Iterator
}

func NewArrayListStackX() *StackX{
	stack := new(StackX)
	stack.dataSource=NewArrayList()
	stack.pos = stack.dataSource.Iterator()
	return stack
}

func (stack *StackX) Clear() {
	stack.dataSource.Clear()
	stack.dataSource.size = 0
}

func (stack *StackX) Size() interface{}{
	return stack.dataSource.Size()
}
func (stack *StackX) Pop() interface{} {
	if !stack.IsEmpty(){
		last := stack.dataSource.dataStore[stack.dataSource.Size()-1]
		stack.dataSource.Delete(stack.dataSource.Size()-1)
		return last
	}
	return nil
}
func (stack *StackX) Push(data interface{}) {
	if !stack.IsFull() {
		stack.dataSource.Append(data)
	}
}
func (stack *StackX) IsFull() bool {
	return stack.dataSource.size >= 10
}
func (stack *StackX) IsEmpty() bool {
	return stack.Size() == 0
}