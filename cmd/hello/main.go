// Command hello is a minimal demo artifact used to learn SLSA Source Track L3
// and SLSA Build Track L3 hands-on. The logic here is intentionally trivial —
// the point of the exercise is the pipeline and attestations around it, not
// this program.
package main

import "fmt"

// version is set at build time via -ldflags so the built binary can prove,
// via provenance, exactly which commit produced it.
var version = "dev"

func main() {
	fmt.Printf("hello from slsa-l3-demo-app (version=%s)\n", version)
}
