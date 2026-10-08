package textutil

import (
	"errors"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple word", "hola", "aloh"},
		{"sentence with spaces", "hola mundo", "odnum aloh"},
		{"unicode characters", "canción", "nóicnac"},
		{"single character", "a", "a"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.input); got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"lowercase word", "murcielago", 5},
		{"uppercase word", "HOLA", 2},
		{"accented vowels", "canción árbol", 5},
		{"no vowels", "xyz", 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountVowels(tt.input); got != tt.want {
				t.Errorf("CountVowels(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"simple palindrome", "reconocer", true},
		{"phrase with spaces and case", "Anita lava la tina", true},
		{"phrase with punctuation", "A man, a plan, a canal: Panama", true},
		{"numeric palindrome", "12321", true},
		{"single character", "a", true},
		{"not a palindrome", "hola", false},
		{"almost a palindrome", "reconocez", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPalindrome(tt.input)
			if err != nil {
				t.Fatalf("IsPalindrome(%q) unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPalindromeErrors(t *testing.T) {
	inputs := []struct {
		name  string
		input string
	}{
		{"empty string", ""},
		{"only spaces", "   "},
		{"only punctuation", "?!.,"},
	}

	for _, tt := range inputs {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPalindrome(tt.input)
			if !errors.Is(err, ErrEmptyString) {
				t.Errorf("IsPalindrome(%q) error = %v, want %v", tt.input, err, ErrEmptyString)
			}
			if got {
				t.Errorf("IsPalindrome(%q) = true on error, want false", tt.input)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"lowercase word", "hola", "HOLA"},
		{"mixed case", "HoLa Mundo", "HOLA MUNDO"},
		{"accented characters", "canción", "CANCIÓN"},
		{"already uppercase", "ABC", "ABC"},
		{"digits and symbols untouched", "a1-b2", "A1-B2"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.input); got != tt.want {
				t.Errorf("ToUpper(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want string
	}{
		{"two words", "hola", "mundo", "holamundo"},
		{"with separator", "hola ", "mundo", "hola mundo"},
		{"unicode", "ca", "nción", "canción"},
		{"empty first", "", "abc", "abc"},
		{"empty second", "abc", "", "abc"},
		{"both empty", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.a, tt.b); got != tt.want {
				t.Errorf("Concat(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
