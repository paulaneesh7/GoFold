package waitgroups

import (
	"fmt"
	"net/http"
	"sync"
)


// usually these are pointers but in this case they are values
var Wg sync.WaitGroup 

func GetStatusCode(endpoint string) {

	// this tell that we are done with the current go routine
	defer Wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Status Code: ", res.StatusCode)
}