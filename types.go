package main

type Person struct {
	Name       string
	Department string
	Score      int
}

type Department struct {
	Name         string
	NumberPeople int
}

type Team struct {
	Members []Person
	Score   int
}

type Match struct {
	Match [2]Person
	Score int
}

type Departments []Department

func (slice Departments) Len() int {
	return len(slice)
}

func (slice Departments) Less(i, j int) bool {
	return slice[i].NumberPeople < slice[j].NumberPeople
}

func (slice Departments) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type People []Person

func (slice People) Len() int {
	return len(slice)
}

func (slice People) Less(i, j int) bool {
	return slice[i].Score < slice[j].Score
}

func (slice People) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
