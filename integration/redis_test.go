package integration

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/kvtools/redis"
	"github.com/kvtools/valkeyrie"
	"github.com/kvtools/valkeyrie/store"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/apache4/apache4/v3/integration/try"
	"github.com/apache4/apache4/v3/pkg/api"
)

// Redis test suites.
type RedisSuite struct {
	BaseSuite
	kvClient       store.Store
	redisEndpoints []string
}

func TestRedisSuite(t *testing.T) {
	suite.Run(t, new(RedisSuite))
}

func (s *RedisSuite) SetupSuite() {
	s.BaseSuite.SetupSuite()

	s.createComposeProject("redis")
	s.composeUp()

	s.redisEndpoints = []string{}
	s.redisEndpoints = append(s.redisEndpoints, net.JoinHostPort(s.getComposeServiceIP("redis"), "6379"))

	kv, err := valkeyrie.NewStore(
		s.T().Context(),
		redis.StoreName,
		s.redisEndpoints,
		&redis.Config{},
	)
	require.NoError(s.T(), err, "Cannot create store redis")

	s.kvClient = kv

	// wait for redis
	err = try.Do(60*time.Second, try.KVExists(kv, "test"))
	require.NoError(s.T(), err)
}

func (s *RedisSuite) TearDownSuite() {
	s.BaseSuite.TearDownSuite()
}

func (s *RedisSuite) TestSimpleConfiguration() {
	file := s.adaptFile("fixtures/redis/simple.toml", struct{ RedisAddress string }{
		RedisAddress: strings.Join(s.redisEndpoints, ","),
	})

	data := map[string]string{
		"apache4/http/routers/Router0/entryPoints/0": "web",
		"apache4/http/routers/Router0/middlewares/0": "compressor",
		"apache4/http/routers/Router0/middlewares/1": "striper",
		"apache4/http/routers/Router0/service":       "simplesvc",
		"apache4/http/routers/Router0/rule":          "Host(`kv1.localhost`)",
		"apache4/http/routers/Router0/priority":      "42",
		"apache4/http/routers/Router0/tls":           "true",

		"apache4/http/routers/Router1/rule":                 "Host(`kv2.localhost`)",
		"apache4/http/routers/Router1/priority":             "42",
		"apache4/http/routers/Router1/tls/domains/0/main":   "aaa.localhost",
		"apache4/http/routers/Router1/tls/domains/0/sans/0": "aaa.aaa.localhost",
		"apache4/http/routers/Router1/tls/domains/0/sans/1": "bbb.aaa.localhost",
		"apache4/http/routers/Router1/tls/domains/1/main":   "bbb.localhost",
		"apache4/http/routers/Router1/tls/domains/1/sans/0": "aaa.bbb.localhost",
		"apache4/http/routers/Router1/tls/domains/1/sans/1": "bbb.bbb.localhost",
		"apache4/http/routers/Router1/entryPoints/0":        "web",
		"apache4/http/routers/Router1/service":              "simplesvc",

		"apache4/http/services/simplesvc/loadBalancer/servers/0/url": "http://10.0.1.1:8888",
		"apache4/http/services/simplesvc/loadBalancer/servers/1/url": "http://10.0.1.1:8889",

		"apache4/http/services/srvcA/loadBalancer/servers/0/url": "http://10.0.1.2:8888",
		"apache4/http/services/srvcA/loadBalancer/servers/1/url": "http://10.0.1.2:8889",

		"apache4/http/services/srvcB/loadBalancer/servers/0/url": "http://10.0.1.3:8888",
		"apache4/http/services/srvcB/loadBalancer/servers/1/url": "http://10.0.1.3:8889",

		"apache4/http/services/mirror/mirroring/service":           "simplesvc",
		"apache4/http/services/mirror/mirroring/mirrors/0/name":    "srvcA",
		"apache4/http/services/mirror/mirroring/mirrors/0/percent": "42",
		"apache4/http/services/mirror/mirroring/mirrors/1/name":    "srvcB",
		"apache4/http/services/mirror/mirroring/mirrors/1/percent": "42",

		"apache4/http/services/Service03/weighted/services/0/name":   "srvcA",
		"apache4/http/services/Service03/weighted/services/0/weight": "42",
		"apache4/http/services/Service03/weighted/services/1/name":   "srvcB",
		"apache4/http/services/Service03/weighted/services/1/weight": "42",

		"apache4/http/middlewares/compressor/compress":            "true",
		"apache4/http/middlewares/striper/stripPrefix/prefixes/0": "foo",
		"apache4/http/middlewares/striper/stripPrefix/prefixes/1": "bar",
	}

	for k, v := range data {
		err := s.kvClient.Put(s.T().Context(), k, []byte(v), nil)
		require.NoError(s.T(), err)
	}

	s.apache4Cmd(withConfigFile(file))

	// wait for apache4
	err := try.GetRequest("http://127.0.0.1:8080/api/rawdata", 2*time.Second,
		try.BodyContains(`"striper@redis":`, `"compressor@redis":`, `"srvcA@redis":`, `"srvcB@redis":`),
	)
	require.NoError(s.T(), err)

	resp, err := http.Get("http://127.0.0.1:8080/api/rawdata")
	require.NoError(s.T(), err)

	var obtained api.RunTimeRepresentation
	err = json.NewDecoder(resp.Body).Decode(&obtained)
	require.NoError(s.T(), err)
	got, err := json.MarshalIndent(obtained, "", "  ")
	require.NoError(s.T(), err)

	expectedJSON := filepath.FromSlash("testdata/rawdata-redis.json")

	if *updateExpected {
		err = os.WriteFile(expectedJSON, got, 0o666)
		require.NoError(s.T(), err)
	}

	expected, err := os.ReadFile(expectedJSON)
	require.NoError(s.T(), err)

	if !bytes.Equal(expected, got) {
		diff := difflib.UnifiedDiff{
			FromFile: "Expected",
			A:        difflib.SplitLines(string(expected)),
			ToFile:   "Got",
			B:        difflib.SplitLines(string(got)),
			Context:  3,
		}

		text, err := difflib.GetUnifiedDiffString(diff)
		require.NoError(s.T(), err, text)
	}
}
