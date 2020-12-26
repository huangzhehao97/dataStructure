package LinerList

// List接口
type List interface {
	Size() int
	Length() int
	Search(x interface{}) bool
	SetData(i int, x interface{})
	Insert(i int, x interface{})
	Remove(i int, x interface{})
	IsEmpty() bool
	IsFull() bool
	Sort()
	Input()
	Output()
	IsEqual(l List) bool
}



