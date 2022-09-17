package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func brightness(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	value := query.Get("value")
	if value == "" {
		fmt.Fprintf(w, "Missing value")
	} else {
		v, err := strconv.Atoi(value)
		if err == nil && v >= 0 && v <= 100 {
			cmd := exec.Command("xbacklight", "-set", value)
			_, err := cmd.Output()
			if err == nil {
				fmt.Fprintf(w, "OK\n")
			} else {
				fmt.Fprintf(w, err.Error())
			}
		} else {
			fmt.Fprintf(w, "Bad value: "+value+" (must be int 0 <= x <= 100)")
		}
	}
}
func main() {
	port := 8090
	if len(os.Args) == 2 {
		p, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid argument: " + os.Args[1])
			os.Exit(1)
		}
		port = p
	}

	http.HandleFunc("/brightness", brightness)

	fmt.Println("Listening on 0.0.0.0:" + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
