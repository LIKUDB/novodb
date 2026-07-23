// Command novodb is the entry point for the NovoDB server/CLI.
// All actual logic lives in the internal/novodb package; this file
// intentionally stays thin.
package main

import "novodb/internal/novodb"

func main() {
	novodb.Run()
}
