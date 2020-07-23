package configuration

import (
	"github.com/ditrit/gandalf/core/models"
	"os"
	"testing"
)

func TestSetStringKeyConfig(t *testing.T) {

	err := SetStringKeyConfig("test", "test_str_key", "t1", "", "test string usage", false)
	if err != nil {
		t.Errorf("Not expecting an error, but got :%v", err)
	}

	//error because we redefine an existing key
	err = SetStringKeyConfig("test", "test_str_key", "t1", "toto", "test string usage", false)
	if err == nil {
		t.Errorf("Expected error: %v", err)
	}
}

func TestSetIntegerKeyConfig(t *testing.T) {
	err := SetIntegerKeyConfig("test", "test_int_key", "t2", -1, "test integer usage", false)
	if err != nil {
		t.Errorf("Not expecting an error, but got :%v", err)
	}
	//error because we redefine an existing key
	err = SetIntegerKeyConfig("test", "test_int_key", "t2", -1, "test integer usage", false)
	if err == nil {
		t.Errorf("Expected error: %v", err)
	}
}

func TestArgParse(t *testing.T) {
	InitMainConfigKeys()
	//error caused for parsing a non integer value
	t.Run("Arg parse test 1", func(t *testing.T) {
		var test = []string{"-t2", "invalid value", "-l", "toto", "-g", "cluster"}
		err := argParse("test config", test)
		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})
	//error from an unknown CLI parameter
	t.Run("Arg parse test 2", func(t *testing.T) {
		var test2 = []string{"-invalidFlag", "10", "-l", "toto", "-g", "cluster"}
		err := argParse("test config", test2)
		if err == nil {
			t.Errorf("expected error: %v", err)
		}
	})

	//Expected case
	t.Run("Arg parse test 3", func(t *testing.T) {
		var test3 = []string{"-t2", "10", "-l", "toto", "-g", "connector"}
		err := argParse("test config", test3)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
	})
}

func TestEnvParse(t *testing.T) {
	//error caused for parsing a non integer value
	t.Run("Env parse test 1", func(t *testing.T) {
		_ = os.Setenv("GANDALF_TEST_INT_KEY", "invalid value")
		err := envParse()
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
		_ = os.Unsetenv("GANDALF_TEST_INT_KEY")
	})
	//expected case
	t.Run("Env parse test 2", func(t *testing.T) {
		_ = os.Setenv("GANDALF_TEST_STR_KEY", "test env string")
		err := envParse()
		_ = os.Unsetenv("GANDALF_TEST_STR_KEY")
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
	})
}

