package main

import (
	"fmt"
	"math/rand"
	"time"
)

//List is the thing that the jsons go int
type List struct {
	column1, column2, column3 []string
}

var shakespearean = List{
	column1: []string{"artless", "bawdy", "beslubbering", "bootless", "churlish", "cockered", "clouted", "craven", "currish", "darkish", "dissembling", "droning", "errant", "fawning", "fobbing", "froward", "frothy", "gleeking", "goatish", "gorbellied", "impertinent", "infectious", "jarring", "loggerheaded", "lumpish", "mammering", "mangled", "mewling", "paunchy", "pribbling", "puking", "puny", "qualling", "rank", "reeky", "roguish", "ruttish", "saucy", "spleeny", "spongy", "surly", "tottering", "unmuzzled", "vain", "venomed", "villainous", "warped", "wayward", "weedy", "yeasty"},
	column2: []string{"base-court", "bat-fowling", "beef-witted", "beetle-headed", "boil-brained", "clapper-clawed", "clay-brained", "common-kissing", "crook-plated", "dismal-dreaming", "dizzy-eyed", "doghearted", "dread-boiled", "earth-vexing", "elf-skinned", "fat-kidneyed", "fen-sucked", "flap-mouthed", "fly-bitten", "folly-fallen", "fool-born", "full-gorged", "guts-gripping", "half-faced", "hasty-witted", "hedge-born", "hell-hated", "idle-headed", "ill-breeding", "ill-nurtured", "knotty-pated", "milk-livered", "motley-minded", "onion-eyed", "plume-plucked", "pottle-deep", "pox-marked", "reeling-ripe", "rough-hewn", "rude-growing", "rump-fed", "shard-borne", "sheep-biting", "spur-galled", "swag-bellied", "tardy-gaited", "tickle-brained", "toad-spotted", "urchin-snouted", "weather-bitten"},
	column3: []string{"artless", "bawdy", "beslubbering", "bootless", "churlish", "cockered", "clouted", "craven", "currish", "darkish", "dissembling", "droning", "errant", "fawning", "fobbing", "froward", "frothy", "gleeking", "goatish", "gorbellied", "impertinent", "infectious", "jarring", "loggerheaded", "lumpish", "mammering", "mangled", "mewling", "paunchy", "pribbling", "puking", "puny", "qualling", "rank", "reeky", "roguish", "ruttish", "saucy", "spleeny", "spongy", "surly", "tottering", "unmuzzled", "vain", "venomed", "villainous", "warped", "wayward", "weedy", "yeasty"},
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	column1 := shakespearean.column1
	column2 := shakespearean.column2
	column3 := shakespearean.column3

	first := r.Intn(len(column1))
	second := r.Intn(len(column2))
	third := r.Intn(len(column3))

	fmt.Printf("thou %v %v %v!\n", column1[first], column2[second], column3[third])
}
