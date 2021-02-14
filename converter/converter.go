package converter

import (
	"github.com/sanches1984/bmp-crypt-info/converter/worker"
	"io/ioutil"
)

type Converter struct {
	level worker.CryptLevel
}

func New(level worker.CryptLevel) *Converter {
	return &Converter{level: level}
}

func (c Converter) DecryptFile(sourceFile string) (string, error) {
	data, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return "", err
	}

	w := worker.New(c.level)
	result, err := w.Decrypt(data)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (c Converter) CryptFile(sourceFile, resultFile, secret string) error {
	data, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	w := worker.New(c.level)
	result, err := w.Encrypt(data, []byte(secret))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(resultFile, result, 0777)
}
