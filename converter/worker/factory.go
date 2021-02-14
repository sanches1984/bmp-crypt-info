package worker

type CryptLevel uint

const (
	CryptLevelLow    = 8
	CryptLevelMedium = 4
	CryptLevelHigh   = 2

	bmpHeaderOffset = 1078
)

type Worker struct {
	level      CryptLevel
	resultData []byte
	position   int
}

func New(level CryptLevel) *Worker {
	return &Worker{level: level}
}

func (w Worker) GetMaxSecretSize(data []byte) int {
	v := (len(data)-bmpHeaderOffset)/int(w.level) - 1
	if v < 0 {
		v = 0
	}
	return v
}
