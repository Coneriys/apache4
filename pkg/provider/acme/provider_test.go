package acme

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/stretchr/testify/assert"
	"github.com/apache4/apache4/v3/pkg/safe"
	"github.com/apache4/apache4/v3/pkg/types"
)

func TestGetUncheckedCertificates(t *testing.T) {
	t.Skip("Needs TLS Manager")
	wildcardMap := make(map[string]*tls.Certificate)
	wildcardMap["*.apache4.wtf"] = &tls.Certificate{}

	wildcardSafe := &safe.Safe{}
	wildcardSafe.Set(wildcardMap)

	domainMap := make(map[string]*tls.Certificate)
	domainMap["apache4.wtf"] = &tls.Certificate{}

	domainSafe := &safe.Safe{}
	domainSafe.Set(domainMap)

	// TODO Add a test for DefaultCertificate
	testCases := []struct {
		desc             string
		dynamicCerts     *safe.Safe
		resolvingDomains map[string]struct{}
		acmeCertificates []*CertAndStore
		domains          []string
		expectedDomains  []string
	}{
		{
			desc:            "wildcard to generate",
			domains:         []string{"*.apache4.wtf"},
			expectedDomains: []string{"*.apache4.wtf"},
		},
		{
			desc:            "wildcard already exists in dynamic certificates",
			domains:         []string{"*.apache4.wtf"},
			dynamicCerts:    wildcardSafe,
			expectedDomains: nil,
		},
		{
			desc:    "wildcard already exists in ACME certificates",
			domains: []string{"*.apache4.wtf"},
			acmeCertificates: []*CertAndStore{
				{
					Certificate: Certificate{
						Domain: types.Domain{Main: "*.apache4.wtf"},
					},
				},
			},
			expectedDomains: nil,
		},
		{
			desc:            "domain CN and SANs to generate",
			domains:         []string{"apache4.wtf", "foo.apache4.wtf"},
			expectedDomains: []string{"apache4.wtf", "foo.apache4.wtf"},
		},
		{
			desc:            "domain CN already exists in dynamic certificates and SANs to generate",
			domains:         []string{"apache4.wtf", "foo.apache4.wtf"},
			dynamicCerts:    domainSafe,
			expectedDomains: []string{"foo.apache4.wtf"},
		},
		{
			desc:    "domain CN already exists in ACME certificates and SANs to generate",
			domains: []string{"apache4.wtf", "foo.apache4.wtf"},
			acmeCertificates: []*CertAndStore{
				{
					Certificate: Certificate{
						Domain: types.Domain{Main: "apache4.wtf"},
					},
				},
			},
			expectedDomains: []string{"foo.apache4.wtf"},
		},
		{
			desc:            "domain already exists in dynamic certificates",
			domains:         []string{"apache4.wtf"},
			dynamicCerts:    domainSafe,
			expectedDomains: nil,
		},
		{
			desc:    "domain already exists in ACME certificates",
			domains: []string{"apache4.wtf"},
			acmeCertificates: []*CertAndStore{
				{
					Certificate: Certificate{
						Domain: types.Domain{Main: "apache4.wtf"},
					},
				},
			},
			expectedDomains: nil,
		},
		{
			desc:            "domain matched by wildcard in dynamic certificates",
			domains:         []string{"who.apache4.wtf", "foo.apache4.wtf"},
			dynamicCerts:    wildcardSafe,
			expectedDomains: nil,
		},
		{
			desc:    "domain matched by wildcard in ACME certificates",
			domains: []string{"who.apache4.wtf", "foo.apache4.wtf"},
			acmeCertificates: []*CertAndStore{
				{
					Certificate: Certificate{
						Domain: types.Domain{Main: "*.apache4.wtf"},
					},
				},
			},
			expectedDomains: nil,
		},
		{
			desc:    "root domain with wildcard in ACME certificates",
			domains: []string{"apache4.wtf", "foo.apache4.wtf"},
			acmeCertificates: []*CertAndStore{
				{
					Certificate: Certificate{
						Domain: types.Domain{Main: "*.apache4.wtf"},
					},
				},
			},
			expectedDomains: []string{"apache4.wtf"},
		},
		{
			desc:    "all domains already managed by ACME",
			domains: []string{"apache4.wtf", "foo.apache4.wtf"},
			resolvingDomains: map[string]struct{}{
				"apache4.wtf":     {},
				"foo.apache4.wtf": {},
			},
			expectedDomains: []string{},
		},
		{
			desc:    "one domain already managed by ACME",
			domains: []string{"apache4.wtf", "foo.apache4.wtf"},
			resolvingDomains: map[string]struct{}{
				"apache4.wtf": {},
			},
			expectedDomains: []string{"foo.apache4.wtf"},
		},
		{
			desc:    "wildcard domain already managed by ACME checks the domains",
			domains: []string{"bar.apache4.wtf", "foo.apache4.wtf"},
			resolvingDomains: map[string]struct{}{
				"*.apache4.wtf": {},
			},
			expectedDomains: []string{},
		},
		{
			desc:    "wildcard domain already managed by ACME checks domains and another domain checks one other domain, one domain still unchecked",
			domains: []string{"apache4.wtf", "bar.apache4.wtf", "foo.apache4.wtf", "acme.wtf"},
			resolvingDomains: map[string]struct{}{
				"*.apache4.wtf": {},
				"apache4.wtf":   {},
			},
			expectedDomains: []string{"acme.wtf"},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.resolvingDomains == nil {
				test.resolvingDomains = make(map[string]struct{})
			}

			acmeProvider := Provider{
				// certificateStore: &apache4tls.CertificateStore{
				// 	DynamicCerts: test.dynamicCerts,
				// },
				certificates:     test.acmeCertificates,
				resolvingDomains: test.resolvingDomains,
			}

			domains := acmeProvider.getUncheckedDomains(t.Context(), test.domains, "default")
			assert.Len(t, domains, len(test.expectedDomains), "Unexpected domains.")
		})
	}
}

