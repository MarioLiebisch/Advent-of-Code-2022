package distress

func (m *Message) AddDivider() {
	var d1, d2, d1r, d2r Package
	d1r.value = 2
	d2r.value = 6
	d1.value = -1
	d1.contents = make([]*Package, 1)
	d1.contents[0] = &d1r
	d2.value = -1
	d2.contents = make([]*Package, 1)
	d2.contents[0] = &d2r
	m.packages = append(m.packages, &d1, &d2)
}

func (m *Package) IsDivider() bool {
	// Can't be a direct value
	if m.value != -1 {
		return false
	}
	// Must have one child package
	if len(m.contents) != 1 {
		return false
	}
	// Child package's value has to be 2 or 6
	return m.contents[0].value == 2 || m.contents[0].value == 6
}

func (m *Message) GetDecoderKey() int {
	res := 1
	for i, p := range m.packages {
		if p.IsDivider() {
			res *= i + 1
		}
	}
	return res
}
