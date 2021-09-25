package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type counters struct {
	sync.Mutex
	view  int
	click int
}

var (
	currentCounter       *counters
	sportCounter         = counters{}
	entertainmentCounter = counters{}
	businessCounter      = counters{}
	educationCounter     = counters{}
	mapCounter           map[string]string

	content = []string{"sports", "entertainment", "business", "education"}
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to EQ Works 😎")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//get random value from content list
	data := content[rand.Intn(len(content))]
	key := data + ": " + string(time.Now().Format("2006.01.02 15:04"))
	value := mapCounter[key]

	switch data {
	case "sports":
		currentCounter = &sportCounter
		processView(&sportCounter, data, key)
		fmt.Fprint(w, string(key), string(value))

	case "entertainment":
		currentCounter = &entertainmentCounter
		processView(&entertainmentCounter, data, key)
		fmt.Fprint(w, string(key), string(value))

	case "business":
		currentCounter = &businessCounter
		processView(&businessCounter, data, key)
		fmt.Fprint(w, string(key), string(value))

	case "education":
		currentCounter = &educationCounter
		processView(&educationCounter, data, key)
		fmt.Fprint(w, string(key), string(value))
	}
	// simulate random click call
	if rand.Intn(100) < 50 {
		processClick(currentCounter, data)
	}

	//c.Lock()
	//c.view++
	//c.Unlock()

	err := processRequest(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

}

func processRequest(r *http.Request) error {
	time.Sleep(time.Duration(rand.Int31n(50)) * time.Millisecond)
	return nil
}

func processClick(c *counters, data string) error {
	c.Lock()
	c.click++
	c.Unlock()

	return nil
}
func processView(c *counters, data string, key string) error {
	c.Lock()
	_, currentTimeIsInMap := mapCounter[key]

	if currentTimeIsInMap == false {
		c.view = 0
		c.click = 0
	}
	c.view++
	value := " view: " + strconv.Itoa(c.view) + " click: " + strconv.Itoa(c.click)
	mapCounter[key] = value

	c.Unlock()
	return nil
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	if !isAllowed() {
		w.WriteHeader(429)
		return
	}
}

func isAllowed() bool {
	return true
}

func uploadCounters() error {
	return nil
}

func main() {
	mapCounter = make(map[string]string)

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/stats/", statsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
