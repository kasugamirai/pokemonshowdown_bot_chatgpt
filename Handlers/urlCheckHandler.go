package Handlers

import (
	"strings"
	"xy.com/pokemonshowdownbot/models"
)

func urlCheckHanler(str string) bool {
	split := strings.Split(str, ",")
	for _, s := range split {
		if !models.URLCheck(s) {
			return false
		}
	}
	return true
}
