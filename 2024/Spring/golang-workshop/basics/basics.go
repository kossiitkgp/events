package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to Go workshop!")

	// Variable Declaration
	var a bool
	var b = 5
	var c, d = 10, 20
	var (
		e = 30
		f int
	)
	var s string

	fmt.Println(a, f, s)    // Variables declared without an explicit value are initialized to their "zero value"
	fmt.Println(b, c, d, e) // Unused declared variables raise error

	x, y := "Hello", "KOSS" // Short Variable Declarations (allowed only inside a function)
	fmt.Println(x, y)

	// For Loop, If-Else Statements
	num, maxguesses := 25, 5
	for i := 0; i < maxguesses; i++ {
		fmt.Printf("You have %d guess(es)\n", maxguesses-i) // similar to printf in C
		fmt.Print("Enter number: ")                         // newline is not appended by default
		var tmp int
		fmt.Scanln(&tmp)

		if tmp == num {
			fmt.Println("Correct!")
			break
		} else {
			fmt.Println("oops!")
		}
	}

	num = 100
	fmt.Println("You have Unlimited Guesses")
	for { // "while" loop
		var tmp int
		fmt.Scanln(&tmp)

		if tmp == num {
			fmt.Println("Correct!")
			break
		} else if tmp > num {
			fmt.Println("Lower..")
		} else {
			fmt.Println("Higher..")
		}
	}

	// Arrays
	var arr1 [5]int // Initialized to 0
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, arr2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	// Slices
	var sli []string // nil Slice
	sli = []string{"a", "b", "c"}
	fmt.Println(sli)
	sli = append(sli, "d")
	sli = append(sli, "a", "c")
	fmt.Println(sli)

	part := arr2[1:4] // Slicing
	part[0] = 7       // Slicing does not copy the slice’s data. It creates a new slice value that points to the original array
	fmt.Println(part, arr2)

	// Maps
	var mp map[string]int // nil map; has no keys, nor keys can be added

	// make - allocates memory and initializes the underlying data structures
	mp = make(map[string]int) // map[key]value
	sl := make([]string, len(sli))
	copy(sl, sli)

	for _, str := range sl {
		mp[str]++
	}
	fmt.Println(sl, mp)

	val, isPresent := mp["d"] //  isPresent is an OPTIONAL return value
	fmt.Println(val, isPresent)
	delete(mp, "d")          // Delete a key
	val, isPresent = mp["d"] // If the key doesn't exist, zero value is returned
	fmt.Println(val, isPresent)

	var mp2 = map[int]string{
		1: "ONE",
		2: "TWO",
		3: "THREE",
	}
	for k, v := range mp2 {
		fmt.Println(k, "->", v)
	}

	// Functions
	nums := []int{10, 2, 3, 4, 5}
	fmt.Println(max(nums[0], nums[1]))
	fmt.Println(maxmin(nums)) // Multiple return values
	changeSlice(nums)         // Slices are passed by reference
	fmt.Println(nums)

	// Structs
	p1 := Point{"pt1", 12.23, 9.7} // Create an instance of struct
	p2 := Point{X: 5.6, Y: 10.0}   // Omitted fields will be zero-valued
	fmt.Println(p1.X, p2.Y)        // Access Values
	fmt.Println(p1.dist(p2))       // Struct Methods

	// Pointers
	num1, num2 := 100, 250
	fmt.Println(num1, num2)
	swap(&num1, &num2) // Pass by Reference
	fmt.Println(num1, num2)

	p3 := Point{"pt", 67.3, 0.0}
	p3ptr := &p3
	fmt.Println(p3ptr.label) // Can also use dots with struct pointers - the pointers are automatically dereferenced.
}

// Functions
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Multiple return values
func maxmin(arr []int) (int, int) {
	mx, mn := -1000, 1000

	for _, num := range arr {
		mx = max(mx, num)
		mn = min(mn, num)
	}

	return mx, mn
}

// Slice is passed by reference
func changeSlice(arr []int) {
	arr[0] = -1
}

// Structs
type Point struct {
	// Fields
	label string
	X     float64
	Y     float64
}

// Struct Methods
func (p1 Point) dist(p2 Point) float64 {
	return math.Sqrt(sq(p1.X-p2.X) + sq(p1.Y-p2.Y))
}

func sq(n float64) float64 {
	return n * n
}

// Pass by Reference
func swap(ptr1 *int, ptr2 *int) {
	temp := *ptr1
	*ptr1 = *ptr2
	*ptr2 = temp
}
