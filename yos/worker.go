package yos

import (
	"log"
	"math"
	"time"
)

type WorkerAction func() error

type WorkItem struct {
	action WorkerAction
}

func retryWithExponentialBackoff(action WorkerAction, maxRetries int, initialDelay time.Duration, backoffFactor float64) error {
	retryCount := 0
	for {
		err := action()
		if err == nil {
			return nil
		}

		retryCount++
		if retryCount > maxRetries {
			return err
		}

		delay := time.Duration(float64(initialDelay) * math.Pow(backoffFactor, float64(retryCount-1)))
		time.Sleep(delay)
	}
}

func Worker(workQueue <-chan WorkItem, done chan<- bool) {
	for item := range workQueue {
		err := retryWithExponentialBackoff(item.action, 5, retryWaitTime, 1.25)

		if err != nil {
			log.Printf("Failed workitem: %v", err)
		}
	}

	done <- true
}
