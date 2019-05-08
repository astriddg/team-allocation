package main

import (
	"testing"
)

func BenchmarkAddDepartment(b *testing.B) {

	for n := 0; n < b.N; n++ {
		addDepartment("test")
	}
}

func BenchmarkAddPerson(b *testing.B) {

	for n := 0; n < b.N; n++ {
		createData()
		addPerson("test", "foo")
		deleteData()
	}

}

func BenchmarkAddToMatches(b *testing.B) {

	for n := 0; n < b.N; n++ {
		createData()
		addToMatches(&Person{
			Name:       "hullo",
			Department: "foo",
		})
		deleteData()
	}

}

func BenchmarkPersistTeams(b *testing.B) {

	for n := 0; n < b.N; n++ {
		teams := createData()
		persistTeams(teams)
		deleteData()
	}

}

func BenchmarkGetTeams(b *testing.B) {

	for n := 0; n < b.N; n++ {
		createData()
		getTeams(3, []string{"hello"})
		deleteData()
	}

}

func createData() []Team {
	addDepartment("foo")
	addDepartment("bar")
	var teams = []Team{
		{
			Members: []Person{
				{
					Name:       "hello",
					Department: "foo",
				},
				{
					Name:       "goodbye",
					Department: "foo",
				},
				{
					Name:       "Hiagain",
					Department: "bar",
				},
			},
			Score: 0,
		},
		{
			Members: []Person{
				{
					Name:       "one",
					Department: "bar",
				},
				{
					Name:       "two",
					Department: "bar",
				},
				{
					Name:       "three",
					Department: "foo",
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
		delPerson(p.Name)
	}
	for _, d := range departments {
		delDepartment(d.Name)
	}

}
