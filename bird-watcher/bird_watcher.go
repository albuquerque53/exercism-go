package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for _, bird := range birdsPerDay {
		total += bird
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birdsOfWeek []int

	size := len(birdsPerDay)

	weekStart := (week - 1) * 7
	weekEnd := weekStart + 7

	if weekEnd > size {
		birdsOfWeek = birdsPerDay[weekStart:]
	} else {
		birdsOfWeek = birdsPerDay[weekStart:weekEnd]
	}

	return TotalBirdCount(birdsOfWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i++ {
		if i != 0 && i%2 != 0 {
			continue
		}

		birdsPerDay[i]++
	}

	return birdsPerDay
}
