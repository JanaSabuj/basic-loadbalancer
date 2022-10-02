package customerror

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Printf("Err detected:!! %v\n", err)
	os.Exit(1)
}
