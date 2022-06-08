package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

func main() {

	//flags
	port_address_flag := flag.String("port_address", ":4040", "the port_address")
	flag.Parse()
	port_address := *port_address_flag

	//define endpoints
	http.HandleFunc("/clientsidestreaming_jpg", clientsidestreamingHandler_jpg)
	http.HandleFunc("/clientsidestreaming_webp", clientsidestreamingHandler_webp)
	http.HandleFunc("/serversidestreaming", serversidestreamingHandler)

	//start server
	fmt.Printf("starting server at port" + port_address + "\n")
	if err := http.ListenAndServe(port_address, nil); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func clientsidestreamingHandler_jpg(w http.ResponseWriter, r *http.Request) {
	var quali int
	var filename_without_ending, filename string
	var err error

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//get filename
		filename = r.FormValue("filename")
		quali, err = strconv.Atoi(r.FormValue("quali"))
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		filename_without_ending = filename[:strings.IndexByte(filename, '.')]
		//create file
		out, err := os.Create("../httpserver/uploadedfiles/" + filename)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		defer out.Close()
		w.WriteHeader(http.StatusOK)
		//write to file
		_, err = io.Copy(out, r.Body)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		defer wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		//convert file
		app := "magick"
		args_jpg := []string{"-quality", strconv.Itoa(quali), "../httpserver/uploadedfiles/" + filename + "[0]", "-background", "white", "-alpha", "remove", "-auto-orient", "-resize", "1920x1920", "+profile", "'!exif,!xmp,!iptc,!8bim,*'", "-strip", "-units", "PixelsPerInch", "-density", "72", "../httpserver/uploadedfiles/converted/" + filename_without_ending + ".jpg"}
		cmd := exec.Command(app, args_jpg...)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		defer wg.Done()
	}()
	wg.Wait()
}

func clientsidestreamingHandler_webp(w http.ResponseWriter, r *http.Request) {
	var quali int
	var filename_without_ending, filename string
	var err error

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//get filename
		filename = r.FormValue("filename")
		quali, err = strconv.Atoi(r.FormValue("quali"))
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		filename_without_ending = filename[:strings.IndexByte(filename, '.')]
		//create file
		out, err := os.Create("../httpserver/uploadedfiles/" + filename)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		defer out.Close()
		w.WriteHeader(http.StatusOK)
		//write to file
		_, err = io.Copy(out, r.Body)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		defer wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		//convert file
		app := "magick"
		args_webp := []string{"-quality", strconv.Itoa(quali), "../httpserver/uploadedfiles/" + filename + "[0]", "-auto-orient", "-resize", "1920x1920", "+profile", "'!exif,!xmp,!iptc,!8bim,*'", "-strip", "-units", "PixelsPerInch", "-density", "72", "../httpserver/uploadedfiles/converted/" + filename_without_ending + ".webp"}
		cmd := exec.Command(app, args_webp...)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		defer wg.Done()
	}()
	wg.Wait()
}

func serversidestreamingHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//get filename
		var filename string = r.FormValue("filename")
		//open file
		file, err := os.Open("../httpserver/uploadedfiles/converted/" + filename)
		if err != nil {
			fmt.Fprintf(w, "could not find file")
			log.Println(err)
		}
		defer file.Close()
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		w.WriteHeader(http.StatusOK)
		//send file as stream
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "could not read body", http.StatusInternalServerError)
		}
		defer wg.Done()
	}()
	wg.Wait()
}
