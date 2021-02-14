package worker

import (
	"fmt"
	"strconv"
)

func convertByteToBits(data byte) []uint {
	dst := make([]uint, 0, 8)
	for i := 0; i < 8; i++ {
		move := uint(7 - i)
		dst = append(dst, uint((data>>move)&1))
	}
	return dst
}

func convertBitsToByte(bits []uint) byte {
	str := ""
	for i := 0; i < 8; i++ {
		str = str + fmt.Sprintf("%d", bits[i])
	}
	v, _ := strconv.ParseUint(str, 2, 8)
	return byte(v)
}

func convertBitsToByteArray(bits []uint) []byte {
	byteArray := make([]byte, 0, len(bits)/8)
	for i := 0; i < len(bits)/8; i++ {
		byteArray = append(byteArray, convertBitsToByte(bits[i*8:(i+1)*8]))
	}
	return byteArray
}
