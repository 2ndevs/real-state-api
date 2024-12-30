package libs

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(value any) {
	bytes, _ := json.MarshalIndent(value, "", "  ")

	output := string(bytes)
	fmt.Println(output)
}
