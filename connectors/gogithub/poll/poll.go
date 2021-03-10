package poll

import "time"

func Scan(actions []string, repository string) {
	for range time.Tick(time.Minute * 1) {
		poll()
	}

}

func poll() {

}
