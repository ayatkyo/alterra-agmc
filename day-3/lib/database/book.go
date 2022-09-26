package database

import "github.com/ayatkyo/alterra-agmc/day-3/models"

var BookDB []models.Book = []models.Book{
	{
		ID:     1,
		Title:  "Software Architecture with C++",
		Author: "Adrian Ostrowsmodels",
		Year:   2021,
	},
	{
		ID:     2,
		Title:  "Modern C++ Programming Cookbook - Second Edition",
		Author: "Marius Bancila",
		Year:   2020,
	},
	{
		ID:     3,
		Title:  "Flutter Apprentice - Third Edition",
		Author: "Vincent Ngo, Kevin D Moore and Michael Katz",
		Year:   2022,
	},
}

var BookDBID = len(BookDB)
