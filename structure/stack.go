package structure

type MyStack struct {
	data  []interface{}
	index int
}

func (self *MyStack) Push(data interface{}) {
	self.data[self.index] = data
	self.index++
}
func (self *MyStack) Pop() interface{} {
	if 0 == self.index {
		return nil
	}
	self.index--
	return self.data[self.index]
}
func (self *MyStack) Empty() bool {
	return self.index == 0
}
