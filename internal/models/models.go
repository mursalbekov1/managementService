package models

import "time"

type User struct {
	id    int       `yaml:"id"`
	name  string    `yaml:"name"`
	email string    `yaml:"email"`
	date  time.Time `yaml:"date"`
	role  string    `yaml:"role"`
}

type Task struct {
	id          int       `yaml:"id"`
	name        string    `yaml:"name"`
	description string    `yaml:"description"`
	priority    string    `yaml:"priority"`
	state       string    `yaml:"state"`
	responsible string    `yaml:"responsible"`
	project     int       `yaml:"project"`
	date        time.Time `yaml:"date"`
	finishDate  time.Time `yaml:"finish_date"`
}

type Project struct {
	id          int       `yaml:"id"`
	name        string    `yaml:"name"`
	description string    `yaml:"description"`
	date        time.Time `yaml:"date"`
	finishDate  time.Time `yaml:"finish_date"`
	manager     int       `yaml:"manager"`
}
