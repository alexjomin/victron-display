package effect

func Translate(buffer []byte, offset, w, h int) []byte {
	size := w * h / 8
	target := []byte{}

	// initialize empty target slice
	for i := 0; i < size; i++ {
		target = append(target, 0)
	}

	// size is incorrect
	if len(buffer) != size {
		return target
	}

	// Translate
	for j, b := range buffer {
		if b == 0 {
			continue
		}
		index := j + offset

		// we need the modulo here as all pixels are proved in a sequential way
		// if fact every elements of the slice contains 8 pixels displayed vertically
		if index > 0 && index%w != 0 {
			target[index] = b
		}
	}
	return target
}
