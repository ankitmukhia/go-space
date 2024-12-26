package main

/* https://www.boot.dev/lessons/9c6f7fc9-8e7a-4857-8b10-70f130410700 */

import "fmt"

func waitForDBs(numDBs int, dbChan chan struct{}) {
	//? initial code, <-dbChan // this is only droping on database, coz there is multiple db
	// solution
	for i := 0; i < numDBs; i++ {
	   <-dbChan
	}	
}

// don't touch below this line

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			// this is not nested struct, this is defining an initilizing same time.
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}

