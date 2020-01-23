package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const (
	ntpPool       string        = "pool.ntp.org"
	maxRetries    int           = 5
	sleepDuration 				= time.Second * 5
)

// retry function retries the given function for the given number of times
func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; i < attempts; i++ {
		err = f()
		if err != nil {
			log.Println("ERROR: ", err)
			log.Printf("Error is retryable. Sleeping for %d seconds and retrying...\n", sleep/1e9)
			time.Sleep(sleep)
		} else {
			return err
		}
	}
	return fmt.Errorf("FATAL ERROR after %d attempts, last error: %s", attempts, err)
}

// printNTPtime function uses a subset of NTP, called SNTP, to print current time.
// To make our local system clock reading more accurate, we add ClockOffset,
// which is the estimated offset relative to the server's clock.
func printNTPtime() (err error) {
	response, err := ntp.Query(ntpPool)
	if err != nil {
		return err
	}
	curTime := time.Now().Add(response.ClockOffset)
	fmt.Println(curTime)
	return err
}

func main() {
	err := retry(maxRetries, sleepDuration, printNTPtime)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
