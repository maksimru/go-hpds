package arraylist

type ArrayList interface {
	Get(index int) interface{}
	Set(index int, value interface{})
	Swap(from int, to int)
	Add(value interface{})
	RemoveAt(index int)
	RemoveLast() interface{}
	GetLength() int
	Clean()
}

type AbstractArrayList struct {
	ArrayList
}

func (list AbstractArrayList) Swap(from int, to int) {
	v := list.Get(to)
	list.Set(to, list.Get(from))
	list.Set(from, v)
}

type IntArrayList struct {
	data []int
	AbstractArrayList
}

func NewIntArrayList(arr []int) *IntArrayList {
	list := IntArrayList{arr, AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = &list
	return &list
}

func (list *IntArrayList) Get(index int) interface{} {
	return list.data[index]
}

func (list *IntArrayList) Set(index int, value interface{}) {
	list.data[index] = value.(int)
}

func (list *IntArrayList) RemoveLast() interface{} {
	valueForRemoval := list.data[list.GetLength()-1]
	list.data = list.data[0 : list.GetLength()-1]
	return valueForRemoval
}

func (list *IntArrayList) Add(value interface{}) {
	list.data = append(list.data, value.(int))
}

func (list *IntArrayList) GetLength() int {
	return len(list.data)
}

func (list *IntArrayList) Clean() {
	list.data = nil
}

type Float32ArrayList struct {
	data []float32
	AbstractArrayList
}

func NewFloat32ArrayList(arr []float32) *Float32ArrayList {
	list := Float32ArrayList{arr, AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = &list
	return &list
}

func (list *Float32ArrayList) Get(index int) interface{} {
	return list.data[index]
}

func (list *Float32ArrayList) Set(index int, value interface{}) {
	list.data[index] = value.(float32)
}

func (list *Float32ArrayList) RemoveLast() interface{} {
	valueForRemoval := list.data[list.GetLength()-1]
	list.data = list.data[0 : list.GetLength()-1]
	return valueForRemoval
}

func (list *Float32ArrayList) Add(value interface{}) {
	list.data = append(list.data, value.(float32))
}

func (list *Float32ArrayList) GetLength() int {
	return len(list.data)
}

func (list *Float32ArrayList) Clean() {
	list.data = nil
}

type Float64ArrayList struct {
	data []float64
	AbstractArrayList
}

func NewFloat64ArrayList(arr []float64) *Float64ArrayList {
	list := Float64ArrayList{arr, AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = &list
	return &list
}

func (list *Float64ArrayList) Get(index int) interface{} {
	return list.data[index]
}

func (list *Float64ArrayList) Set(index int, value interface{}) {
	list.data[index] = value.(float64)
}

func (list *Float64ArrayList) RemoveLast() interface{} {
	valueForRemoval := list.data[list.GetLength()-1]
	list.data = list.data[0 : list.GetLength()-1]
	return valueForRemoval
}

func (list *Float64ArrayList) Add(value interface{}) {
	list.data = append(list.data, value.(float64))
}

func (list *Float64ArrayList) GetLength() int {
	return len(list.data)
}

func (list *Float64ArrayList) Clean() {
	list.data = nil
}

type StringArrayList struct {
	data []string
	AbstractArrayList
}

func NewStringArrayList(arr []string) *StringArrayList {
	list := StringArrayList{arr, AbstractArrayList{}}
	list.AbstractArrayList.ArrayList = &list
	return &list
}

func (list *StringArrayList) Get(index int) interface{} {
	return list.data[index]
}

func (list *StringArrayList) Set(index int, value interface{}) {
	list.data[index] = value.(string)
}

func (list *StringArrayList) RemoveLast() interface{} {
	valueForRemoval := list.data[list.GetLength()-1]
	list.data = list.data[0 : list.GetLength()-1]
	return valueForRemoval
}

func (list *StringArrayList) Add(value interface{}) {
	list.data = append(list.data, value.(string))
}

func (list *StringArrayList) GetLength() int {
	return len(list.data)
}

func (list *StringArrayList) Clean() {
	list.data = nil
}
