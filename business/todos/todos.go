package todos

import "time"

type Todos struct {
	Id         int32
	User_Id    int32
	Content    string
	Created_At time.Time
	Updated_At time.Time
	Deleted    bool
}
