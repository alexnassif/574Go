
package tree

var SIZE = 100000

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New() *Tree {
	var t *Tree = &Tree{
		Left:  nil,
		Value: 0,
		Right: nil,
	}
	for v := 1; v < SIZE; v++ {
		t = insert(t, v)
	}
	return t
}

func insert(t *Tree, v int) *Tree {

	var current = t
	var previous *Tree

	for current != nil {
		previous = current

		if v < current.Value{
			current = current.Left
		}else{
			current = current.Right
		}
	}

	current = &Tree{
		Left:  nil,
		Value: v,
		Right: nil,
	}

	if v > previous.Value{
		previous.Right = current
	}else{
		previous.Left = current
	}

	return t
}


