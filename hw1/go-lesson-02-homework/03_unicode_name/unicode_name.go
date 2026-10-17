package unicode_name

import "unicode"

// AnalyzeName считает байты, Unicode-кодовые точки, латинские и кириллические
// буквы и возвращает байтовые позиции начала каждой Unicode-кодовой точки.
// Предполагается, что входная строка содержит корректный UTF-8.
func AnalyzeName(name string) (byteCount, runeCount, latinCount, cyrillicCount int, positions []int) {
	byteCount = len(name)
	for i, letter := range name {
		positions = append(positions, i)
		runeCount++
		if unicode.IsLetter(letter) && unicode.Is(unicode.Latin, letter) {
			latinCount++
		}
		if unicode.IsLetter(letter) && unicode.Is(unicode.Cyrillic, letter) {
			cyrillicCount++
		}
	}
	if byteCount == 0 {
		return 0, 0, 0, 0, nil
	}
	return len(name), runeCount, latinCount, cyrillicCount, positions
}
