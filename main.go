// Copyright (c) TEST USER.
// SPDX-License-Identifier: SSPL-1.0

package main

import (
	"fmt"

	"github.com/hikarukin/incompat/incompatversion"
)

var Version string = "test"

func main() {
	fmt.Printf("Version: %s", Version)
	fmt.Printf("Incompat Version: %s", incompatversion.Version)
}