func TestProvider_sanitizeDomains(t *testing.T) {
	testCases := []struct {
		desc            string
		domains         types.Domain
		dnsChallenge    *DNSChallenge
		expectedErr     string
		expectedDomains []string
	}{
		{
			desc:            "valid wildcard",
			domains:         types.Domain{Main: "*.apache4.wtf"},
			dnsChallenge:    &DNSChallenge{},
			expectedErr:     "",
			expectedDomains: []string{"*.apache4.wtf"},
		},
		{
			desc:            "no wildcard",
			domains:         types.Domain{Main: "apache4.wtf", SANs: []string{"foo.apache4.wtf"}},
			dnsChallenge:    &DNSChallenge{},
			expectedErr:     "",
			expectedDomains: []string{"apache4.wtf", "foo.apache4.wtf"},
		},
		{
			desc:            "no domain",
			domains:         types.Domain{},
			dnsChallenge:    nil,
			expectedErr:     "no domain was given",
			expectedDomains: nil,
		},
		{
			desc:            "unauthorized wildcard with SAN",
			domains:         types.Domain{Main: "*.*.apache4.wtf", SANs: []string{"foo.apache4.wtf"}},
			dnsChallenge:    &DNSChallenge{},
			expectedErr:     "unable to generate a wildcard certificate in ACME provider for domain \"*.*.apache4.wtf,foo.apache4.wtf\" : ACME does not allow '*.*' wildcard domain",
			expectedDomains: nil,
		},
		{
			desc:            "wildcard and SANs",
			domains:         types.Domain{Main: "*.apache4.wtf", SANs: []string{"apache4.wtf"}},
			dnsChallenge:    &DNSChallenge{},
			expectedErr:     "",
			expectedDomains: []string{"*.apache4.wtf", "apache4.wtf"},
		},
		{
			desc:            "wildcard SANs",
			domains:         types.Domain{Main: "*.apache4.wtf", SANs: []string{"*.acme.wtf"}},
			dnsChallenge:    &DNSChallenge{},
			expectedErr:     "",
			expectedDomains: []string{"*.apache4.wtf", "*.acme.wtf"},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			acmeProvider := Provider{Configuration: &Configuration{DNSChallenge: test.dnsChallenge}}

			domains, err := acmeProvider.sanitizeDomains(t.Context(), test.domains)

			if len(test.expectedErr) > 0 {
				assert.EqualError(t, err, test.expectedErr, "Unexpected error.")
			} else {
				assert.Len(t, domains, len(test.expectedDomains), "Unexpected domains.")
			}
		})
	}
}

