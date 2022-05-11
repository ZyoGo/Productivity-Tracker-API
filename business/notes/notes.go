package notes

import "time"

type Notes struct {
	Id         int32
	Status     string
	Content    string
	Tags       []string
	Created_At time.Time
	Updated_At time.Time
	Deleted    bool
}
