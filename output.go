package main

import (
	"encoding/hex"
	"image/color"
	"strconv"
)

func printXfce(colors []color.Color) string {
	output := ""
	output += "ColorPalette="
	for _, c := range colors {
		bytes := []byte{byte(c.(color.NRGBA).R), byte(c.(color.NRGBA).R), byte(c.(color.NRGBA).G), byte(c.(color.NRGBA).G), byte(c.(color.NRGBA).B), byte(c.(color.NRGBA).B)}
		output += "#"
		output += hex.EncodeToString(bytes)
		output += ";"
	}
	output += "\n"

	return output
}

func printLilyTerm(colors []color.Color) string {
	output := ""
	for i, c := range colors {
		cc := c.(color.NRGBA)
		bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
		output += "Color"
		output += strconv.Itoa(i)
		output += " = "
		output += "#"
		output += hex.EncodeToString(bytes)
		output += "\n"
	}
	return output
}

func printTerminator(colors []color.Color) string {
	output := "palette = \""
	for i, c := range colors {
		cc := c.(color.NRGBA)
		bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
		if i < len(colors)-1 {
			output += "#"
			output += hex.EncodeToString(bytes)
			output += ":"
		} else if i == len(colors)-1 {
			output += "#"
			output += hex.EncodeToString(bytes)
		}
	}
	return output
}
