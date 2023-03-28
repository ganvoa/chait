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
  rol1: "You are a dinasour"
  rol2: "You are a vegetable"
  replies: 3`

		expectedConfig := &internal.Config{}
		expectedConfig.Chait.Replies = 3
		expectedConfig.Chait.Rol1 = "You are a dinasour"
		expectedConfig.Chait.Rol2 = "You are a vegetable"

		conf, err := internal.NewConfig(strings.NewReader(testConfig))
		if err != nil {
			t.Fatalf("got error %v, expected none", err)
		}

		assert.Equal(t, expectedConfig, conf)
	})

	t.Run("Replies > 0", func(t *testing.T) {

		testConfig := `
chait:
  rol1: "You are a dinasour"
  rol2: "You are a vegetable"
`
		_, err := internal.NewConfig(strings.NewReader(testConfig))
		assert.EqualError(t, err, "replies must be > 0")
	})

	t.Run("Rol1 required", func(t *testing.T) {

		testConfig := `
chait:
  rol2: "You are a vegetable"
  replies: 3
`
		_, err := internal.NewConfig(strings.NewReader(testConfig))
		assert.EqualError(t, err, "rol1 required")
	})

	t.Run("Rol2 required", func(t *testing.T) {

		testConfig := `
chait:
  rol1: "You are a dinasour"
  replies: 3
`
		_, err := internal.NewConfig(strings.NewReader(testConfig))
		assert.EqualError(t, err, "rol2 required")
	})

	t.Run("Invalid File", func(t *testing.T) {

		testConfig := `
chait:  rol1: "You are a dinasour"replies: 3
`
		_, err := internal.NewConfig(strings.NewReader(testConfig))
		assert.EqualError(t, err, "error decoding yaml")
	})

}
