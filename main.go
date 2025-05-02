package main

import (
	"bancho-kill/compliance_tests"
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))

	compliance_tests.RunAllComplianceTests("http://c.staging.titanic.sh")
}
