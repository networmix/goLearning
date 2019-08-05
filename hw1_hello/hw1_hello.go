package main

import (
    "fmt"
    "log"
    "time"

    "github.com/beevik/ntp"
)

const (
    ntpPool       string        = "pool.ntp.org"
    maxRetries    int           = 5
    sleepDuration time.Duration = 5e9
)

// retry function retries the given function for the given number of times
func retry(attempts int, sleep time.Duration, f func() error) (err error) {
    for i := 0; i < attempts; i++ {
        err = f()
        if err == nil {
            return err
        }
        log.Println("ERROR: ", err)
        log.Printf("Error is retryable. Sleeping for %ds and retrying...\n", sleep/1e9)
        time.Sleep(sleep)
    }
    return fmt.Errorf("FATAL ERROR after %d attempts, last error: %s", attempts, err)
}

// printNTPtime function uses a subset of NTP, called SNTP, to print current time.
func printNTPtime() (err error) {
    response, err := ntp.Query(ntpPool)
    if err == nil {
        time := time.Now().Add(response.ClockOffset)
        fmt.Println(time)
    }
    return err
}

func main() {
    retry(maxRetries, sleepDuration, printNTPtime)
}
