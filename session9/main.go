package main

import "log"
import "fmt"
import "errors"

var(
	scores = map[string]int{
		"John": 8,
		"Nick": 5,
		"Bubu": 3,
	}
	NotEnoughScore = errors.New("score is not enough")
)

// interface error {
// 	Error() string 
// }

type FindError struct {
	Err error
	Msg string
}


func (fe FindError) Error() string {
	return fe.Msg + " :==: " + fe.Err.Error() 
}

func (fe FindError) Unwrap() error {
	return fe.Err
}

func FindScore(n string) (error, int) {
	score, ok := scores[n]
	if !ok {
		return fmt.Errorf("%s does not have a score", n), 0
	}
	return nil, score
}
func Find(n string) (int, error) {
	err, s := FindScore(n)
	if err!= nil {
		return 0, FindError{
			Msg: "Not Found",
			Err: err,
		}
	}
	return s, NotEnoughScore
}



func main() {
	name := "Fred"
	s, err := Find(name)
	if err != nil {
		log.Println(err.Error())
		log.Fatal(errors.Unwrap(err))
	}
	fmt.Println(name, "'s score is ", s)
}