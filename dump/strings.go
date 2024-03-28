package main

import (
	"fmt"
	"strings"
)

func BuilderInStrings() {
	var sb strings.Builder
	sb.WriteString("Sathish")
	fmt.Println(sb.String())
}
