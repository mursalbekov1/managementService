package models

import "time"

type User struct {
	Id    int       `yaml:"id"`
	Name  string    `yaml:"name"`
	Email string    `yaml:"email"`
	Date  time.Time `yaml:"date"`
	Role  string    `yaml:"role"`
}

type Task struct {
	Id          int       `yaml:"id"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Priority    string    `yaml:"priority"`
	State       string    `yaml:"state"`
	Responsible string    `yaml:"responsible"`
	Project     int       `yaml:"project"`
	Date        time.Time `yaml:"date"`
	FinishDate  time.Time `yaml:"finish_date"`
}

type Project struct {
	Id          int       `yaml:"id"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Date        time.Time `yaml:"date"`
	FinishDate  time.Time `yaml:"finish_date"`
	Manager     int       `yaml:"manager"`
}
