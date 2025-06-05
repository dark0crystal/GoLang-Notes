package testing

import(
	"testing"
)


// === RUN   TestAdd
// --- PASS: TestAdd (0.00s)
// PASS
// ok      GoLang-Notes/testing    1.132s
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// === RUN   TestSubtract
// --- PASS: TestSubtract (0.00s)
// PASS
// ok      GoLang-Notes/testing    1.137s
func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	expected := 6

	if result != expected {
		t.Errorf("Subtract(10, 4) = %d; want %d", result, expected)
	}
}