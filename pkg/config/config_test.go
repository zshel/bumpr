package config_test

import (
	"github.com/analogj/go-util/utils"
	"github.com/packagrio/bumpr/pkg/config"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestConfiguration_init_ShouldCorrectlyInitializeConfiguration(t *testing.T) {
	t.Parallel()

	//setup
	defer utils.UnsetEnv("PACKAGR_")()

	//test
	testConfig, _ := config.Create()

	//assert
	require.Equal(t, "generic", testConfig.GetString(config.PACKAGR_PACKAGE_TYPE), "should populate package_type with generic default")
	require.Equal(t, "default", testConfig.GetString(config.PACKAGR_SCM), "should populate scm with default")
	require.Equal(t, "patch", testConfig.GetString(config.PACKAGR_VERSION_BUMP_TYPE), "should populate runner with default")
}

func TestConfiguration_init_EnvVariablesShouldLoadProperly(t *testing.T) {
	//setup
	os.Setenv("PACKAGR_VERSION_BUMP_TYPE", "major")

	//test
	testConfig, _ := config.Create()

	//assert
	require.Equal(t, "major", testConfig.GetString(config.PACKAGR_VERSION_BUMP_TYPE), "should populate Engine Version Bump Type from environmental variable")

	//teardown
	os.Unsetenv("PACKAGR_VERSION_BUMP_TYPE")
}

func TestConfiguration_ReadConfig(t *testing.T) {
	//setup
	testConfig, _ := config.Create()
	testConfig.SetDefault(config.PACKAGR_PACKAGE_TYPE, "generic")
	testConfig.SetDefault(config.PACKAGR_SCM, "default")
	testConfig.SetDefault(config.PACKAGR_VERSION_BUMP_TYPE, "patch")

	//test
	err := testConfig.ReadConfig(path.Join("testdata", "simple_overrides.yml"))

	//assert
	require.NoErrorf(t, err, "No error")
	require.Equal(t, "golang", testConfig.GetString(config.PACKAGR_PACKAGE_TYPE), "should populate Package Type from overrides config file")
	require.Equal(t, "github", testConfig.GetString(config.PACKAGR_SCM), "should populate SCM from overrides config file")
	require.Equal(t, "major", testConfig.GetString(config.PACKAGR_VERSION_BUMP_TYPE), "should populate Engine Version Bump Type from overrides config file")

}
