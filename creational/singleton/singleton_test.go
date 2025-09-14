package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.True(t, GetInstance() == GetInstance())
	assert.False(t, GetInstance() == GetLazyInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("test fail")
			}
		}
	})

}

var sink *ConfigSingleton

func BenchmarkGetInstanceParallel2(b *testing.B) {
	Init()
	b.Run("New", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				sink = NewInstance()
			}
		})
	})
	b.Run("Get", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				sink = GetInstance()
			}
		})
	})
}
