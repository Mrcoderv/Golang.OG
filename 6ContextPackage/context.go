package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)  // stop on 2 second 
	defer cancel()

	select {   // case for the operation completre or not .  
	case <-time.After(3 * time.Second):  
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Println("Operation timed out:", ctx.Err())
	}
}
