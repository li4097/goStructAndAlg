package utility

type Compare interface {
	Less(c Compare) bool
}
type MyData struct {
	data int
}

func (i *MyData) Less(data Compare) bool {
	if o, err := data.(*MyData); err == false {
		return false
	} else {
		return i.data < o.data
	}
}
