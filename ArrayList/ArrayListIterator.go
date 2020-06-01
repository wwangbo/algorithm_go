package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool  //是否有下一个
	Next() (interface{},error) //下一个
	Remove()  //删除
	GetIndex() int //获取索引
}

type Iterable interface {
	Iterator() Iterator //构造初始化接口
}
//构造指针访问数组
type ArrayListIterator struct {
	list *ArrayList //数组指针
	currentIndex int //当前索引
}
//构造迭代器
func (list *ArrayList) Iterator() Iterator{
	it := new(ArrayListIterator)
	it.list = list
	it.currentIndex = 0
	return it
}

func (it *ArrayListIterator)HasNext() bool  {
	return it.currentIndex < it.list.size
}

func (it *ArrayListIterator)Next() (interface{},error)  {
	if it.HasNext() {
		value,err := it.list.Get(it.currentIndex)
		it.currentIndex++
		return value,err
	}
	return nil,errors.New("不存在")
}

func (it *ArrayListIterator)Remove()  {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}

func (it *ArrayListIterator) GetIndex() int  {
	return it.currentIndex
}