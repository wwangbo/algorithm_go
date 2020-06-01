package ArrayList

import (
	"errors"
	"fmt"
)

//接口
type List interface {
	Size() int //数组大小
	Get(index int) (interface{},error)//获取第index个元素
	Set(index int,newVal interface{}) error //修改数据
	Insert(index int,newVal interface{}) error
	Append(newVal interface{}) //追加
	Clear() //清空
	Delete(index int) error //删除
	String() string
}

//数据结构
type ArrayList struct {
	dataStore [] interface{} //范型，存储数据
	size int //数组大小
}

func NewArrayList()*ArrayList{
	list := new(ArrayList)
	list.dataStore = make([]interface{},0,10)
	list.size = 0
	return list
}

func (list *ArrayList)Size() int  {
	return list.size  //返回数组的大小
}

func (list *ArrayList) Get(index int) (interface{},error) {
	if index < 0 || index >= list.size {
		return nil,errors.New("索引越界")
	}
	return list.dataStore[index],nil
}

func (list *ArrayList)Append(newVal interface{}) {
	list.dataStore = append(list.dataStore,newVal)
	list.size++
}

func (list *ArrayList) String() string{
	return fmt.Sprint(list)
}