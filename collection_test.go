package introspector_test

import (
	"fmt"
	"testing"

	"github.com/codeclysm/introspector"
)

func TestCollection(t *testing.T) {
	cases := []struct {
		Token            string
		ExpIntrospection introspector.Introspection
	}{{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJleGFtcGxlIiwiZXhwIjoxNDg2Mzk0Nzc5LCJqdGkiOiJmODVjNTM0Yy03M2JhLTQ3NjMtYTU4MS0yMzkxN2I5Nzc5MjUiLCJpYXQiOjE0ODYzODg3NzksImlzcyI6ImFwaS5leGFtcGxlLmNjIiwibmJmIjoxNDg2Mzg4Nzc5LCJzdWIiOiJ0ZXN0IiwidXNlciI6eyJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJpZCI6InRlc3QiLCJ1aWQiOiJ0ZXN0In19.mpRwH7Klc2P1X93N1f0Qf_W3RcNfxm97xwSLEpgSlIw", introspector.Introspection{Active: false, Subject: "test", ExpiresAt: 1486394779, IssuedAt: 1486388779, NotBefore: 1486388779, Issuer: "api.example.cc", Audience: "example", Extra: map[string]interface{}{"user": map[string]interface{}{"email": "test@example.com", "id": "test", "uid": "test"}, "jti": "f85c534c-73ba-4763-a581-23917b977925"}}}, {"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJleGFtcGxlIiwiZXhwIjoxNDg2Mzk0Nzc5LCJqdGkiOiJmODVjNTM0Yy03M2JhLTQ3NjMtYTU4MS0yMzkxN2I5Nzc5MjUiLCJpYXQiOjE0ODYzODg3NzksImlzcyI6ImFwaS5leGFtcGxlLmNjIiwibmJmIjoxNDg2Mzg4Nzc5LCJzdWIiOiJ0ZXN0IiwidXNlciI6eyJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJpZCI6InRlc3QiLCJ1aWQiOiJ0ZXN0In19.8vc9IRexqGZ4UtyQatgsCpe020Rq7Jpz4FJNjIK114Y", introspector.Introspection{Active: false, Subject: "test", ExpiresAt: 1486394779, IssuedAt: 1486388779, NotBefore: 1486388779, Issuer: "api.example.cc", Audience: "example", Extra: map[string]interface{}{"user": map[string]interface{}{"email": "test@example.com", "id": "test", "uid": "test"}, "jti": "f85c534c-73ba-4763-a581-23917b977925"}}}}
	for _, tc := range cases {
		jwt1 := introspector.JWT{
			Key: []byte("secret"),
		}
		jwt2 := introspector.JWT{
			Key: []byte("terces"),
		}

		coll := introspector.Collection{jwt1, jwt2}

		t.Run(fmt.Sprintf("%s", tc.Token), func(t *testing.T) {
			i, err := coll.Introspect(tc.Token)
			if err != nil {
				t.Fatalf("err: %s", err.Error())
			}
			equal(t, i, &tc.ExpIntrospection)
		})
	}
}
