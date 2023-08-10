package singleton_test

import (
	"testing"

	singleton "github.com/EASYGOING45/DesignPattern_Golang/Singleton"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.True(t, singleton.GetInstance() == singleton.GetInstance())
	assert.False(t, singleton.GetInstance() == singleton.GetLazyInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
