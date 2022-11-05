package config_test

import (
	"os"
	"testing"

	"delta.nitt.edu/dion/config"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	os.Setenv("ENV", "test")
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

		mergedMap := config.Merge(m1, m2)

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
		conf, err := config.LoadConfig()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, conf.Environment, "test")
		assert.Equal(t, conf.Db.Host, dbHost)
	})
}
