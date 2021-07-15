package worker

import (
	"fmt"
)

func (w *Worker) Encrypt(data, secret []byte) ([]byte, error) {
	maxSize := w.GetMaxSecretSize(data)
	if maxSize < len(secret) {
		return nil, fmt.Errorf("secret is too big, max size %d", maxSize)
	}

	w.position = bmpHeaderOffset + 1
	w.resultData = make([]byte, 0, len(data))
	for _, b := range data {
		w.resultData = append(w.resultData, b)
	}

	w.writeSize(uint16(len(secret)))

	for _, b := range secret {
		w.encryptByte(b)
	}

	return w.resultData, nil
}

func (w *Worker) encryptByte(b byte) {
	sourceBits := convertByteToBits(b)
	for i := 0; i < int(w.level); i++ {
		currentBits := convertByteToBits(w.resultData[w.position])
		for j := 0; j < 8/int(w.level); j++ {
			currentBits[int(w.level)*(j+1)-1] = sourceBits[i*8/int(w.level)+j]
		}
		w.resultData[w.position] = convertBitsToByte(currentBits)
		w.position++
	}
}

func (w *Worker) writeSize(size uint16) {
	var h, l = uint8(size >> 8), uint8(size & 0xff)
	w.resultData[w.position] = byte(h)
	w.position++
	w.resultData[w.position] = byte(l)
	w.position++
}