func TestDeleteUnnecessaryDomains(t *testing.T) {
	testCases := []struct {
		desc            string
		domains         []types.Domain
		expectedDomains []types.Domain
	}{
		{
			desc: "no domain to delete",
			domains: []types.Domain{
				{
					Main: "acme.wtf",
					SANs: []string{"apache4.acme.wtf", "foo.bar"},
				},
				{
					Main: "*.foo.acme.wtf",
				},
				{
					Main: "acme02.wtf",
					SANs: []string{"apache4.acme02.wtf", "bar.foo"},
				},
			},
			expectedDomains: []types.Domain{
				{
					Main: "acme.wtf",
					SANs: []string{"apache4.acme.wtf", "foo.bar"},
				},
				{
					Main: "*.foo.acme.wtf",
					SANs: []string{},
				},
				{
					Main: "acme02.wtf",
					SANs: []string{"apache4.acme02.wtf", "bar.foo"},
				},
			},
		},
		{
			desc: "wildcard and root domain",
			domains: []types.Domain{
				{
					Main: "acme.wtf",
				},
				{
					Main: "*.acme.wtf",
					SANs: []string{"acme.wtf"},
				},
			},
			expectedDomains: []types.Domain{
				{
					Main: "acme.wtf",
					SANs: []string{},
				},
				{
					Main: "*.acme.wtf",
					SANs: []string{},
				},
			},
		},
		{
			desc: "2 equals domains",
			domains: []types.Domain{
				{
					Main: "acme.wtf",
					SANs: []string{"apache4.acme.wtf", "foo.bar"},
				},
				{
					Main: "acme.wtf",
					SANs: []string{"apache4.acme.wtf", "foo.bar"},
				},
			},
			expectedDomains: []types.Domain{
				{
					Main: "acme.wtf",
					SANs: []string{"apache4.acme.wtf", "foo.bar"},
				},
			},
		},
		{
			desc: "2 domains with same values",
			domains: []types.Domain{
				{
					Main: "acme.wtf",
					SANs: []string{"apache4.acme.wtf"},
				},
				{
					Main: "acme.wtf",
					SANs: []string{"apache4.acme.wtf", "foo.bar"},
				},
			},
			expectedDomains: []types.Domain{
				{
					Main: "acme.wtf",
					SANs: []string{"apache4.acme.wtf"},
				},
				{
					Main: "foo.bar",
					SANs: []string{},
				},
			},
		},
		{
			desc: "domain totally checked by wildcard",
			domains: []types.Domain{
				{
					Main: "who.acme.wtf",
					SANs: []string{"apache4.acme.wtf", "bar.acme.wtf"},
				},
				{
					Main: "*.acme.wtf",
				},
			},
			expectedDomains: []types.Domain{
				{
					Main: "*.acme.wtf",
					SANs: []string{},
				},
			},
		},
		{
			desc: "duplicated wildcard",
			domains: []types.Domain{
				{
					Main: "*.acme.wtf",
					SANs: []string{"acme.wtf"},
				},
				{
					Main: "*.acme.wtf",
				},
			},
			expectedDomains: []types.Domain{
				{
					Main: "*.acme.wtf",
					SANs: []string{"acme.wtf"},
				},
			},
		},
		{
			desc: "domain partially checked by wildcard",
			domains: []types.Domain{
				{
					Main: "apache4.acme.wtf",
					SANs: []string{"acme.wtf", "foo.bar"},
				},
				{
					Main: "*.acme.wtf",
				},
				{
					Main: "who.acme.wtf",
					SANs: []string{"apache4.acme.wtf", "bar.acme.wtf"},
				},
			},
			expectedDomains: []types.Domain{
				{
					Main: "acme.wtf",
					SANs: []string{"foo.bar"},
				},
				{
					Main: "*.acme.wtf",
					SANs: []string{},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			domains := deleteUnnecessaryDomains(t.Context(), test.domains)
			assert.Equal(t, test.expectedDomains, domains, "unexpected domain")
		})
	}
}

func TestIsAccountMatchingCaServer(t *testing.T) {
	testCases := []struct {
		desc       string
		accountURI string
		serverURI  string
		expected   bool
	}{
		{
			desc:       "acme staging with matching account",
			accountURI: "https://acme-staging-v02.api.letsencrypt.org/acme/acct/1234567",
			serverURI:  "https://acme-staging-v02.api.letsencrypt.org/acme/directory",
			expected:   true,
		},
		{
			desc:       "acme production with matching account",
			accountURI: "https://acme-v02.api.letsencrypt.org/acme/acct/1234567",
			serverURI:  "https://acme-v02.api.letsencrypt.org/acme/directory",
			expected:   true,
		},
		{
			desc:       "http only acme with matching account",
			accountURI: "http://acme.api.letsencrypt.org/acme/acct/1234567",
			serverURI:  "http://acme.api.letsencrypt.org/acme/directory",
			expected:   true,
		},
		{
			desc:       "different subdomains for account and server",
			accountURI: "https://test1.example.org/acme/acct/1234567",
			serverURI:  "https://test2.example.org/acme/directory",
			expected:   false,
		},
		{
			desc:       "different domains for account and server",
			accountURI: "https://test.example1.org/acme/acct/1234567",
			serverURI:  "https://test.example2.org/acme/directory",
			expected:   false,
		},
		{
			desc:       "different tld for account and server",
			accountURI: "https://test.example.com/acme/acct/1234567",
			serverURI:  "https://test.example.org/acme/directory",
			expected:   false,
		},
		{
			desc:       "malformed account url",
			accountURI: "//|\\/test.example.com/acme/acct/1234567",
			serverURI:  "https://test.example.com/acme/directory",
			expected:   false,
		},
		{
			desc:       "malformed server url",
			accountURI: "https://test.example.com/acme/acct/1234567",
			serverURI:  "//|\\/test.example.com/acme/directory",
			expected:   false,
		},
		{
			desc:       "malformed server and account url",
			accountURI: "//|\\/test.example.com/acme/acct/1234567",
			serverURI:  "//|\\/test.example.com/acme/directory",
			expected:   false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			result := isAccountMatchingCaServer(t.Context(), test.accountURI, test.serverURI)

			assert.Equal(t, test.expected, result)
		})
	}
}

