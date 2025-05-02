package main

import (
	"bancho-kill/compliance_tests"
)

func main() {
	compliance_tests.RunAllComplianceTests("http://c.staging.titanic.sh")
}
