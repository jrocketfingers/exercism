package bookstore

var discount map[int]int = map[int]int{
	1: 0,
	2: 5,
	3: 10,
	4: 20,
	5: 25,
}

func Cost(books []int) int {
	count := make(map[int]int)
	for _, book := range books {
		count[book] += 1
	}

	groups := []int{}

	for {
		total := 0
		for k, v := range count {
			if v > 0 {
				count[k] -= 1
				total += 1
			}
		}
		if total == 0 {
			break
		}
		groups = append(groups, total)
	}

	sum := 0
	for _, count := range groups {
		sum += count * 800 * (100 - discount[count]) / 100
	}
	return sum
}
