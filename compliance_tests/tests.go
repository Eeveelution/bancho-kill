package compliance_tests

func RunAllComplianceTests(addr string) {
	context := TestContext{
		CurrentTestNumber: 0,
		ServerAddress:     addr,
	}

	TestOsuLogin(&context)
}
