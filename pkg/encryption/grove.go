package encryption

func (f *File) GetGroveCoordinates() int {
	l := len(*f)
	for i, v := range *f {
		if v.value == 0 {
			return (*f)[(i+1000)%l].value + (*f)[(i+2000)%l].value + (*f)[(i+3000)%l].value
		}
	}
	return 0
}
