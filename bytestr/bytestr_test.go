package bytestr

import (
	"bytes"
	"testing"
)

func TestString2Byte(t *testing.T) {
	s1 := "Hello World!"
	b1 := String2Bytes(s1)
	b2 := []byte(s1)

	if !bytes.Equal(b1, b2) {
		t.Fail()
	}
}

func TestBytes2String(t *testing.T) {
	b1 := []byte("Hello World!")
	s1 := Bytes2String(b1)
	s2 := string(b1)

	if s1 != s2 {
		t.Fail()
	}
}

func Benchmark_NormalBytes2String(b *testing.B) {
	x := []byte("Hello World! Hello World! Hello World!")
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

func Benchmark_Byte2String(b *testing.B) {
	x := []byte("Hello World! Hello World! Hello World!")
	for i := 0; i < b.N; i++ {
		_ = Bytes2String(x)
	}
}

func Benchmark_NormalString2Bytes(b *testing.B) {
	x := "Hello World! Hello World! Hello World!"
	for i := 0; i < b.N; i++ {
		_ = []byte(x)
	}
}

func Benchmark_String2Bytes(b *testing.B) {
	x := "Hello World! Hello World! Hello World!"
	for i := 0; i < b.N; i++ {
		_ = String2Bytes(x)
	}
}
