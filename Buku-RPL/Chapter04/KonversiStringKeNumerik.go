package main

import "fmt"

func main() {
	// Konversi dari int ke float32
	var numInt int = 42
	numFloat32 := float32(numInt)
	fmt.Printf("Hasil konversi int ke float32: %f\n", numFloat32)

	// Konversi dari float64 ke int
	var numFloat64 float64 = 3.14
	numInt2 := int(numFloat64)
	fmt.Printf("Hasil konversi float64 ke int: %d\n", numInt2)

	// Konversi dari int64 ke int32
	var numInt64 int64 = 1000
	numInt32 := int32(numInt64)
	fmt.Printf("Hasil konversi int64 ke int32: %d\n", numInt32)

	// Konversi dari int16 ke int8
	var numInt16 int16 = 255
	numInt8 := int8(numInt16)
	fmt.Printf("Hasil konversi int16 ke int8: %d\n", numInt8)

	// Konversi dari uint8 ke uint16
	var numUint8 uint8 = 200
	numUint16 := uint16(numUint8)
	fmt.Printf("Hasil konversi uint8 ke uint16: %d\n", numUint16)

}
