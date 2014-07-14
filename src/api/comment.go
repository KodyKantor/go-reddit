package api

import (
	"strconv"
)

//Comment is a struct that represents
//comments on a link
type Comment struct {
	Body   string
	Author string
	Edited bool
}

func (comment Comment) String() string {
	result := ""

	result += "Body is " + comment.Body + "\n"
	result += "Author is " + comment.Author + "\n"
	result += "Edited? " + strconv.FormatBool(comment.Edited) + "\n"

	return result
}
