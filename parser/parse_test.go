package parser

import "testing"

func TestParseThrift(t *testing.T) {
	err := ParseThrift(
		`
		namespace go face.test
		enum Sex {
			Male = 1,
			Female = 2,
		}
		struct LoginRequest {
			1: required string userName
			2: required string password
		}
		struct UserInfo {
			1: required string openId
			2: optional string userName
			3: optional Sex sex
		}
		service TestService {
			string login(1: LoginRequest req)(api.post="/api/login")
			UserInfo getUserInfo(1: string openId)(api.get="/api/getUserInfo")
		}
		`,
	)
	if err != nil {
		t.Errorf("ParseThrift failed: %v", err)
		return
	}
}

func TestParseThriftFile(t *testing.T) {
	ParseThriftFile("../idl/service.thrift")
}
