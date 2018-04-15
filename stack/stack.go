package stack

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Popp removes the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

// func main() {
// 	stack := new(Stack)

// 	stack.Push("Things")
// 	stack.Push("and")
// 	stack.Push("Stuff")

// 	fmt.Println(stack)

// 	for stack.Len() > 0 {
// 		// We have to do a type assertion because we get back a variable of type
// 		// interface{} while the underlying type is a string.
// 		fmt.Printf("%s ", stack.Pop().(string))
// 	}
// 	fmt.Println()
// }
