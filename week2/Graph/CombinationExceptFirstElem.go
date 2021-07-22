package Graph

type CombinationExceptFirstElem struct {
	Elements       []int
	ChosenElements [][]int
	M              int

	currentIndices      []int
	hasIterationStopped bool
}

func (c *CombinationExceptFirstElem) nextIndices() []int {
	if c.currentIndices == nil {
		c.currentIndices = make([]int, c.M)

		for i := 0; i < len(c.currentIndices); i++ {
			c.currentIndices[i] = i
		}

		c.hasIterationStopped = false
		return c.currentIndices
	}

	// determine which rightmost index to increase
	// we keep the first element (first city) unchanged
	rightmostToIncrease := -1
	for i := len(c.currentIndices) - 1; i >= 1; i-- {
		atMost := (len(c.Elements) - 1) + (-c.M + 1 + i)
		if c.currentIndices[i] < atMost {
			rightmostToIncrease = i
			break
		}
	}

	// increase the rightmost, and the right indices to the right of rightmost
	c.hasIterationStopped = rightmostToIncrease == -1
	if !c.hasIterationStopped {
		c.currentIndices[rightmostToIncrease] += 1
		for i := rightmostToIncrease + 1; i < len(c.currentIndices); i++ {
			c.currentIndices[i] = c.currentIndices[i-1] + 1
		}
	}

	return c.currentIndices
}

func (c *CombinationExceptFirstElem) GetCombinations() [][]int {
	c.nextIndices()
	c.ChosenElements = append(c.ChosenElements, make([]int, c.M))
	for i := 0; i < c.M; i++ {
		c.ChosenElements[len(c.ChosenElements)-1][i] = c.Elements[c.currentIndices[i]]
	}

	c.nextIndices()
	for !c.hasIterationStopped {
		c.ChosenElements = append(c.ChosenElements, make([]int, c.M))
		for i := 0; i < c.M; i++ {
			c.ChosenElements[len(c.ChosenElements)-1][i] = c.Elements[c.currentIndices[i]]
		}

		c.nextIndices()
	}

	return c.ChosenElements
}
