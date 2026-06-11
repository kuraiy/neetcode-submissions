type Car struct {
	Position int
	Speed    int
	Time     float32
	Colored  bool
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	for i := range position {
		cars[i].Position = position[i]
		cars[i].Time = float32(target-position[i]) / float32(speed[i])
		cars[i].Speed = speed[i]
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position > cars[j].Position
	})

	// fmt.Println(cars)
	res := 0
	for i := range cars {
		if cars[i].Colored == false {
			for j := i + 1; j < len(cars); j++ {
				if cars[j].Time <= cars[i].Time {
					// fmt.Println("hello", i, j)
					cars[j].Colored = true
				}
			}
		}
	}

	for i := range cars {
		if !cars[i].Colored {
			res++
		}
	}

	return res
}