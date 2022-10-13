package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	size := len(slice)

	if outOfBound(size, index) {
		return -1
	}

	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	size := len(slice)

	if outOfBound(size, index) {
		return append(slice, value)
	}

	slice[index] = value

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	size := len(slice)

	if outOfBound(size, index) {
		return slice
	}

	firstPart := slice[:index]
	secondPart := slice[index+1:]

	return append(firstPart, secondPart...)
}

// outOfBound verify if the index is out of bound of slice
// "out of bound" = if it is negative or after the end of the stack
func outOfBound(size, index int) bool {
	return index < 0 || index+1 > size
}
