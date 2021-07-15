package worker

import "encoding/binary"

func (w *Worker) Decrypt(data []byte) ([]byte, error) {
	w.position = bmpHeaderOffset + 1
	size := w.getDataBlockSize(data)

	bitArray := make([]uint, 0)
	for i := 0; i < int(size); i++ {
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

func (w *Worker) getDataBlockSize(data []byte) uint16 {
	var h, l = data[w.position], data[w.position+1]
	w.position += 2

	size := binary.BigEndian.Uint16([]byte{h, l})
	return size * uint16(w.level)
}
