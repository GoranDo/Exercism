package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

//Robot is just a name (for now)
type Robot string

//namesInUse keeps track of which names are in use
var namesInUse = make(map[Robot]bool)

/*Name returns a robots name*/
func (r *Robot) Name() (string, error) {
	rand.Seed(time.Now().UnixNano())
	for inUse, ok := string(*r) == "", true; ok && inUse; inUse, ok = namesInUse[*r] {
		*r = Robot(letter() + letter() + number())
	}
	namesInUse[*r] = true
	return string(*r), nil
}

/*Reset give a robot a new name*/
func (r *Robot) Reset() {
	*r = Robot("")
}

/*letter generates a random upper case letter*/
func letter() string {
	return string(rune(rand.Intn('Z'-'A') + 'A'))
}

/*number generates a random three digit number*/
func number() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}
