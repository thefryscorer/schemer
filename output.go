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
			output += "\n"
		}
	}
	return output
}

func printXterm(colors []color.Color) string {
	output := ""
	output += "! Terminal colors"
	output += "\n"
	for i, c := range colors {
		cc := c.(color.NRGBA)
		bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
		output += "*color"
		output += strconv.Itoa(i)
		output += ": #"
		output += hex.EncodeToString(bytes)
		output += "\n"
	}

	return output
}

func printKonsole(colors []color.Color) string {
	output := ""
	for i, c := range colors {
		cc := c.(color.NRGBA)
		output += "[Color"
		if i > 7 {
			output += strconv.Itoa(i - 8)
			output += "Intense"
		} else {
			output += strconv.Itoa(i)
		}
		output += "]\n"
		output += "Color="
		output += strconv.Itoa(int(cc.R)) + ","
		output += strconv.Itoa(int(cc.G)) + ","
		output += strconv.Itoa(int(cc.B)) + "\n"
		output += "Transparency=false\n\n"
	}

	return output
}

func printRoxTerm(colors []color.Color) string {
	output := "[roxterm colour scheme]\n"
	output += "pallete_size=16\n"

	for i, c := range colors {
		cc := c.(color.NRGBA)
		bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
		output += "color"
		output += strconv.Itoa(i)
		output += " = "
		output += "#"
		output += hex.EncodeToString(bytes)
		output += "\n"
	}

	return output
}
