package logger

import (
	"io/ioutil"
	"testing"
)

func BenchmarkInfof(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	Init(ioutil.Discard, LStdFlags)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Infof("test message %d", 1)
	}
}

func BenchmarkInfo(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	Init(ioutil.Discard, LStdFlags)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Info("test message 1")
	}
}

func BenchmarkErrorf(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	Init(ioutil.Discard, LStdFlags)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Errorf("test message %d", 1)
	}
}

func BenchmarkError(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	Init(ioutil.Discard, LStdFlags)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Error("test message 1")
	}
}

func BenchmarkWarningf(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	Init(ioutil.Discard, LStdFlags)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Warningf("test message %d", 1)
	}
}

func BenchmarkWarning(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	Init(ioutil.Discard, LStdFlags)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Warning("test message 1")
	}
}
