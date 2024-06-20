package main

import (
	"fmt"

	"studyum_notifications/internal"
)

func main() {
	fmt.Printf("error launching server: %s", internal.App().Run())
}
