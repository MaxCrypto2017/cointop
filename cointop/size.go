package cointop

// Size returns window width and height
func (ct *Cointop) size() (int, int) {
	return ct.g.Size()
}

// Width returns window width
func (ct *Cointop) width() int {
	w, _ := ct.size()
	return w
}

// Height returns window height
func (ct *Cointop) height() int {
	_, h := ct.size()
	return h
}

// viewWidth returns view width
func (ct *Cointop) viewWidth(view string) int {
	v, err := ct.g.View(view)
	if err != nil {
		return 0
	}
	w, _ := v.Size()
	return w
}

// viewHeight returns view height
func (ct *Cointop) viewHeight(view string) int {
	v, err := ct.g.View(view)
	if err != nil {
		return 0
	}
	_, h := v.Size()
	return h
}
