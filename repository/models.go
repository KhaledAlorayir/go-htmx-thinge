package repository

import "time"

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type MuscleGroup struct {
	Id        int
	Name      string
	CreatedAt time.Time
}

type Exercise struct {
	Id            int
	Name          string
	CreatedAt     time.Time
	MuscleGroupId int
}

type UserExercise struct {
	Id           int
	CreatedAt    time.Time
	UserId       int
	ExerciseId   int
	Weight       float32
	InclineLevel int
	Note         string
	Link         string
	NumberOfSets int
}
