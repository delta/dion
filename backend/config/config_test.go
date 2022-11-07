package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	// this is so that the current working directory will
	// be root directory, it's causing problem in tests,
	// so had to do this, refer to the package to see what this does
	_ "delta.nitt.edu/dion/testing"
)

func Test(t *testing.T) {
	t.Run("Merge test", func(t *testing.T) {
		type msi = map[string]interface{}

		m1 := msi{
			"a": msi{"b": "c"},
		}
		m2 := msi{
			"b": "d",
			"a": msi{
				"e": "f",
			},
		}

		mergedMap := merge(m1, m2)

		resultantMap := msi{
			"a": msi{
				"b": "c",
				"e": "f",
			},
			"b": "d",
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
}

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")
	goleak.VerifyTestMain(m)
}
