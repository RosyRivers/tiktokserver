package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func SaveCover(InputPath string, filmname string) (CoverName string) {
	pngName := fmt.Sprintf("%spng", filmname[0:len(filmname)-3])
	saveFile := filepath.Join("./public/cover", pngName)
	cmd := exec.Command("ffmpeg", "-i", InputPath, "-ss", "00:00:00.001", "-vframes", "1", saveFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//cmd.Dir = "/"
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("First frame extracted and saved as image:", saveFile)
	return pngName
}
