package logger

import "testing"

func TestLogger(t *testing.T) {
	InitLogger()
	Infof("test logger %s", "info")
	Errorf("test logger %s", "error")
	Warnf("test logger %s", "warn")
	Debugf("test logger %s", "debug")
	// Panicf("test logger %s", "panic")
	Fatalf("test logger %s", "fatal")
}
