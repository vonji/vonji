package controllers

import "strconv"

func parseUint(s string) (uint, error) { //TODO move
	n, err := strconv.ParseUint(s, 10, 64)
	return uint(n), err
}
