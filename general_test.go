package main

import "testing"

func BenchmarkAddDepartment(b *testing.B) {

	for n := 0; n < b.N; n++ {
		addDepartment("test")
	}
}

func BenchmarkAddPerson(b *testing.B) {
	createData()

	for n := 0; n < b.N; n++ {
		addPerson("test", "foo")
	}

	deleteData()
}

func BenchmarkAddToMatches(b *testing.B) {
	createData()

	for n := 0; n < b.N; n++ {
		addToMatches(&Person{
			Name:       "hullo",
			Department: departments[0],
		})
	}

	deleteData()
}

func BenchmarkPersistTeams(b *testing.B) {
	teams := createData()

	for n := 0; n < b.N; n++ {
		persistTeams(teams)
	}

	deleteData()
}

func BenchmarkGetTeams(b *testing.B) {
	createData()

	for n := 0; n < b.N; n++ {
		getTeams(3)
	}

	deleteData()
}

func createData() []Team {
	addDepartment("foo")
	addDepartment("bar")
	var teams = []Team{
		{
			Members: []*Person{
				{
					Name:       "hello",
					Department: departments[0],
				},
				{
					Name:       "goodbye",
					Department: departments[0],
				},
				{
					Name:       "Hiagain",
					Department: departments[1],
				},
			},
			Score: 0,
		},
		{
			Members: []*Person{
				{
					Name:       "one",
					Department: departments[1],
				},
				{
					Name:       "two",
					Department: departments[1],
				},
				{
					Name:       "three",
					Department: departments[0],
				},
			},
			Score: 0,
		},
	}

	addPerson("hello", "foo")
	addPerson("goodbye", "foo")
	addPerson("three", "foo")
	addPerson("Hiagain", "bar")
	addPerson("one", "bar")
	addPerson("two", "bar")

	return teams
}

func deleteData() {
	for _, p := range people {
		delPerson(p.Name, true)
	}
	for _, d := range departments {
		delDepartment(d.Name)
	}

}
