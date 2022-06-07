package main

import (
	"fmt"
	"os/exec"
	"strconv"
)


func main() {

	for i := 0; i < 20; i++ {
		quali := (i + 1) * 5

		app := "magick"

		args_jpg := []string{"-quality", strconv.Itoa(quali), "../webpthumbnails/Ursprungsdateien/wallpaper.png[0]", "-background", "white", "-alpha", "remove", "-auto-orient", "-resize", "1920x1920", "+profile", "'!exif,!xmp,!iptc,!8bim,*'", "-strip", "-units", "PixelsPerInch", "-density", "72", "../webpthumbnails/Comparison/output_" + strconv.Itoa(quali) + ".jpg"}
		args_webp := []string{"-quality", strconv.Itoa(quali), "../webpthumbnails/Ursprungsdateien/wallpaper.png[0]", "-auto-orient", "-resize", "1920x1920", "+profile", "'!exif,!xmp,!iptc,!8bim,*'", "-strip", "-units", "PixelsPerInch", "-density", "72", "../webpthumbnails/Comparison/output_" + strconv.Itoa(quali) + ".webp"}

		cmd_jpg := exec.Command(app, args_jpg...)
		cmd_webp := exec.Command(app, args_webp...)
		_, err := cmd_jpg.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		_, err = cmd_webp.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
