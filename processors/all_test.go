package processors

import "testing"

func TestHex(t *testing.T) {
	got := "1E (hex) files were added"
	expected := "30 files were added"

	res := Hex(got)
	if res != expected {
		t.Errorf("Expected %q, got %q: ", expected, res)
	}
}

func TestBin(t *testing.T) {
	got := "It has been 10 (bin) years"
	expected := "It has been 2 years"

	res := Bin(got)
	if res != expected {
		t.Errorf("Expected %q, got %q: ", expected, res)
	}
}

func TestUp(t *testing.T) {
	got := "Ready, set, go (up) !"
	expected := "Ready, set, GO !"

	res := Up(got)
	if res != expected {
		t.Errorf("Expected %q, got %q: ", expected, res)
	}
}

func TestLow(t *testing.T) {
	got := "I should stop SHOUTING (low)"
	expected := "I should stop shouting"

	res := Low(got)
	if res != expected {
		t.Errorf("Expected %q, got %q: ", expected, res)
	}
}

func TestCap(t *testing.T) {
	got := "Welcome to the Brooklyn bridge (cap)"
	expected := "Welcome to the Brooklyn Bridge"

	res := Cap(got)
	if res != expected {
		t.Errorf("Expected %q, got %q: ", expected, res)
	}
}

func TestUpN(t *testing.T) {
	got := "This is so exciting (up, 2)"
	expected := "This is SO EXCITING"

	res := UpN(got)
	if res != expected {
		t.Errorf("Expected %q, got %q: ", expected, res)
	}
}

func TestLowN(t *testing.T) {
	got := "IT WAS THE (low, 3) start of fun."
	expected := "it was the start of fun."

	res := LowN(got)
	if res != expected {
		t.Errorf("Expected %q, got %q: ", expected, res)
	}
}

func TestCapN(t *testing.T) {
	got := "it was the age of foolishness (cap, 6)"
	expected := "It Was The Age Of Foolishness"

	res := CapN(got)
	if res != expected {
		t.Errorf("Expected %q, got %q: ", expected, res)
	}
}
