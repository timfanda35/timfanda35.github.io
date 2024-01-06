package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"flag"

	"github.com/fogleman/gg"

	// Fix Google Font issue
	// https://www.reddit.com/r/golang/comments/14ldjiu/comment/jq1du7t/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
	"github.com/goki/freetype"
	"github.com/goki/freetype/truetype"
)

const (
	FontFile           = "NotoSansTC-Black.ttf"
	OutputDir          = "static/images"
	OgImageWidth       = 1200
	OgImageHeight      = 600
	OgTitleFontSize    = 48
	OgSubTitleFontSize = 32
	OgPaddingLeft      = 72.0
	OgPaddingTop       = 90.0
	OgBottomLineY      = 576

	SubTitle      = "Bear Su's blog"
	IconImagePath = "icon.png"
)

// Load font with goki freetype
func loadFont(path string) *truetype.Font {
	var fontBytes []byte
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not load font file: %s", path)
		os.Exit(1)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not parse font file: %s", path)
		os.Exit(1)
	}

	return font
}

func loadPostContent(postFilePath string) string {
	postFile, err := os.ReadFile(postFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not read input file: %s", postFilePath)
		os.Exit(1)
	}

	return string(postFile)
}

func extractTitle(postContent string) string {
	re := regexp.MustCompile(`\ntitle:(.+)\n`)
	matches := re.FindStringSubmatch(string(postContent))

	if len(matches) < 1 {
		fmt.Fprintf(os.Stderr, "Can not find post title")
		os.Exit(1)
	}

	return strings.Trim(matches[1], " \"")
}

func extractDate(postContent string) string {
	re := regexp.MustCompile(`\ndate: (\d{4}-\d{2}-\d{2})T`)
	matches := re.FindStringSubmatch(string(postContent))

	if len(matches) < 1 {
		fmt.Fprintf(os.Stderr, "Can not find post date")
		os.Exit(1)
	}

	return matches[1]
}

func imageOutputPath(outputDir string, postFilePath string) string {
	re := regexp.MustCompile(`posts(/\d{4}-\d{2}-\d{2})-(.+).md`)
	matches := re.FindStringSubmatch(postFilePath)
	if len(matches) < 2 {
		fmt.Fprintf(os.Stderr, "Can not parse file name: %s", postFilePath)
		os.Exit(1)
	}

	dir := outputDir + matches[1]

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "Can not create output directory: %s", dir)
			os.Exit(1)
		}
	}

	return dir + "/" + matches[2] + ".png"
}

func executableDir() string {
	executablePath, err := os.Executable()
	if err != nil {
		return ""
	}

	return filepath.Dir(executablePath)
}

type OGRenRequest struct {
	InputPath     string
	IconImagePath string
	OutputDir     string
}

func drawTitle(dc *gg.Context, postTitle string, font *truetype.Font, x float64, y float64) {
	titleFace := truetype.NewFace(font, &truetype.Options{Size: OgTitleFontSize})
	dc.SetFontFace(titleFace)
	dc.SetRGB255(0, 0, 0)
	dc.DrawStringWrapped(postTitle, x, y, 0, 0, 1100, 1, gg.AlignLeft)
}

func drawIcon(dc *gg.Context, iconImagePath string, y float64) {
	iconImage, err := gg.LoadImage(iconImagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not load icon image: %s", iconImage)
		os.Exit(1)
	} else {
		w := iconImage.Bounds().Size().X
		iconImageAxisX := OgImageWidth - OgPaddingLeft - w
		dc.DrawImage(iconImage, int(iconImageAxisX), int(y))
	}
}

func drawSubTitle(dc *gg.Context, postTitle string, font *truetype.Font, x float64, y float64) {
	subTitleFace := truetype.NewFace(font, &truetype.Options{Size: OgSubTitleFontSize})
	dc.SetFontFace(subTitleFace)
	dc.SetRGB255(150, 150, 150)
	dc.DrawStringWrapped(SubTitle, x, y, 0, 0, 1100, 1, gg.AlignLeft)
}

func (og OGRenRequest) Do() string {
	// Load Post Content
	postContent := loadPostContent(og.InputPath)

	// Get Title
	postTitle := extractTitle(postContent)

	// Get Date
	postDate := extractDate(postContent)

	// Create context
	dc := gg.NewContext(OgImageWidth, OgImageHeight)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Load Font
	font := loadFont(executableDir() + "/" + FontFile)

	axisX := OgPaddingLeft
	axisY := OgPaddingTop

	// Draw Title
	drawTitle(dc, postTitle, font, axisX, axisY)
	axisY += OgTitleFontSize * 4

	// Draw Icon
	drawIcon(dc, og.IconImagePath, axisY)

	// Draw SubTitle
	drawSubTitle(dc, SubTitle, font, axisX, axisY)
	axisY += OgTitleFontSize

	// Draw Date
	dc.DrawStringWrapped(postDate, axisX, axisY, 0, 0, 1100, 1, gg.AlignLeft)

	// Draw Bottom Line
	dc.DrawRectangle(0, OgBottomLineY, OgImageWidth, OgImageHeight-OgBottomLineY)
	dc.SetRGB255(227, 76, 37)
	dc.Fill()

	// Save
	ogImage := imageOutputPath(og.OutputDir, og.InputPath)
	dc.SavePNG(ogImage)

	return ogImage
}

func main() {
	var dir string
	var webdir string

	flag.StringVar(&dir, "dir", "./", "The hugo project root path")
	flag.StringVar(&webdir, "webdir", "static", "The hugo website root path")
	flag.Parse()

	// Read file
	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) == 0 {
		fmt.Fprint(os.Stderr, "Error: Miss input file path")
		os.Exit(1)
	}

	postFilePath := nonFlagArgs[0]

	req := OGRenRequest{
		InputPath:     postFilePath,
		IconImagePath: path.Join(dir, IconImagePath),
		OutputDir:     path.Join(dir, OutputDir),
	}

	ogImagePath := req.Do()

	pattern := regexp.MustCompile(`^.*/static(.*)$`)

	// 使用正規表達式替換
	resultString := pattern.ReplaceAllString(ogImagePath, "$1")

	fmt.Println(resultString)
}
