package breathfirstsearch

type Person struct {
	Name     string
	Friends  []*Person
	Visited  bool
}

func BreadthFirstSearch(root *Person, name string) *Person {
	if root == nil {
		return nil
	}

	queue := []*Person{root}
	visited := make(map[*Person]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // Dequeue 

		if current.Name == name {
			return current
		}
		visited[current] = true

		for _, friend := range current.Friends {
			if !visited[friend] {
				queue = append(queue, friend) //Enqueue
				visited[friend] = true
			}
		}
	}

	return nil
}


