package worker

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestConvertToByteArray(t *testing.T) {
	arr := []uint{
		0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 1, 0,
		0, 0, 0, 0, 0, 0, 1, 1,
	}
	actual := convertBitsToByteArray(arr)
	assert.Equal(t, actual, []byte{1, 2, 3})
}

func TestConvertToBitArray(t *testing.T) {
	b := byte(3)
	actual := convertByteToBits(b)
	assert.Equal(t, actual, []uint{0, 0, 0, 0, 0, 0, 1, 1})
}
