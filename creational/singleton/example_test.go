package singleton

import "testing"

func TestEagerSingleton(t *testing.T) {
	s1 := GetEagerInstance()
	s2 := GetEagerInstance()

	if s1 != s2 {
		t.Error("Expected same instance")
	}

	if s1.GetData() != "eager" {
		t.Error("Expected 'eager'")
	}
}

func TestLazySingleton(t *testing.T) {
	s1 := GetLazyInstance()
	s2 := GetLazyInstance()

	if s1 != s2 {
		t.Error("Expected same instance")
	}

	if s1.GetData() != "lazy" {
		t.Error("Expected 'lazy'")
	}
}
