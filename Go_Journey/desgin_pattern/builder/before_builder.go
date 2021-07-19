package main

import (
	"fmt"
	"strings"
)

func before() string {
	sb := strings.Builder{}
	words := []string{"hello", "wolrd"}
	sb.WriteString("<ul>\n")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>\n")
	}
	sb.WriteString("</ul>")
	return sb.String()
}

func main() {
	fmt.Println(before())
}
