package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	"image/color"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	selected int32    = 0
	labels   []string = []string{
		"small (1024x768)",
		"tv (1920x1080)",
		"tv + small (1920x1080 + 1024x768) ",
		"billys (1600x900)",
	}
	resolutions []string = []string{
		"output DVI-D-1 pos 0 0 res 1024x768",
		"output HDMI-A-1 pos 0 0 res 1920x1080",
		"output HDMI-A-1 pos 0 0 res 1920x1080\noutput DVI-D-1 pos 1920 0 res 1024x768",
		"output DVI-D-1 pos 0 0 res 1600x900",
	}
	colorBg   color.Color = color.RGBA{50, 50, 50, 255}
	colorText color.Color = color.RGBA{241, 93, 14, 255}
	fontSize  float32     = 32.0
)

func writeToDisk() {

	configFilePath := filepath.Join(os.Getenv("HOME"), ".config", "sway", "config")
	// Read the data of the file
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	fileText := string(data)
	truncatedText := fileText[:strings.Index(fileText, "#resolutions")]
	// Example: Append a new line to the file
	newLine := resolutions[selected]
	result := truncatedText + "\n" + "#resolutions" + "\n" + newLine
	// Write the modified content back to the file
	err = os.WriteFile(configFilePath, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	err = exec.Command("swaymsg", "reload").Run()
	if err != nil {
		fmt.Println("Error reloading sway:", err)
	}

	os.Exit(1)

}

func loop() {
	g.SingleWindow().Layout(
		g.Style().SetFontSize(fontSize).To(
			g.Style().SetColor(g.StyleColorText, colorText).To(
				g.Style().SetColor(g.StyleColorWindowBg, colorText).To(
					g.Align(g.AlignCenter).To(
						g.Combo("", labels[selected], labels, &selected),
						g.Button("Change").OnClick(writeToDisk),
					)))))
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 500, 250, g.MasterWindowFlagsNotResizable)
	g.Context.FontAtlas.SetDefaultFont("RubikDoodleShadow.ttf", fontSize)
	g.Context.FontAtlas.AddFont("RubikDoodleShadow.ttf", fontSize)
	g.PushColorWindowBg(colorBg)
	wnd.Run(loop)
}
