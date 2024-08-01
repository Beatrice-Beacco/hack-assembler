package tests

import (
	"bytes"
	"testing"
)

func assertNotError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Fatalf("Was expecting error, but got nil instead")
	}
}

func assertEqualBoolean(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Fatalf("Expected %t, but got %t instead", want, got)
	}
}

func assertEqualString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Fatalf("Expected %s, but got %s instead", want, got)
	}
}

func assertEqualByteArray(t testing.TB, got, want []byte) {
	t.Helper()

	if !bytes.Equal(got, want) {
		t.Fatalf("Expected\n%s\nBut got\n%s", want, got)
	}
}
