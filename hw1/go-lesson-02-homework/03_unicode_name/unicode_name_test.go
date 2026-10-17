package unicode_name

import (
	"reflect"
	"testing"
)

func TestAnalyzeName(t *testing.T) {
	tests := []struct {
		name, input                   string
		bytes, runes, latin, cyrillic int
		positions                     []int
	}{
		{"mixed", "Кофе ☕ Go", 15, 9, 2, 4, []int{0, 2, 4, 6, 8, 9, 12, 13, 14}},
		{"ascii", "Go123", 5, 5, 2, 0, []int{0, 1, 2, 3, 4}},
		{"cyrillic", "Яя", 4, 2, 0, 2, []int{0, 2}},
		{"emoji", "🙂", 4, 1, 0, 0, []int{0}},
		{"empty", "", 0, 0, 0, 0, []int{}},
		{"combining_mark", "é", 3, 2, 1, 0, []int{0, 1}},
		{"extended_latin", "ÉéĀā", 8, 4, 4, 0, []int{0, 2, 4, 6}},
		{"extended_cyrillic", "ЁёЇїӁӂ", 12, 6, 0, 6, []int{0, 2, 4, 6, 8, 10}},
		{"other_scripts", "Ω中", 5, 2, 0, 0, []int{0, 2}},
		{"cyrillic_nonletters", "\u0482\u0483", 4, 2, 0, 0, []int{0, 2}},
		{"ascii_boundaries", "@AZ[\x60az{09 \t\n", 13, 13, 4, 0, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
		{"mixed_byte_widths", "Aя☕🙂Z", 11, 5, 2, 1, []int{0, 1, 3, 6, 10}},
		{"latin_nonletter", "\u2160", 3, 1, 0, 0, []int{0}},
		{"three_byte_letters", "\u1e00\u1e01\ua640\ua641", 12, 4, 2, 2, []int{0, 3, 6, 9}},
		{"nul_and_replacement_character", "A\x00\ufffdя", 7, 4, 1, 1, []int{0, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, r, l, c, p := AnalyzeName(tt.input)
			if b != tt.bytes || r != tt.runes || l != tt.latin || c != tt.cyrillic {
				t.Errorf("counts = (%d,%d,%d,%d), want (%d,%d,%d,%d)", b, r, l, c, tt.bytes, tt.runes, tt.latin, tt.cyrillic)
			}
			if len(p) == 0 && len(tt.positions) == 0 {
				return
			}
			if !reflect.DeepEqual(p, tt.positions) {
				t.Errorf("positions = %v, want %v", p, tt.positions)
			}
		})
	}
}
