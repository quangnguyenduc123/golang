package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {

	// calling http.Get() which gives a response and an error value.
	resp, err := http.Get(url)

	// error handling
	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		c <- s

	} else {
		// a good practice is to close the Response Body if there is one
		// If you forget to close it there will be a resource leak.
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d  \n", url, resp.StatusCode)

		// fetching the page if HTTP Status Code is 200 (success)
		if resp.StatusCode == 200 {

			// The resp.Body implements the io.Reader interface
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			// Creating the file's name
			file := strings.Split(url, "//")[1]
			file += ".txt" // and I concatenate .txt to file value

			fmt.Printf("Writing response Body to %s\n", file)

			// Writing the response Body to File
			// If the file doesn't exist WriteFile() creates it and if it already exists
			// the function will truncate it before writing to file.
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				s += "Error writting file"
				c <- s
			}

			s += fmt.Sprintf("%s is UP!\n", url)
			c <- s
		}
	}
}

func checkUrl(url string, c chan string) {

	// calling http.Get() which gives a response and an error value.
	resp, err := http.Get(url)

	// error handling
	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		c <- s

	} else {
		// a good practice is to close the Response Body if there is one
		// If you forget to close it there will be a resource leak.
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d  \n", url, resp.StatusCode)
		fmt.Println(s)
		c <- url

		// fetching the page if HTTP Status Code is 200 (success)
		if resp.StatusCode == 200 {
			s += fmt.Sprintf("%s is UP!\n", url)
			fmt.Println(s)
			c <- url
		}
	}
}

func main() {
	urls := []string{"https://www.golang.org", "https://www.google.com"}

	c := make(chan string)

	// Iterating over the URLs and call the function for each URL
	for _, url := range urls {
		go checkUrl(url, c)
	}

	fmt.Println("Number of goroutines ", runtime.NumGoroutine())

	for {
		go checkUrl(<-c, c)
		fmt.Println(strings.Repeat("#", 20))
	}
}
