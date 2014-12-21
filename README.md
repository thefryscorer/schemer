Terminal Colorscheme Generator
==============================

## Screenshot 
![Screenshot](http://i.imgur.com/TSLluID.png)

## Installation 

### Short version

> go get github.com/thefryscorer/schemer

### Long Version

#### Installing and configuring Go
To build this program, you will need to have Go installed and properly configured. After installing the Go package, you will need to configure a GOPATH. This is a directory in which Go will keep its programs and source files. I recommend making the GOPATH directory in your home folder. If your GOPATH is in your root directory a kitten will die. 

> mkdir ~/Go

You will also need to set the GOPATH variable so that Go knows where to put things. You can do this by running:

> export $GOPATH=/home/<USER>/Go

Where <USER> is your username.

NOTE: You don't need to (and shouldn't) set the $GOROOT variable. This is handled for you and you shouldn't mess with it.

#### Installing SDL1.2
This program also makes use of SDL1.2 for the color preview window (which you can use by adding the "-d" flag in schemer). As such, SDL1.2 will need to be installed on your system for you to build and run schemer.

*To install SDL1.2 in ArchLinux:*

> sudo pacman -S sdl sdl_image sdl_ttf sdl_mixer

*To install SDL1.2 on another system:*

Learn how to use Google and your package manager. Both are very useful. 

#### Installing schemer
You should now be able to install schemer using the command:

> go get github.com/thefryscorer/schemer

And it will be built in your GOPATH directory, in a subdirectory named 'bin'. To run it, you can either add /home/<USER>/Go/bin to your system path and run it as you would any other command. Or cd into the bin directory and run it with:

> ./schemer

## Usage 

> schemer -term="xfce" Image.png

Then copy the generated config lines into your terminal config file.

## Features 

- Outputs configuration in several different formats for different terminals.
- Configurable color difference threshold
- Configurable minimum and maximum brightness value
- Can preview colorscheme in SDL window

## Supported output formats

- Colours in just plain text (default)
- Konsole
- xterm/rxvt/aterm
- urxvt
- iTerm2
- XFCE Terminal
- Roxterm
- LilyTerm
- Terminator
- Chrome Shell
