package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const levels = "▁▂▃▄▅▆▇█"

var levelSize float64

func init() {
	levelSize = 100.0 / float64(len([]rune(levels)))
}

func get(ps string) int {
	data, err := ioutil.ReadFile("/sys/class/power_supply/" + ps)
	if err != nil {
		log.Fatal(err)
	}
	value, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func levelIndex(capacity int) int {
	if capacity >= 100 {
		return len([]rune(levels)) - 1
	}
	if float64(capacity) < levelSize {
		return 0
	}
	return int(((float64(capacity) + levelSize) / levelSize) - 1)
}

// LevelGlyph gets the power supply level glyph that represents capacity
func LevelGlyph(capacity int) string {
	return string([]rune(levels)[levelIndex(capacity)])
}

func main() {
	capacity := get("BAT0/capacity")
	fmt.Printf("%s▏%d%%\n", LevelGlyph(capacity), capacity)
}
