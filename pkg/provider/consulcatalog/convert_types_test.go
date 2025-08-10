package consulcatalog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tagsToNeutralLabels(t *testing.T) {
	testCases := []struct {
		desc     string
		tags     []string
		prefix   string
		expected map[string]string
	}{
		{
			desc:     "without tags",
			expected: nil,
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

			labels := tagsToNeutralLabels(test.tags, test.prefix)

			assert.Equal(t, test.expected, labels)
		})
	}
}
