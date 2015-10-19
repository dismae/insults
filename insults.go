//Package insults is for generating insults.
package insults

import (
	"fmt"
	"math/rand"
	"time"
)

//InsultList columns should be in  order that words would appear in a sentence
//you should be the word that would appear at the beginning of the sentence,
//ie Thou for the shakespeaerean list.
type InsultList struct {
	column1, column2, column3 []string
	you                       string
}

//RandInsult takes an InsultList and generates an insult.
func RandInsult(l InsultList) string {
	return fmt.Sprintf("%v %v %v %v!", l.you, l.column1, l.column2, l.column3)
}
func init() {
	rand.Seed(time.Now().UnixNano())
}
