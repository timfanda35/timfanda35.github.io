package main

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/fogleman/gg"

	// Fix Google Font issue
	// https://www.reddit.com/r/golang/comments/14ldjiu/comment/jq1du7t/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
	"github.com/goki/freetype"
	"github.com/goki/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	FontPath      = "NotoSansTC-Black.ttf"
	OutputDir     = "static/images"
	OgImageWidth  = 1200
	OgImageHeight = 600
)

// Load font with goki freetype
func loadFontFace(path string) font.Face {
	var fontBytes []byte
	fontBytes, err := os.ReadFile(FontPath)
	if err != nil {
		panic(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(font, &truetype.Options{
		Size: 72,
	})

	return face
}

func imageOutputPath(postFile string) string {
	re := regexp.MustCompile(`posts(/\d{4}-\d{2}-\d{2})-(.+).md`)
	matches := re.FindStringSubmatch(postFile)
	if len(matches) < 2 {
		panic("Can not parse file name")
	}

	dir := OutputDir + matches[1]

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			panic("Can not create output directory")
		}
	}

	return dir + "/" + matches[2] + ".png"
}

func main() {
	// Read file
	postFile := os.Args[1]
	postContent, err := os.ReadFile(postFile)
	if err != nil {
		panic(err)
	}
	// Get Title
	re := regexp.MustCompile(`\ntitle:(.+)\n`)
	matches := re.FindStringSubmatch(string(postContent))

	if len(matches) < 1 {
		panic("Can not find title")
	}
	postTitle := strings.Trim(matches[1], " \"")
	log.Println("postTitle:", postTitle)

	// Create context
	dc := gg.NewContext(OgImageWidth, OgImageHeight)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Load Font
	face := loadFontFace("NotoSansTC-Black.ttf")
	dc.SetFontFace(face)

	// Draw Text
	dc.SetRGB(0, 0, 0)
	dc.DrawStringWrapped(postTitle, 50, 122, 0, 0, 1100, 1, gg.AlignLeft)

	// Save
	ogImage := imageOutputPath(postFile)
	dc.SavePNG(ogImage)

	log.Println("OG Image Save to:", ogImage)
}
