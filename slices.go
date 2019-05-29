package slices

// This package implements some slice operations not provided by Golang's internal packages.

func RemoveFromStringSlice (stringSlice []string, element string) (newSlice []string) {
/* This function deletes all occurences of a string, from a slice.

	INPUT
	input 0: The slice to be modified.
	input 1: The string whose occurences should be deleted.

	OUTPT
	outpt 0: The result of the operation. */

	for _, memberElement := range stringSlice {
		if element == memberElement {
			continue
		}

		newSlice = append (newSlice, memberElement)
	}

	return newSlice
}

func IsElementInStringSlice (stringSlice []string, element string) (bool) { /* This function
	tells if there is any occurence of a string, in a slice.

	INPUT
	input 0: The slice where the string would be searched.
	input 1: The string to be searched.

	OUTPT
	outpt 0: Value will be "true" if the string exists in the slice. Otherwise, value
		will be false */

	for _, memberElement := range stringSlice {
		if element == memberElement {
			return true
		}
	}

	return false
}

func IndexInStringSlice (stringSlice []string, element string) (int) { /* This function tells
	the index of the first occurence of a string.

	INPUT
	input 0: The slice where the string should be searched.
	input 1: The string whose index should be provided.

	OUTPT
	outpt 0: The index of the first occurence of the input 1. If the input 1 isn't in the
		slice, value would be "-1". */

	for index, memberElement := range stringSlice {
		if element == memberElement {
			return index
		}
	}

	return -1
}
