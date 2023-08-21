package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

func main() {
	var b bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		b.WriteString(scanner.Text())
		b.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	output := string(caddyfile.Format(b.Bytes()))

	fmt.Printf("%s", output)
}
