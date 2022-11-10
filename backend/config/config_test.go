package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
	"gopkg.in/yaml.v3"
)

func Test(t *testing.T) {
	t.Run("Merge test", func(t *testing.T) {
		type msi = map[string]interface{}

		m1 := msi{
			"a": msi{"b": "c"},
			"c": "d",
		}
		m2 := msi{
			"b": "d",
			"a": msi{
				"e": "f",
			},
			"c": msi{
				"f": "g",
			},
		}

		mergedMap := merge(m1, m2)

		resultantMap := msi{
			"a": msi{
				"b": "c",
				"e": "f",
			},
			"b": "d",
			"c": msi{
				"f": "g",
			},
		}

		assert.Equal(t, mergedMap, resultantMap)
	})

	t.Run("Base Config and Env Specific config merging check", func(t *testing.T) {
		dbHost := "random_value_for_testing"
		t.Setenv("DB_HOST", dbHost)
		conf, err := loadConfig()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, "test", conf.Environment)
		assert.Equal(t, dbHost, conf.Db.Host)
	})
	t.Run("handleEnvVarAndDefaultValues works properly", func(t *testing.T) {
		yamlContent := "x: ${X:x}\ny: ${Y:-y}\nz: ${Z:}\na: ${A}\nb: ${B}\nc: ${C:-c}\nd: ${D:-d}"
		// Gets the env var
		t.Setenv("X", "notX")

		// Gets the default var
		// t.Setenv("Z", "")

		// Gets the default val as its empty
		t.Setenv("Y", "")

		// C gets the default val as its not defined
		// t.Setenv("C", "")
		// D gets the env value as its defined
		t.Setenv("D", "notD")

		// Gets the env var, whatever it is, if it is not defined gets empty
		// ${B} wont be defined above
		t.Setenv("A", "notA")
		// t.Setenv("B", "")

		conf := make(map[string]string)
		err := yaml.Unmarshal(handleEnvVarAndDefaultValues([]byte(yamlContent)), &conf)
		if err != nil {
			panic(err)
		}
		expected := map[string]string{
			"a": "notA",
			"b": "",
			"c": "c",
			"d": "notD",
			"x": "notX",
			"y": "y",
			"z": "",
		}
		assert.Equal(t, expected, conf)
	})
}

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")
	goleak.VerifyTestMain(m)
}
