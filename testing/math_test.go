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

func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	expected := 20
	if result != expected {
		t.Errorf("Multiply(4, 5) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(20, 4)
	expected := 5
	if err != nil {
		t.Errorf("Divide(20, 4) returned error: %v", err)
	}
	if result != expected {
		t.Errorf("Divide(20, 4) = %d; want %d", result, expected)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Errorf("Divide(10, 0) should return an error for division by zero")
	}
}

func TestPower(t *testing.T) {
	result := Power(2, 3)
	expected := 8.0
	if result != expected {
		t.Errorf("Power(2, 3) = %f; want %f", result, expected)
	}
}
