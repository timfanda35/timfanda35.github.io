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
)

const (
	FontPath           = "NotoSansTC-Black.ttf"
	OutputDir          = "static/images"
	OgImageWidth       = 1200
	OgImageHeight      = 600
	OgTitleFontSize    = 48
	OgSubTitleFontSize = 32
	OgPaddingLeft      = 72.0
	OgPaddingTop       = 90.0

	SubTitle      = "Bear Su's blog"
	IconImagePath = "icon.png"
)

// Load font with goki freetype
func loadFont(path string) *truetype.Font {
	var fontBytes []byte
	fontBytes, err := os.ReadFile(FontPath)
	if err != nil {
		panic(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	return font
}

func extractTitle(postContent string) string {
	re := regexp.MustCompile(`\ntitle:(.+)\n`)
	matches := re.FindStringSubmatch(string(postContent))

	if len(matches) < 1 {
		panic("Can not find title")
	}

	return strings.Trim(matches[1], " \"")
}

func extractDate(postContent string) string {
	re := regexp.MustCompile(`\ndate: (\d{4}-\d{2}-\d{2})T`)
	matches := re.FindStringSubmatch(string(postContent))

	if len(matches) < 1 {
		panic("Can not find date")
	}

	return matches[1]
}

func imageOutputPath(postFilePath string) string {
	re := regexp.MustCompile(`posts(/\d{4}-\d{2}-\d{2})-(.+).md`)
	matches := re.FindStringSubmatch(postFilePath)
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
	postFilePath := os.Args[1]
	postFile, err := os.ReadFile(postFilePath)
	if err != nil {
		panic(err)
	}

	postContent := string(postFile)

	// Get Title
	postTitle := extractTitle(postContent)
	log.Println("postTitle:", postTitle)

	// Get Date
	postDate := extractDate(postContent)
	log.Println("postDate:", postDate)

	// Create context
	dc := gg.NewContext(OgImageWidth, OgImageHeight)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Load Font
	font := loadFont("NotoSansTC-Black.ttf")

	axisX := OgPaddingLeft
	axisY := OgPaddingTop

	// Draw Title
	titleFace := truetype.NewFace(font, &truetype.Options{Size: OgTitleFontSize})
	dc.SetFontFace(titleFace)
	dc.SetRGB(0, 0, 0)
	dc.DrawStringWrapped(postTitle, axisX, axisY, 0, 0, 1100, 1, gg.AlignLeft)

	axisY += OgTitleFontSize * 4

	// Draw Icon
	iconImage, err := gg.LoadImage(IconImagePath)
	if err != nil {
		log.Println(err)
	} else {
		w := iconImage.Bounds().Size().X
		iconImageAxisX := OgImageWidth - OgPaddingLeft - w
		dc.DrawImage(iconImage, int(iconImageAxisX), int(axisY))
	}

	// Draw SubTitle
	subTitleFace := truetype.NewFace(font, &truetype.Options{Size: OgSubTitleFontSize})
	dc.SetFontFace(subTitleFace)
	dc.SetRGB(150, 150, 150)
	dc.DrawStringWrapped(SubTitle, axisX, axisY, 0, 0, 1100, 1, gg.AlignLeft)

	axisY += OgTitleFontSize

	// Draw Date
	dc.DrawStringWrapped(postDate, axisX, axisY, 0, 0, 1100, 1, gg.AlignLeft)

	// Save
	ogImage := imageOutputPath(postFilePath)
	dc.SavePNG(ogImage)

	log.Println("OG Image Save to:", ogImage)
}
