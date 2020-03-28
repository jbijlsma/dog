// Package dog allows us to more fully understand dogs
package dog

import "time"

// Years turns human years into dog years
func Years(humanYears int) int {

	time.Sleep(2 * time.Second)
	return humanYears * 7
}