package main

import (
	"encoding/hex"
	"image/color"
	"strconv"
)

type outputFunction (func([]color.Color) string)

type Terminal struct {
	friendlyName string
	flagName     string
	output       outputFunction
}

// Terminals are defined here
var terminals = []Terminal{
	{
		friendlyName: "XFCE4Terminal",
		flagName:     "xfce",
		output:       printXfce,
	},
	{
		friendlyName: "LilyTerm",
		flagName:     "lilyterm",
		output:       printLilyTerm,
	},
	{
		friendlyName: "Terminator",
		flagName:     "terminator",
		output:       printTerminator,
	},
	{
		friendlyName: "ROXTerm",
		flagName:     "roxterm",
		output:       printRoxTerm,
	},
	{
		friendlyName: "rxvt/xterm/aterm",
		flagName:     "xterm",
		output:       printXterm,
	},
	{
		friendlyName: "Konsole",
		flagName:     "konsole",
		output:       printKonsole,
	},
	{
		friendlyName: "iTerm2",
		flagName:     "iterm2",
		output:       printITerm2,
	},
	{
		friendlyName: "urxvt",
		flagName:     "urxvt",
		output:       printURxvt,
	},
}

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

func printITerm2(colors []color.Color) string {
	output := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	output += "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n"
	output += "<plist version=\"1.0\">\n"
	output += "<dict>\n"
	for i, c := range colors {
		cc := c.(color.NRGBA)
		output += "\t<key>Ansi "
		output += strconv.Itoa(i)
		output += " Color</key>\n"
		output += "\t<dict>\n"
		output += "\t\t<key>Blue Component</key>\n"
		output += "\t\t<real>"
		output += strconv.FormatFloat(float64(cc.B)/255, 'f', 17, 64)
		output += "</real>\n"
		output += "\t\t<key>Green Component</key>\n"
		output += "\t\t<real>"
		output += strconv.FormatFloat(float64(cc.G)/255, 'f', 17, 64)
		output += "</real>\n"
		output += "\t\t<key>Red Component</key>\n"
		output += "\t\t<real>"
		output += strconv.FormatFloat(float64(cc.R)/255, 'f', 17, 64)
		output += "</real>\n"
		output += "\t</dict>\n"
	}
	output += "</dict>\n"
	output += "</plist>\n"
	return output
}

func printURxvt(colors []color.Color) string {
	output := ""
	for i, c := range colors {
		cc := c.(color.NRGBA)
		bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
		output += "URxvt*color"
		output += strconv.Itoa(i)
		output += ": "
		output += "#"
		output += hex.EncodeToString(bytes)
		output += "\n"
	}
	return output
}
