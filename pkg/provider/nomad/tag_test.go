package nomad

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tagsToLabels(t *testing.T) {
	testCases := []struct {
		desc     string
		tags     []string
		prefix   string
		expected map[string]string
	}{
		{
			desc:     "no tags",
			tags:     []string{},
			prefix:   "apache4",
			expected: map[string]string{},
		},
		{
			desc:   "minimal global config",
			tags:   []string{"apache4.enable=false"},
			prefix: "apache4",
			expected: map[string]string{
				"apache4.enable": "false",
			},
		},
		{
			desc: "config with domain",
			tags: []string{
				"apache4.enable=true",
				"apache4.domain=example.com",
			},
			prefix: "apache4",
			expected: map[string]string{
				"apache4.enable": "true",
				"apache4.domain": "example.com",
			},
		},
		{
			desc: "config with custom prefix",
			tags: []string{
				"custom.enable=true",
				"custom.domain=example.com",
			},
			prefix: "custom",
			expected: map[string]string{
				"apache4.enable": "true",
				"apache4.domain": "example.com",
			},
		},
		{
			desc: "config with spaces in tags",
			tags: []string{
				"custom.enable = true",
				"custom.domain = example.com",
			},
			prefix: "custom",
			expected: map[string]string{
				"apache4.enable": "true",
				"apache4.domain": "example.com",
			},
		},
		{
			desc:   "with a prefix",
			prefix: "test",
			tags: []string{
				"test.aaa=01",
				"test.bbb=02",
				"ccc=03",
				"test.ddd=04=to",
			},
			expected: map[string]string{
				"apache4.aaa": "01",
				"apache4.bbb": "02",
				"apache4.ddd": "04=to",
			},
		},
		{
			desc:   "with an empty prefix",
			prefix: "",
			tags: []string{
				"test.aaa=01",
				"test.bbb=02",
				"ccc=03",
				"test.ddd=04=to",
			},
			expected: map[string]string{
				"apache4.test.aaa": "01",
				"apache4.test.bbb": "02",
				"apache4.ccc":      "03",
				"apache4.test.ddd": "04=to",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			labels := tagsToLabels(test.tags, test.prefix)

			assert.Equal(t, test.expected, labels)
		})
	}
}
