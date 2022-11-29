package ldap_api

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap"
	"log"
)

const (
	BindUsername = "admin@example.com"
	BindPassword = "Topdog88"
	FQDN         = "ldap.example.com"
	BaseDN       = "cn=ldap,dc=example,dc=com"
	Filter       = "(objectClass=*)"
)

// Ldap Connection without TLS
func Connect() (*ldap.Conn, error) {
	// You can also use IP instead of FQDN
	l, err := ldap.DialURL(fmt.Sprintf("ldap://%s:389", FQDN))
	if err != nil {
		return nil, err
	}

	return l, nil
}

// Ldap Connection with TLS
func ConnectTLS() (*ldap.Conn, error) {
	// You can also use IP instead of FQDN
	l, err := ldap.DialURL(fmt.Sprintf("ldaps://%s:636", FQDN))
	if err != nil {
		return nil, err
	}

	return l, nil
}

func ConnectTLSConfig() (*ldap.Conn, error) {
	// You can also use IP instead of FQDN
	l, err := ldap.DialURL(
		fmt.Sprintf("ldaps://%s:636", FQDN),
		ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	)
	if err != nil {
		return nil, err
	}

	return l, nil
}

// Anonymous Bind and Search
func AnonymousBindAndSearch(l *ldap.Conn) (*ldap.SearchResult, error) {
	l.UnauthenticatedBind("")

	anonReq := ldap.NewSearchRequest(
		"",
		ldap.ScopeBaseObject, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		Filter,
		[]string{},
		nil,
	)
	result, err := l.Search(anonReq)
	if err != nil {
		return nil, fmt.Errorf("Anonymous Bind Search Error: %s", err)
	}

	if len(result.Entries) > 0 {
		result.Entries[0].Print()
		return result, nil
	} else {
		return nil, fmt.Errorf("Couldn't fetch anonymous bind search entries")
	}
}

// Normal Bind and Search
func BindAndSearch(l *ldap.Conn) (*ldap.SearchResult, error) {
	l.Bind(BindUsername, BindPassword)

	searchReq := ldap.NewSearchRequest(
		BaseDN,
		ldap.ScopeBaseObject, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		Filter,
		[]string{},
		nil,
	)
	result, err := l.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("Search Error: %s", err)
	}

	if len(result.Entries) > 0 {
		return result, nil
	} else {
		return nil, fmt.Errorf("Couldn't fetch search entries")
	}
}

func main() {
	// Non-TLS Connection
	l, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Anonymous Bind and Search
	result, err := AnonymousBindAndSearch(l)
	if err != nil {
		log.Fatal(err)
	}
	result.Entries[0].Print()

	// TLS Connection
	ltls, err := ConnectTLS()
	if err != nil {
		log.Fatal(err)
	}
	defer ltls.Close()

	// Normal Bind and Search
	result, err = BindAndSearch(ltls)
	if err != nil {
		log.Fatal(err)
	}
	result.Entries[0].Print()
}
