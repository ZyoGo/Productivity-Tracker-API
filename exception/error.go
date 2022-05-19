package exception

import "errors"

var (
	//ErrHasBeenModified Error when update data that has been modified
	ErrDataHasBeenModified = errors.New("Data has been modified")

	//ErrInvalidSpec Error when data given is not valid on update or insert
	ErrInvalidSpec = errors.New("Given spec is not valid")

	//ErrZeroAffected Data not found
	ErrNotFound = errors.New("No record affected")

	ErrInternalServer = errors.New("Internal server error")
)