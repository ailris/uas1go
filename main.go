package main

import (
	"fmt"
	"net/http"
)

type Student struct {
	NIM      string
	Name     string
	Class    string
	Major    string
}

func main() {
	student := Student{
		NIM:      "2101160001",
		Name:     "Christine Wielia",
		Class:    "SI-TI-IT-5 M-1",
		Major:    "Jurusan Teknologi Informasi",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "NIM: %s\n", student.NIM)
		fmt.Fprintf(w, "Nama: %s\n", student.Name)
		fmt.Fprintf(w, "Kelas: %s\n", student.Class)
		fmt.Fprintf(w, "Jurusan: %s\n", student.Major)
	})

	http.ListenAndServe(":8080", nil)
}
