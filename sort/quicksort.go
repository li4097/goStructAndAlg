package sort

type Compare interface {
	Compare1(interface{}) bool
}

func quickSort(s []Compare, left int, right int) {
	if s == nil || left > right {
		return
	}

	flag := s[left]
	low, high := left, right

	for low < high {
		for low < high && s[high].Compare1(flag) == true {
			high--
		}
		if low < high {
			s[low] = s[high]
		}
		for low < high && s[low].Compare1(flag) == false {
			low++
		}
		if low < high {
			s[high] = s[low]
		}
	}
	s[low] = flag
	quickSort(s, left, low-1)
	quickSort(s, high+1, right)

	return
}
func QuickSort(s []Compare) {
	quickSort(s, 0, len(s)-1)
}
