package react

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.

type reactor struct {
}

type inputCell struct {
	val      int
	children []*computeCell
}

type computeCell struct {
	val        int
	numParents int
	updater1   func(int) int
	updater2   func(int, int) int
	parents    []Cell
	children   []*computeCell
	callBacks  map[int]func(int)
	nextID     int
}

type canceler struct {
	index          int
	callBackHolder *computeCell
}

func (c *canceler) Cancel() {
	delete(c.callBackHolder.callBacks, c.index)
}

func (c *inputCell) Value() int {
	return c.val
}

func (c *computeCell) Value() int {
	return c.val
}

func (c *inputCell) SetValue(value int) {
	if c.val == value {
		return
	}
	c.val = value

	visited := map[*computeCell]bool{}
	queue := append([]*computeCell{}, c.children...)

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		if visited[cell] {
			continue
		}
		visited[cell] = true

		changed := cell.recompute()
		if changed {
			queue = append(queue, cell.children...)
		}
	}
}

func (c *computeCell) recompute() bool {
	var newVal int
	if c.numParents == 1 {
		newVal = c.updater1(c.parents[0].Value())
	} else {
		newVal = c.updater2(c.parents[0].Value(), c.parents[1].Value())
	}

	if newVal != c.val {
		c.val = newVal
		for _, cb := range c.callBacks {
			cb(newVal)
		}
		return true
	}
	return false
}
func (c *computeCell) AddCallback(callback func(int)) Canceler {
	c.callBacks[c.nextID] = callback
	c.nextID += 1
	return &canceler{index: c.nextID - 1, callBackHolder: c}
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	return &inputCell{val: initial, children: []*computeCell{}}
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	newComputeCell := &computeCell{val: compute(dep.Value()), nextID: 1, numParents: 1, updater1: compute, children: []*computeCell{}}
	if inputCell, ok := dep.(*inputCell); ok {
		inputCell.children = append(inputCell.children, newComputeCell)
	} else if computeCell, ok := dep.(*computeCell); ok {
		computeCell.children = append(computeCell.children, newComputeCell)
	}
	newComputeCell.parents = []Cell{dep}
	newComputeCell.callBacks = map[int]func(int){}
	return newComputeCell
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	newComputeCell := &computeCell{val: compute(dep1.Value(), dep2.Value()), nextID: 1, numParents: 2, updater2: compute, children: []*computeCell{}}
	if inputCell, ok := dep1.(*inputCell); ok {
		inputCell.children = append(inputCell.children, newComputeCell)
	} else if computeCell, ok := dep1.(*computeCell); ok {
		computeCell.children = append(computeCell.children, newComputeCell)
	}

	if inputCell, ok := dep2.(*inputCell); ok {
		inputCell.children = append(inputCell.children, newComputeCell)
	} else if computeCell, ok := dep2.(*computeCell); ok {
		computeCell.children = append(computeCell.children, newComputeCell)
	}
	newComputeCell.parents = []Cell{dep1, dep2}
	newComputeCell.callBacks = map[int]func(int){}
	return newComputeCell
}
