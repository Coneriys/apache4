package kv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	testCases := []struct {
		desc     string
		rootName string
		pairs    map[string]string
		expected *sample
	}{
		{
			desc:     "simple case",
			rootName: "apache4",
			pairs: map[string]string{
				"apache4/fielda":        "bar",
				"apache4/fieldb":        "1",
				"apache4/fieldc":        "true",
				"apache4/fieldd/0":      "one",
				"apache4/fieldd/1":      "two",
				"apache4/fielde":        "",
				"apache4/fieldf/Test1":  "A",
				"apache4/fieldf/Test2":  "B",
				"apache4/fieldg/0/name": "A",
				"apache4/fieldg/1/name": "B",
				"apache4/fieldh/":       "foo",
			},
			expected: &sample{
				FieldA: "bar",
				FieldB: 1,
				FieldC: true,
				FieldD: []string{"one", "two"},
				FieldE: &struct {
					Name string
				}{},
				FieldF: map[string]string{
					"Test1": "A",
					"Test2": "B",
				},
				FieldG: []sub{
					{Name: "A"},
					{Name: "B"},
				},
				FieldH: "foo",
			},
		},
		{
			desc:     "multi-level root name",
			rootName: "foo/bar/apache4",
			pairs: map[string]string{
				"foo/bar/apache4/fielda":        "bar",
				"foo/bar/apache4/fieldb":        "2",
				"foo/bar/apache4/fieldc":        "true",
				"foo/bar/apache4/fieldd/0":      "one",
				"foo/bar/apache4/fieldd/1":      "two",
				"foo/bar/apache4/fielde":        "",
				"foo/bar/apache4/fieldf/Test1":  "A",
				"foo/bar/apache4/fieldf/Test2":  "B",
				"foo/bar/apache4/fieldg/0/name": "A",
				"foo/bar/apache4/fieldg/1/name": "B",
				"foo/bar/apache4/fieldh/":       "foo",
			},
			expected: &sample{
				FieldA: "bar",
				FieldB: 2,
				FieldC: true,
				FieldD: []string{"one", "two"},
				FieldE: &struct {
					Name string
				}{},
				FieldF: map[string]string{
					"Test1": "A",
					"Test2": "B",
				},
				FieldG: []sub{
					{Name: "A"},
					{Name: "B"},
				},
				FieldH: "foo",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			element := &sample{}

			err := Decode(mapToPairs(test.pairs), element, test.rootName)
			require.NoError(t, err)

			assert.Equal(t, test.expected, element)
		})
	}
}

type sample struct {
	FieldA string
	FieldB int
	FieldC bool
	FieldD []string
	FieldE *struct {
		Name string
	} `kv:"allowEmpty"`
	FieldF map[string]string
	FieldG []sub
	FieldH string
}

type sub struct {
	Name string
}
