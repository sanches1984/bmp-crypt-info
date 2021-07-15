# BMP Crypt Info
Encrypt information in BMP-image and decrypt it.

## Use cases
### Encrypt
```
./bmp-crypt-info -encrypt -level=normal -src=./test.bmp -dst=./result.bmp -phrase=hello
```
### Decrypt
```
./bmp-crypt-info -decrypt -level=normal -src=./result.bmp
```