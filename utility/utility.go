package utility

type Compare interface {
	Less(c Compare) bool
}
type MyData struct {
	Data int
}

func (i *MyData) Less(data Compare) bool {
	if o, err := data.(*MyData); err == false {
		return false
	} else {
		return i.Data < o.Data
	}
}
