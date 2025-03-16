package log

import "testing"

func TestPrint(t *testing.T) {
	Print()
}

func TestPanic(t *testing.T) {
	Panic()
}

func TestFatal(t *testing.T) {
	Fatal()
}

func TestFlag(t *testing.T) {
	Flag()
}

func TestPrefix(t *testing.T) {
	Prefix()
}
