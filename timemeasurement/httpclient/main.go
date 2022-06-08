package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	//---------------------------------- set the flags ----------------------------------
	http_url_flag := flag.String("http_url", "http://localhost:4040", "the address")
	quali_flag := flag.Int("quali", 85, "the quali returned from the server")
	flag.Parse()
	http_url := *http_url_flag
	quali := *quali_flag

	//---------------------------------- define variables ----------------------------------

	images := 20
	var benchmark_slice_jpg = make([]int, images)
	var benchmark_slice_webp = make([]int, images)
	var filename_upload, file_endung string

	//---------------------------------- Time Measurements ----------------------------------

	for i := 1; i <= images; i++ {
		switch i {
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
		filename_upload = strconv.Itoa(i) + "." + file_endung

		//Stream a file to the server
		start := time.Now()
		clientsidestreamingclient(http_url, "/clientsidestreaming_jpg?filename="+filename_upload+"&quali="+strconv.Itoa(quali), "../../../Ursprungsdateien/"+filename_upload)
		//Stream a file from the server
		serversidestreamingclient(http_url, "/serversidestreaming?filename="+strconv.Itoa(i)+".jpg", "../httpclient/downloadedfiles/"+strconv.Itoa(i)+".jpg")
		elapsed := int(time.Since(start))
		benchmark_slice_jpg[i-1] = int(elapsed)

		//Stream a file to the server
		start = time.Now()
		clientsidestreamingclient(http_url, "/clientsidestreaming_webp?filename="+filename_upload+"&quali="+strconv.Itoa(quali), "../../../Ursprungsdateien/"+filename_upload)
		//Stream a file from the server
		serversidestreamingclient(http_url, "/serversidestreaming?filename="+strconv.Itoa(i)+".webp", "../httpclient/downloadedfiles/"+strconv.Itoa(i)+".webp")
		elapsed = int(time.Since(start))
		benchmark_slice_webp[i-1] = int(elapsed)
	}

	//---------------------------------- print benchmark results to files ----------------------------------

	file, err := os.OpenFile("../results/results_"+strconv.Itoa(quali)+"_jpg.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)
	for i := 1; i <= images; i++ {
		_, _ = datawriter.WriteString(strconv.Itoa(i) + "\t" + strconv.Itoa(benchmark_slice_jpg[i-1]) + "\n")
	}
	datawriter.Flush()
	file.Close()

	file_2, err := os.OpenFile("../results/results_"+strconv.Itoa(quali)+"_webp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter_2 := bufio.NewWriter(file_2)
	for i := 1; i <= images; i++ {
		_, _ = datawriter_2.WriteString(strconv.Itoa(i) + "\t" + strconv.Itoa(benchmark_slice_webp[i-1]) + "\n")
	}
	datawriter_2.Flush()
	file_2.Close()

}

//---------------------------------- client funcs ----------------------------------

func clientsidestreamingclient(http_url string, endpoint string, filepath string) {
	//open file
	r, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//send request
	resp, err := http.Post(http_url+endpoint, "multipart/form-data", r)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()
	return
}

func serversidestreamingclient(http_url string, endpoint string, filepath string) {
	//get the data
	resp, err := http.Get(http_url + endpoint)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()
	//create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer out.Close()
	//write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return
}
