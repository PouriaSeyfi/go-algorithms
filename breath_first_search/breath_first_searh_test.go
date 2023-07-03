package breathfirstsearch

import "testing"

func TestBreadthFirstSearch(t *testing.T) {
	// Create a social network graph for testing
	//   Alice -> Bob -> Charlie -> David
	//          |
	//          -> John

	alice := &Person{Name: "Alice"}
	bob := &Person{Name: "Bob"}
	charlie := &Person{Name: "Charlie"}
	david := &Person{Name: "David"}
	john := &Person{Name: "John"}

	alice.Friends = []*Person{bob, john}
	bob.Friends = []*Person{charlie}
	charlie.Friends = []*Person{david}


	// Search for a specific person in the social network
	person := BreadthFirstSearch(alice, "Charlie")
	if person == nil {
		t.Errorf("Expected to find person Charlie, but got nil")
	} else if person.Name != "Charlie" {
		t.Errorf("Expected person Charlie, but got %s", person.Name)
	}
}
