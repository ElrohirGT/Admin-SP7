// Package textutil ofrece funciones simples de manipulación de texto.
package textutil

import (
	"errors"
	"strings"
	"unicode"
)

// ErrEmptyString se retorna cuando la operación requiere una cadena no vacía.
var ErrEmptyString = errors.New("la cadena no puede estar vacía")

// Reverse retorna la cadena en orden inverso (respeta caracteres Unicode).
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountVowels retorna el total de vocales (con o sin tilde) de la cadena.
func CountVowels(s string) int {
	count := 0
	for _, r := range strings.ToLower(s) {
		if strings.ContainsRune("aeiouáéíóúü", r) {
			count++
		}
	}
	return count
}

// IsPalindrome indica si la cadena es palíndromo, ignorando mayúsculas,
// espacios y signos de puntuación. Retorna error si no hay caracteres
// alfanuméricos que evaluar.
func IsPalindrome(s string) (bool, error) {
	var cleaned []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}
	if len(cleaned) == 0 {
		return false, ErrEmptyString
	}
	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false, nil
		}
	}
	return true, nil
}

// ToUpper retorna la cadena en mayúsculas.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Concat retorna la unión de dos cadenas.
func Concat(a, b string) string {
	return a + b
}
