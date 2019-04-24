package utility

type Compare interface {
	Less(c Compare) bool
}
