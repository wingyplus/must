package must

import (
	"errors"
	"testing"
)

func TestMust(t *testing.T) {
	t.Parallel()

	mustEqual(t, 1, Must(1, nil))
	mustEqual(t, "string", Must("string", nil))
	mustEqual(t, struct{}{}, Must(struct{}{}, nil))

	ref := &struct{}{}
	mustEqual(t, ref, Must(ref, nil))
}

func TestMust_Panic(t *testing.T) {
	t.Parallel()

	anerr := errors.New("error: found an error")
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("Must must panic with error")
		}
		if err != anerr {
			t.Fatalf("Expected %v, got %v", anerr, err)
		}
	}()
	Must("", anerr)
	t.Fatal("This test should not reach statement here.")
}

func mustEqual[T comparable](t *testing.T, expected T, actual T) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
