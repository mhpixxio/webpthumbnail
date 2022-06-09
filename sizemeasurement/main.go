package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	images := 20
	benchmark_size_entries := 20 //Anzahl Versionen pro Image
	benchmark_size_jpg := make([][]int, images)
	for k := range benchmark_size_jpg {
		benchmark_size_jpg[k] = make([]int, benchmark_size_entries)
	}
	benchmark_size_webp := make([][]int, images)
	for k := range benchmark_size_webp {
		benchmark_size_webp[k] = make([]int, benchmark_size_entries)
	}

	for j := 1; j <= images; j++ {

		for i := 0; i < benchmark_size_entries; i++ {
			quali := (i + 1) * 5

			app := "magick"

			var file_endung string
			file_endung = "jpg"

			switch j {
			case 1:
				file_endung = "jfif"
			case 2:
				file_endung = "png"
			case 3:
				file_endung = "avif"
			case 4:
				file_endung = "webp"
			case 5:
				file_endung = "jpg"
			case 6:
				file_endung = "jpg"
			case 7:
				file_endung = "png"
			case 8:
				file_endung = "svg"
			case 9:
				file_endung = "jfif"
			case 10:
				file_endung = "jpg"
			case 11:
				file_endung = "jpeg"
			case 12:
				file_endung = "png"
			case 13:
				file_endung = "svg"
			case 14:
				file_endung = "jpg"
			case 15:
				file_endung = "jpg"
			case 16:
				file_endung = "jpg"
			case 17:
				file_endung = "jpg"
			case 18:
				file_endung = "CR2"
			case 19:
				file_endung = "png"
			case 20:
				file_endung = "tiff"
			}

			args_jpg := []string{"-quality", strconv.Itoa(quali), "../../Ursprungsdateien/" + strconv.Itoa(j) + "." + file_endung + "[0]", "-background", "white", "-alpha", "remove", "-auto-orient", "-resize", "1920x1920", "+profile", "'!exif,!xmp,!iptc,!8bim,*'", "-strip", "-units", "PixelsPerInch", "-density", "72", "../../Comparison_size/output_" + strconv.Itoa(j) + "_" + strconv.Itoa(quali) + "_jpg.jpg"}
			args_webp := []string{"-quality", strconv.Itoa(quali), "../../Ursprungsdateien/" + strconv.Itoa(j) + "." + file_endung + "[0]", "-auto-orient", "-resize", "1920x1920", "+profile", "'!exif,!xmp,!iptc,!8bim,*'", "-strip", "-units", "PixelsPerInch", "-density", "72", "../../Comparison_size/output_" + strconv.Itoa(j) + "_" + strconv.Itoa(quali) + "_webp.webp"}

			cmd_jpg := exec.Command(app, args_jpg...)
			cmd_webp := exec.Command(app, args_webp...)
			_, err := cmd_jpg.Output()
			if err != nil {
				fmt.Println(err)
			}
			_, err = cmd_webp.Output()
			if err != nil {
				fmt.Println(err)
			}

			info_jpg, err := os.Stat("../../Comparison_size/output_" + strconv.Itoa(j) + "_" + strconv.Itoa(quali) + "_jpg.jpg")
			if err != nil {
				fmt.Println(err)
			}
			size_jpg := info_jpg.Size()
			benchmark_size_jpg[j-1][i] = int(size_jpg)

			info_webp, err := os.Stat("../../Comparison_size/output_" + strconv.Itoa(j) + "_" + strconv.Itoa(quali) + "_webp.webp")
			if err != nil {
				fmt.Println(err)
			}
			size_webp := info_webp.Size()
			benchmark_size_webp[j-1][i] = int(size_webp)

		}
	}

	file_jpg, err := os.OpenFile("../../Comparison_size/output_jpg.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	datawriter_jpg := bufio.NewWriter(file_jpg)
	for k := 0; k < images; k++ {
		for t := 0; t < benchmark_size_entries; t++ {
			_, err = datawriter_jpg.WriteString(strconv.Itoa(k+1) + "\t" + strconv.Itoa((t+1)*5) + "\t" + strconv.Itoa(benchmark_size_jpg[k][t]))
			_, err2 := datawriter_jpg.WriteString("\n")
			if err != nil || err2 != nil {
				fmt.Println(err)
			}
		}
		_, err := datawriter_jpg.WriteString("\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	datawriter_jpg.Flush()
	file_jpg.Close()

	file_webp, err := os.OpenFile("../../Comparison_size/output_webp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	datawriter_webp := bufio.NewWriter(file_webp)
	for k := 0; k < images; k++ {
		for t := 0; t < benchmark_size_entries; t++ {
			_, err = datawriter_webp.WriteString(strconv.Itoa(k+1) + "\t" + strconv.Itoa((t+1)*5) + "\t" + strconv.Itoa(benchmark_size_webp[k][t]))
			_, err2 := datawriter_webp.WriteString("\n")
			if err != nil || err2 != nil {
				fmt.Println(err)
			}
		}
		_, err = datawriter_webp.WriteString("\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	datawriter_webp.Flush()
	file_webp.Close()
}
