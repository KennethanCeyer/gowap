package exception

import "fmt"

// Error is generate new Error to using Pair struct
func Error(p Pair, a ...interface{}) error {
	return fmt.Errorf(fmt.Sprintf("exception(%d): %s", p.Code, p.Message), a...)
}