func TestYamlFileToMap(t *testing.T) {
	t.Run("test yaml to map with default value", func(t *testing.T) {
		_, err := yamlFileToMap()
		if err != nil {
			t.Errorf("Expected error : %v", err)
		}
	})

	//error because we give an invalid path
	t.Run("test yaml to map with wrong path", func(t *testing.T) {
		var wrongPathTest = []string{"-t2", "10", "-l", "toto", "-g", "connector", "-f", "/core/"}
		err := argParse("test config", wrongPathTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		_, err = yamlFileToMap()
		if err == nil {
			t.Errorf("Expected error : %v", err)
		}
	})

	//error because we try to unmarshal a wrong file
	t.Run("test yaml to map with wrong directory", func(t *testing.T) {
		var errorInFileTest = []string{"-t2", "10", "-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/testing/"}
		err := argParse("test config", errorInFileTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		_, err = yamlFileToMap()
		if err == nil {
			t.Errorf("Expected error : %v", err)
		}
	})
}

func TestYamlFileParse(t *testing.T) {

	//error when parsing a value of the yaml file
	t.Run("yaml parsing error", func(t *testing.T) {
		var wrongPathTest = []string{"-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/"}
		err := argParse("test config", wrongPathTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		err = yamlFileParse()
		if err == nil {
			t.Errorf("Expected error : %v", err)
		}
	})

	//error when parsing a value of the yaml file
	t.Run("yaml parsing error", func(t *testing.T) {
		var wrongPathTest = []string{"-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/"}
		err := argParse("test config", wrongPathTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		err = yamlFileParse()
		if err == nil {
			t.Errorf("Expected error : %v", err)
		}
	})

	//expected case
	t.Run("yaml parsing test", func(t *testing.T) {
		var yamlTest = []string{"-t2", "10", "-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/"}
		err := argParse("test config", yamlTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		err = yamlFileParse()
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
	})
}

func TestDefaultParse(t *testing.T) {
	err := defaultParse()
	if err != nil {
		t.Errorf("Not expecting an error, but got : %v", err)
	}
}

func TestParseConfig(t *testing.T) {
	//error while parsing CLI parameters
	t.Run("arg parse error", func(t *testing.T) {
		var parseConfigTest = []string{"-t2", "test", "-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/"}
		err := ParseConfig("global parse test", parseConfigTest)
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})

	//error while parsing environment variables
	t.Run("env parse error", func(t *testing.T) {
		var parseConfigTest = []string{"-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/"}
		_ = os.Setenv("GANDALF_TEST_INT_KEY", "test")
		err := ParseConfig("global parse test", parseConfigTest)
		_ = os.Unsetenv("GANDALF_TEST_INT_KEY")
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})

	//error while parsing the yaml file
	t.Run("yaml parse error", func(t *testing.T) {
		var parseConfigTest = []string{"-l", "toto", "-g", "connector", "-f", "test_file.yaml"}
		err := ParseConfig("global parse test", parseConfigTest)
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})

	//expected case
	t.Run("Expected parse", func(t *testing.T) {
		var parseConfigTest = []string{"-t2", "10", "-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/"}
		err := ParseConfig("global parse test", parseConfigTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
	})
}

func TestWorkerKeyParse(t *testing.T) {
	var testIntList = []models.ConfigurationKeys{
		{"test_int_worker","2","integer",true},
	}
	t.Run("env parse error", func(t *testing.T) {
		_ = os.Setenv("GANDALF_TEST_INT_WORKER", "test")
		err := WorkerKeyParse(testIntList)
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
		_ = os.Unsetenv("GANDALF_TEST_INT_WORKER")
	})

	t.Run("yaml parse error", func(t *testing.T) {
		var parseConfigTest = []string{"-l", "toto", "-g", "connector", "-f", "test_file.yaml"}
		err := ParseConfig("global parse test", parseConfigTest)
		err = WorkerKeyParse(testIntList)
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})
}

func TestKeysValidation(t *testing.T) {
	//Error while parsing the file into a map
	t.Run("yaml parsing error", func(t *testing.T) {
		var wrongKeyTest = []string{"-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/wrongTest/"}
		err := argParse("test config", wrongKeyTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		err = KeysValidation()
		if err == nil {
			t.Errorf("Expected error : %v", err)
		}
	})

	t.Run("env not needed key error", func(t *testing.T) {
		var wrongKeyTest = []string{"-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/"}
		err := argParse("test config", wrongKeyTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		_ = os.Setenv("GANDALF_NOT_NEEDED_KEY", "test")
		err = KeysValidation()
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		_ = os.Unsetenv("GANDALF_NOT_NEEDED_KEY")
	})

	//error when parsing a key not existing in gandalf configuration
	t.Run("yaml parsing not needed key error", func(t *testing.T) {
		var wrongKeyTest = []string{"-l", "toto", "-g", "connector", "-f", homePath + "/gandalf/core/configuration/test/testingKey/"}
		err := argParse("test config", wrongKeyTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		err = KeysValidation()
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
	})
}

func TestGetVersionsList(t *testing.T) {
	t.Run("unexpected GetVersionsList", func(t *testing.T) {
		var parseConfigTest = []string{"-t2", "10", "-l", "toto", "-g", "connector", "-v", "test,test2", "-f", homePath + "/gandalf/core/configuration/test/"}
		err := ParseConfig("global parse test", parseConfigTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		//giving GetVersionsList an invalid value
		gandalfVersionsString, _ := GetStringConfig("versions")
		_, err = GetVersionsList(gandalfVersionsString)
		if err == nil {
			t.Errorf("Not expecting an error, but got: %v", err)
		}
	})

	//expected case
	t.Run("expected GetVersionsList", func(t *testing.T) {
		var parseConfigTest = []string{"-t2", "10", "-l", "toto", "-g", "connector", "-v", "1,2", "-f", homePath + "/gandalf/core/configuration/test/"}
		err := ParseConfig("global parse test", parseConfigTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		gandalfVersionsString, _ := GetStringConfig("versions")
		_, err = GetVersionsList(gandalfVersionsString)
		if err != nil {
			t.Errorf("Not expecting an error, but got: %v", err)
		}
	})

}

func TestIsConfigValid(t *testing.T) {
	//expected case
	t.Run("Valid config test", func(t *testing.T) {
		err := IsConfigValid()
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
	})

	//error because of an invalid gandalf type component
	t.Run("Invalid gandalf type test", func(t *testing.T) {
		var gandalfTypeTest = []string{"-t2", "10", "-l", "toto", "-g", "test"}
		err := argParse("test config", gandalfTypeTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		err = IsConfigValid()
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})

	//Error in a cluster type component
	t.Run("Invalid cluster test", func(t *testing.T) {
		var clusterTest = []string{"-t2", "10", "-l", "toto", "-g", "cluster"}
		err := argParse("test config", clusterTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		//setting an empty key to test configuration
		_ = SetStringKeyConfig("cluster", "test_cluster", "", "", "test cluster usage", true)
		err = IsConfigValid()
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})

	//Error in an aggregator component
	t.Run("Invalid aggregator test", func(t *testing.T) {
		var aggregatorTest = []string{"-t2", "10", "-l", "toto", "-g", "aggregator"}
		err := argParse("test config", aggregatorTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		//setting an empty key to test configuration
		_ = SetStringKeyConfig("aggregator", "test_aggregator", "", "", "test aggregator usage", true)
		err = IsConfigValid()
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})

	//Error in a connector component
	t.Run("Invalid connector test", func(t *testing.T) {
		var connectorTest = []string{"-t2", "10", "-l", "toto", "-g", "connector"}
		err := argParse("test config", connectorTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		//setting an empty key to test configuration
		_ = SetStringKeyConfig("connector", "test_connector", "", "", "test connector usage", true)
		err = IsConfigValid()
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})

	t.Run("Invalid worker test", func(t *testing.T) {
		var testList = []models.ConfigurationKeys{
			{"testWorker","","string",true},
		}
		var workerTest = []string{"-t2", "10", "-l", "toto", "-g", "connector","-test_connector","test","-t","tenantTest"}
		err := ParseConfig("test config", workerTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		err = WorkerKeyParse(testList)
		_,_ = GetStringConfig("testWorker")
		if err != nil {
			t.Errorf("Unexpected Error :%v", err)
		}
		err = IsConfigValid()
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})

	t.Run("Invalid worker test", func(t *testing.T) {
		var testList = []TestKey{
			{"testWorker","","string",true},
		}
		var workerTest = []string{"-t2", "10", "-l", "toto", "-g", "connector","-testConnector","test","-t","tenantTest"}
		err := ParseConfig("test config", workerTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		err = WorkerKeyParse(testList)
		_,_ = GetStringConfig("testWorker")
		if err != nil {
			t.Errorf("Unexpected Error :%v", err)
		}
		err = IsConfigValid()
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
	})
}

func TestConfigMain(t *testing.T){
		var configTest = []string{"-t2", "10", "-l", "toto", "-g", "cluster", "-f", homePath + "/gandalf/core/configuration/test/","-test_cluster","test"}
		ConfigMain("test config", configTest)
}

func TestGetTlS(t *testing.T) {
	//Error while getting the TLS files
	t.Run("fail getting a file", func(t *testing.T) {
		var fileErrorTest = []string{"-t2", "10", "-l", "toto", "-g", "connector", "-cert_pem", "/gandalf/core/certs"}
		err := argParse("test config", fileErrorTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		_, err = GetTLS()
		if err == nil {
			t.Errorf("Expected error : %v", err)
		}
	})

	//Get TLS using multiple paths
	t.Run("get TLS paths test", func(t *testing.T) {
		var pathMapTest = []string{ "-t2", "10",
									"-l", "toto",
									"-g", "connector",
									"-cert_dir", "gandalf/core/certs",
									"-cert_pem", homePath + "/gandalf/core/certs/cert.pem",
									"-key_pem", homePath + "/gandalf/core/certs/key.pem",
									"-ca_cert_pem", homePath + "/gandalf/core/certs/ca.pem",
									"-ca_key_pem", homePath + "/gandalf/core/certs/cakey.pem"}
		err := argParse("test config", pathMapTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		_, err = GetTLS()
		if err != nil {
			t.Errorf("Not expecting an error, but got: %v", err)
		}
	})

	//Get TLS using a directory
	t.Run("get TLS via directory test", func(t *testing.T) {
		var pathMapTest = []string{"-t2", "10", "-l", "toto", "-g", "connector", "-cert_dir", homePath + "/gandalf/core/certs"}
		err := argParse("test config", pathMapTest)
		if err != nil {
			t.Errorf("Not expecting an error, but got : %v", err)
		}
		_, err = GetTLS()
		if err != nil {
			t.Errorf("Not expecting an error, but got: %v", err)
		}
	})

}

func TestGetStringConfig(t *testing.T) {
	//Expected case
	_, err := GetStringConfig("test_str_key")

	if err != nil {
		t.Errorf("Not expecting an error, but got : %v", err)
	}

	//error for trying to get an Integer type Key with GetStringKey
	_, err = GetStringConfig("test_int_key")
	if err == nil {
		t.Errorf("Expected error : %v", err)
	}

	//error because of an unknown key
	_, err = GetStringConfig("test_string_key")
	if err == nil {
		t.Errorf("Expected error: %v", err)

	}
}

func TestGetIntegerConfig(t *testing.T) {
	//Expected case
	_, err := GetIntegerConfig("test_int_key")
	if err != nil {
		t.Errorf("Not expecting an error, but got : %v", err)
	}

	//error because of an unknown key
	_, err = GetIntegerConfig("test_integer_key")
	if err == nil {
		t.Errorf("Expected error: %v", err)
	}

	//error for trying to get a String type Key with GetIntegerKey
	_, err = GetIntegerConfig("test_str_key")
	if err == nil {
		t.Errorf("Expected error: %v", err)
	}

	tmp := "inconvertible value"
	ConfigKeys["test_int_conversion_key"] = configKey{&tmp, "test", "t3", "integer", "none", "fail convert test", false}
	//error for an invalid conversion
	_, err = GetIntegerConfig("test_int_conversion_key")
	if err == nil {
		t.Errorf("Expected error: %v", err)
	}

}
