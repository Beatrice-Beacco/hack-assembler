package tests

import "testing"

func assertNotError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func assertEqualBoolean(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Fatalf("Expected %t, but got %t instead", want, got)
	}
}
