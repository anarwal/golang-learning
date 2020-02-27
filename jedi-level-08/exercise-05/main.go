package main

import (
	"fmt"
	"sort"
)

// Sort the []user by
//   - age
//   - last
// also sort each []string "Sayings" for each user
// Print everything out using the range clause

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// ByAge : method use to sort by age
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByLast : method use to sort by lastname
type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is not guarantee of innovation",
		},
	}
	u2 := user{
		First: "M",
		Last:  "B",
		Age:   50,
		Sayings: []string{
			"Oh James, you killed him",
			"Dear James",
		},
	}

	u3 := user{
		First: "Bob",
		Last:  "Call",
		Age:   22,
		Sayings: []string{
			"Bob, its me",
			"Dear M, its Bob",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)
	fmt.Println("------")
	sort.Sort(ByAge(users))
	fmt.Println(users)
	fmt.Println("------")
	sort.Sort(ByLast(users))
}
