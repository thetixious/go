package unicode_name

// AnalyzeName считает байты, Unicode-кодовые точки, латинские и кириллические
// буквы и возвращает байтовые позиции начала каждой Unicode-кодовой точки.
// Предполагается, что входная строка содержит корректный UTF-8.
func AnalyzeName(name string) (byteCount, runeCount, latinCount, cyrillicCount int, positions []int) {
	// TODO: используйте range по строке; для классификации букв можно
	// использовать стандартный пакет unicode.
	return 0, 0, 0, 0, nil
}
