// Package parse contiene el análisis léxico de la sintaxis NQL
// (el lenguaje de comandos de NovoDB): partir una línea de entrada
// en tokens respetando comillas, corchetes y llaves.
//
// Este archivo vivía como tokenizer.go dentro de internal/novodb.
// Se movió aquí porque, a diferencia del resto del paquete, no toca
// ningún tipo del motor (Engine, Document, Config, Session, Filter,
// Transaction...): solo depende de strings de la librería estándar.
// Eso lo hace el único archivo que se puede separar en su propio
// paquete sin exportar estado interno ni arriesgar ciclos de
// imports. Ver docs/known-limitations.md para el resto del análisis.
package parse

import "strings"

// Tokenize divide una línea de comando NQL en tokens, respetando
// cadenas entre comillas simples/dobles y el anidamiento de
// corchetes `[]` y llaves `{}` (para no partir arrays/objetos JSON
// embebidos en el comando).
func Tokenize(input string) []string {
	var tokens []string
	var buf strings.Builder
	inStr := false
	var strCh rune
	inBracket := 0
	inBrace := 0

	for _, ch := range input {
		switch {
		case !inStr && (ch == '"' || ch == '\''):
			inStr = true
			strCh = ch
			buf.WriteRune(ch)
		case inStr && ch == strCh:
			inStr = false
			buf.WriteRune(ch)
		case !inStr && ch == '[':
			inBracket++
			buf.WriteRune(ch)
		case !inStr && ch == ']':
			inBracket--
			buf.WriteRune(ch)
		case !inStr && ch == '{':
			inBrace++
			buf.WriteRune(ch)
		case !inStr && ch == '}':
			inBrace--
			buf.WriteRune(ch)
		case !inStr && inBracket == 0 && inBrace == 0 &&
			(ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'):
			if buf.Len() > 0 {
				tokens = append(tokens, buf.String())
				buf.Reset()
			}
		default:
			buf.WriteRune(ch)
		}
	}
	if buf.Len() > 0 {
		tokens = append(tokens, buf.String())
	}
	return tokens
}
