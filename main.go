package main

import (
	"fmt"
	"time"
)

type File struct {
	time time.Time
	desc string
	name string
}

type Folder struct {
	time  time.Time
	desc  string
	name  string
	files []File
}

type User struct {
	name    string
	folders []Folder
}

func main() {
	f := File{time.Now(), "description file 1", "file name 1"}
	folder := Folder{time.Now(), "description folder 2", "folder name 2", []File{f}}
	user1 := User{"firstUser", []Folder{folder}}
	users := []User{user1}

	fmt.Printf("time: %s, description: %s, file name: %s \n", &f.time, f.desc, f.name)
	fmt.Printf("time: %s, description: %s, folder name: %s, contains: %s \n", &folder.time, folder.desc, folder.name, folder.files[0].name)
	fmt.Printf("username: %s, contains: %s", users[0].name, users[0].folders[0].name)
}
