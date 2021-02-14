package worker

func (w *Worker) Decrypt(data []byte) ([]byte, error) {
	w.position = bmpHeaderOffset + 1
	size := int(data[w.position])
	w.position++

	bitArray := make([]uint, 0)
	for i := 0; i < size; i++ {
		bitArray = append(bitArray, w.getDecryptedBits(data[w.position])...)
		w.position++
	}

	return convertBitsToByteArray(bitArray), nil
}

func (w *Worker) getDecryptedBits(b byte) []uint {
	size := 8 / int(w.level)
	v := make([]uint, 0, size)
	bits := convertByteToBits(b)

	for i := 0; i < size; i++ {
		v = append(v, bits[int(w.level)*(i+1)-1])
	}
	return v
}
