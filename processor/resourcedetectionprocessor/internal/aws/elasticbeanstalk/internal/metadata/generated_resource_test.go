// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetCloudPlatform("cloud.platform-val")
			rb.SetCloudProvider("cloud.provider-val")
			rb.SetDeploymentEnvironment("deployment.environment-val")
			rb.SetServiceInstanceID("service.instance.id-val")
			rb.SetServiceVersion("service.version-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch test {
			case "default":
				assert.Equal(t, 5, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 5, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("cloud.platform")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "cloud.platform-val", val.Str())
			}
			val, ok = res.Attributes().Get("cloud.provider")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "cloud.provider-val", val.Str())
			}
			val, ok = res.Attributes().Get("deployment.environment")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "deployment.environment-val", val.Str())
			}
			val, ok = res.Attributes().Get("service.instance.id")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "service.instance.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("service.version")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "service.version-val", val.Str())
			}
		})
	}
}
