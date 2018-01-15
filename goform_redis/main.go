package main

import (
	"bytes"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	REDISHOST = "redis"
	TEMPLATE  = "/usr/local/goform/html/index.html.tpl"
)

func top(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Output template html
	list := getAllData()
	t := template.Must(template.ParseFiles(TEMPLATE))
	if err := t.ExecuteTemplate(w, "index.html.tpl", list); err != nil {
		log.Fatal(err)
	}

	// Write form data
	value := r.FormValue("key")
	if len(value) != 0 {
		writeData(value)
	}
}

func getAllData() []string {
	// Get all post data
	c, err := redis.Dial("tcp", REDISHOST+":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	list, err := redis.Strings(c.Do("LRANGE", "POST", 0, -1))
	if err != nil {
		log.Fatal(err)
	}

	return list
}

func writeData(data string) {
	// Connect redis server
	c, err := redis.Dial("tcp", REDISHOST+":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Get time of key and create value
	const layout = "2006-01-02-15:04:05"
	t := time.Now().Format(layout)

	var buffer bytes.Buffer
	buffer.WriteString(t)
	buffer.WriteString(" ")
	buffer.WriteString(data)

	// Write value
	_, serr := c.Do("RPUSH", "POST", buffer.String())
	if serr != nil {
		log.Fatal(serr)
	}
}

func main() {
	http.HandleFunc("/", top)

	fmt.Printf("Starting web server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
