package bioutil

func LeftAlign(ref, alt string) (string, string) {
	r, a := len(ref), len(alt)
	if r<=1 || a<=1 {
		return ref, alt
	}
	var index int
	for i := 1; true; i++ {
		if i == r || i == a {
			index = i - 1
			break
		}

		if ref[r-i] != alt[a-i] {
			index = i - 1
			break
		}
	}

	return ref[0 : r-index], alt[0 : a-index]
}
