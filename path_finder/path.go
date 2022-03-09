package path_finder

import (
	"errors"
)

var (
	errBrokenChain = errors.New("broken chain")
	errBrokenPair  = errors.New("broken pair")
)

// Path Accepts a list of airport tag pairs as input,
// returns a pair from the first and last airport in the chain, or an error
func Path(data [][]string) (path []string, err error) {
	firstPair := data[0]
	left := data[1:]
	if len(left) == 0 {
		return firstPair, nil
	}
	path, _, err = findStep(firstPair, left)

	return
}

// findStep Iterates over the list of pairs recursively.
// For each iteration, it tries to pair the longest path starting from the first pair.
// All pairs that are not included in the expansion generate a list for the next iteration.
func findStep(fl []string, input [][]string) (result []string, output [][]string, err error) {
	result = fl
	for _, pair := range input {
		if len(pair) != 2 {
			return nil, nil, errBrokenPair
		}
		if fl[0] == pair[1] {
			fl[0] = pair[0]
			continue
		} else if fl[1] == pair[0] {
			fl[1] = pair[1]
			continue
		} else {
			output = append(output, pair)
		}
	}

	// If the length of the incoming and subsequent list for iterations is equal,
	// this means that we are in a cycle of a broken chain or a chain with duplicates.
	if len(input) == len(output) {
		return nil, nil, errBrokenChain
	}

	// Iterate until we have elements in the slice for the next iteration
	if len(output) != 0 {
		return findStep(fl, output)
	} else {
		return fl, output, nil
	}
}
