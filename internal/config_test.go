package internal_test

import (
	"strings"
	"testing"

	"github.com/ganvoa/chait/internal"
	"github.com/stretchr/testify/assert"
)

func Test_ReadConfigFile(t *testing.T) {

	t.Run("No Error", func(t *testing.T) {

		testConfig := `
chait:
  roleU1: "You are a dinasour"
  roleU2: "You are a vegetable"
  iterations: 3`

		expectedConfig := &internal.Config{}
		expectedConfig.Chait.Iterations = 3
		expectedConfig.Chait.RoleU1 = "You are a dinasour"
		expectedConfig.Chait.RoleU2 = "You are a vegetable"

		conf, err := internal.NewConfig(strings.NewReader(testConfig))
		if err != nil {
			t.Fatalf("got error %v, expected none", err)
		}

		assert.Equal(t, expectedConfig, conf)
	})

	t.Run("iterations > 0", func(t *testing.T) {

		testConfig := `
chait:
  roleU1: "You are a dinasour"
  roleU2: "You are a vegetable"
`
		_, err := internal.NewConfig(strings.NewReader(testConfig))
		assert.EqualError(t, err, "iterations must be > 0")
	})

	t.Run("Rol1 required", func(t *testing.T) {

		testConfig := `
chait:
  roleU2: "You are a vegetable"
  iterations: 3
`
		_, err := internal.NewConfig(strings.NewReader(testConfig))
		assert.EqualError(t, err, "roleU1 required")
	})

	t.Run("RoleU2 required", func(t *testing.T) {

		testConfig := `
chait:
  roleU1: "You are a dinasour"
  iterations: 3
`
		_, err := internal.NewConfig(strings.NewReader(testConfig))
		assert.EqualError(t, err, "roleU2 required")
	})

	t.Run("Invalid File", func(t *testing.T) {

		testConfig := `
chait:  roleU1: "You are a dinasour"iterations: 3
`
		_, err := internal.NewConfig(strings.NewReader(testConfig))
		assert.EqualError(t, err, "error decoding yaml")
	})

}
