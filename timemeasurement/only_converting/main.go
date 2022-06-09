package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	images := 20
	benchmark_time_entries := 20 //Anzahl Versionen pro Image
	benchmark_time_jpg := make([][]int, images)
	for k := range benchmark_time_jpg {
		benchmark_time_jpg[k] = make([]int, benchmark_time_entries)
	}
	benchmark_time_webp := make([][]int, images)
	for k := range benchmark_time_webp {
		benchmark_time_webp[k] = make([]int, benchmark_time_entries)
	}

	for j := 1; j <= images; j++ {

		for i := 0; i < benchmark_time_entries; i++ {
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

			args_jpg := []string{"-quality", strconv.Itoa(quali), "../../../Ursprungsdateien/" + strconv.Itoa(j) + "." + file_endung + "[0]", "-background", "white", "-alpha", "remove", "-auto-orient", "-resize", "1920x1920", "+profile", "'!exif,!xmp,!iptc,!8bim,*'", "-strip", "-units", "PixelsPerInch", "-density", "72", "../../../Comparison_time_converting/output_" + strconv.Itoa(j) + "_" + strconv.Itoa(quali) + "_jpg.jpg"}
			args_webp := []string{"-quality", strconv.Itoa(quali), "../../../Ursprungsdateien/" + strconv.Itoa(j) + "." + file_endung + "[0]", "-auto-orient", "-resize", "1920x1920", "+profile", "'!exif,!xmp,!iptc,!8bim,*'", "-strip", "-units", "PixelsPerInch", "-density", "72", "../../../Comparison_time_converting/output_" + strconv.Itoa(j) + "_" + strconv.Itoa(quali) + "_webp.webp"}

			start := time.Now()
			cmd_jpg := exec.Command(app, args_jpg...)
			_, err := cmd_jpg.Output()
			if err != nil {
				fmt.Println(err)
			}
			elapsed := int(time.Since(start))
			benchmark_time_jpg[j-1][i] = int(elapsed)

			start = time.Now()
			cmd_webp := exec.Command(app, args_webp...)
			_, err = cmd_webp.Output()
			if err != nil {
				fmt.Println(err)
			}
			elapsed = int(time.Since(start))
			benchmark_time_webp[j-1][i] = int(elapsed)
		}
	}

	file_jpg, err := os.OpenFile("../../../Comparison_time_converting/output_time_converting_jpg.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	datawriter_jpg := bufio.NewWriter(file_jpg)
	for k := 0; k < images; k++ {
		for t := 0; t < benchmark_time_entries; t++ {
			_, err = datawriter_jpg.WriteString(strconv.Itoa(k+1) + "\t" + strconv.Itoa((t+1)*5) + "\t" + strconv.Itoa(benchmark_time_jpg[k][t]))
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

	file_webp, err := os.OpenFile("../../../Comparison_time_converting/output_time_converting_webp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	datawriter_webp := bufio.NewWriter(file_webp)
	for k := 0; k < images; k++ {
		for t := 0; t < benchmark_time_entries; t++ {
			_, err = datawriter_webp.WriteString(strconv.Itoa(k+1) + "\t" + strconv.Itoa((t+1)*5) + "\t" + strconv.Itoa(benchmark_time_webp[k][t]))
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