func TestInitAccount(t *testing.T) {
	testCases := []struct {
		desc            string
		account         *Account
		email           string
		keyType         string
		expectedAccount *Account
	}{
		{
			desc: "Existing account with all information",
			account: &Account{
				Email:   "foo@foo.net",
				KeyType: certcrypto.EC256,
			},
			expectedAccount: &Account{
				Email:   "foo@foo.net",
				KeyType: certcrypto.EC256,
			},
		},
		{
			desc:    "Account nil",
			email:   "foo@foo.net",
			keyType: "EC256",
			expectedAccount: &Account{
				Email:   "foo@foo.net",
				KeyType: certcrypto.EC256,
			},
		},
		{
			desc: "Existing account with no email",
			account: &Account{
				KeyType: certcrypto.RSA4096,
			},
			email:   "foo@foo.net",
			keyType: "EC256",
			expectedAccount: &Account{
				Email:   "foo@foo.net",
				KeyType: certcrypto.EC256,
			},
		},
		{
			desc: "Existing account with no key type",
			account: &Account{
				Email: "foo@foo.net",
			},
			email:   "bar@foo.net",
			keyType: "EC256",
			expectedAccount: &Account{
				Email:   "foo@foo.net",
				KeyType: certcrypto.EC256,
			},
		},
		{
			desc: "Existing account and provider with no key type",
			account: &Account{
				Email: "foo@foo.net",
			},
			email: "bar@foo.net",
			expectedAccount: &Account{
				Email:   "foo@foo.net",
				KeyType: certcrypto.RSA4096,
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			acmeProvider := Provider{account: test.account, Configuration: &Configuration{Email: test.email, KeyType: test.keyType}}

			actualAccount, err := acmeProvider.initAccount(t.Context())
			assert.NoError(t, err, "Init account in error")
			assert.Equal(t, test.expectedAccount.Email, actualAccount.Email, "unexpected email account")
			assert.Equal(t, test.expectedAccount.KeyType, actualAccount.KeyType, "unexpected keyType account")
		})
	}
}

func Test_getCertificateRenewDurations(t *testing.T) {
	testCases := []struct {
		desc                  string
		certificatesDurations int
		expectRenewPeriod     time.Duration
		expectRenewInterval   time.Duration
	}{
		{
			desc:                  "Less than 24 Hours certificates: 20 minutes renew period, 1 minutes renew interval",
			certificatesDurations: 1,
			expectRenewPeriod:     time.Minute * 20,
			expectRenewInterval:   time.Minute,
		},
		{
			desc:                  "1 Year certificates: 4 months renew period, 1 week renew interval",
			certificatesDurations: 24 * 365,
			expectRenewPeriod:     time.Hour * 24 * 30 * 4,
			expectRenewInterval:   time.Hour * 24 * 7,
		},
		{
			desc:                  "265 Days certificates: 30 days renew period, 1 day renew interval",
			certificatesDurations: 24 * 265,
			expectRenewPeriod:     time.Hour * 24 * 30,
			expectRenewInterval:   time.Hour * 24,
		},
		{
			desc:                  "90 Days certificates: 30 days renew period, 1 day renew interval",
			certificatesDurations: 24 * 90,
			expectRenewPeriod:     time.Hour * 24 * 30,
			expectRenewInterval:   time.Hour * 24,
		},
		{
			desc:                  "30 Days certificates: 10 days renew period, 12 hour renew interval",
			certificatesDurations: 24 * 30,
			expectRenewPeriod:     time.Hour * 24 * 10,
			expectRenewInterval:   time.Hour * 12,
		},
		{
			desc:                  "7 Days certificates: 1 days renew period, 1 hour renew interval",
			certificatesDurations: 24 * 7,
			expectRenewPeriod:     time.Hour * 24,
			expectRenewInterval:   time.Hour,
		},
		{
			desc:                  "24 Hours certificates: 6 hours renew period, 10 minutes renew interval",
			certificatesDurations: 24,
			expectRenewPeriod:     time.Hour * 6,
			expectRenewInterval:   time.Minute * 10,
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			renewPeriod, renewInterval := getCertificateRenewDurations(test.certificatesDurations)
			assert.Equal(t, test.expectRenewPeriod, renewPeriod)
			assert.Equal(t, test.expectRenewInterval, renewInterval)
		})
	}
}
