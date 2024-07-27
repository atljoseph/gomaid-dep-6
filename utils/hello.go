package utils

import (
	"fmt"

	"github.com/atljoseph/gomaid-dep-5/http"
)

func SayHelloFrom(name string) {
	fmt.Printf("Hello from %s (w/ gomaid-deps-4)\n", name)
	http.SayHello()
}
