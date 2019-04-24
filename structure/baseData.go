package structure

import "struct/utility"

type MyData struct {
	data int
}

func (i *MyData) Less(data utility.Compare) bool {
	if o, err := data.(*MyData); err == false {
		return false
	} else {
		return i.data < o.data
	}
}
