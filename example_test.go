package ptr_test

import (
	"fmt"
	"strconv"

	"github.com/hightech-ninja/ptr"
)

func ExampleTo() {
	// take pointers to constant values without need to create a temporary variable
	ptr1 := ptr.To(42)
	fmt.Println(*ptr1)
	ptr2 := ptr.To("string")
	fmt.Println(*ptr2)

	// take pointer to map elements
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	ptr3 := ptr.To(m["three"])
	fmt.Println(*ptr3)

	// take pointer to the function returned value
	f := func() string { return "result" }
	ptr4 := ptr.To(f())
	fmt.Println(*ptr4)

	// value and pointers are independent and not affect each other
	original := "dragon"
	pointer := ptr.To(original)
	*pointer = "age"
	fmt.Println(original, *pointer)

	// Output:
	// 42
	// string
	// 3
	// result
	// dragon age
}

func ExampleDeref() {
	// get optional fields
	type user struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}
	user1 := user{
		FirstName:  "John",
		LastName:   "Doe",
		MiddleName: nil,
	}

	fullName := fmt.Sprintf(
		"%s %s %s",
		user1.FirstName,
		ptr.Deref(user1.MiddleName),
		user1.LastName,
	)
	fmt.Println("User:", fullName)

	// Output:
	// User: John  Doe
}

func ExampleDerefToDefault() {
	// get optional fields, but when zero value for nil is not suitable
	var clientID1 *string
	value1 := ptr.DerefToDefault(clientID1, "unknown")
	fmt.Println(value1)

	clientID2 := ptr.To("your-best-client@gmail.com")
	value2 := ptr.DerefToDefault(clientID2, "unknown")
	fmt.Println(value2)

	// Output:
	// unknown
	// your-best-client@gmail.com
}

func ExampleReset() {
	// Reusing Memory in Loops
	number := ptr.To(5)
	for i := 0; i < 3; i++ {
		fmt.Println(*number)
		ptr.Reset(number)
	}

	// Clearing User Input in GUI Applications
	userInput := ptr.To("User input data")
	// Assume user triggers a reset action
	ptr.Reset(userInput)
	fmt.Println("User input after reset:", *userInput)

	// Output:
	// 5
	// 0
	// 0
	// User input after reset:
}

func ExampleResetTo() {
	// Updating Configurations Dynamically
	maxUsers := ptr.To(100)
	fmt.Println("Max Users:", *maxUsers)
	ptr.ResetTo(maxUsers, 150)
	fmt.Println("Updated Max Users:", *maxUsers)

	// Modifying State in Stateful Applications
	gameState := ptr.To("paused")
	fmt.Println("Game State:", *gameState)
	ptr.ResetTo(gameState, "running")
	fmt.Println("Game State after update:", *gameState)

	// Output:
	// Max Users: 100
	// Updated Max Users: 150
	// Game State: paused
	// Game State after update: running
}

func ExampleCompare() {
	// Compare data structures by pointers
	type node struct {
		left  *node
		right *node
		Value int
	}
	node1 := &node{
		Value: 5,
	}
	node2 := &node{
		left:  nil,
		right: nil,
		Value: 5,
	}
	fmt.Println("Nodes equal:", ptr.Compare(node1, node2))

	// Equality Checks in Unit Testing
	expected := ptr.To(42)
	actual := ptr.To(42)
	fmt.Println("Test result:", ptr.Compare(expected, actual))

	// Output:
	// Nodes equal: true
	// Test result: true
}

func ExampleMap() {
	age := ptr.To(33)
	str  := ptr.Map(age, strconv.Itoa)
	fmt.Printf("%q",*str)
	
	// Output:
	// "33"
}