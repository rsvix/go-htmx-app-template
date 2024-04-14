package structs

import "time"

type User struct {
	ID                            uint
	Email                         string
	Firstname                     string
	Lastname                      string
	Password                      string
	Activationtoken               string
	Activationtokenexpiration     time.Time
	Passwordchangetoken           string
	Passwordchangetokenexpiration time.Time
	Registerip                    string
	Enabled                       int
}
