package parser

import "testing"

func TestParseThrift(t *testing.T) {
	ParseThrift(
		`
		namespace go face.test
		struct UserInfo {
			1: required string openId
			2: optional string userName
		}
		service TestService {
			string login(1:string password)(api.post="/api/login")
			UserInfo getUserInfo(1: string openId)(api.get="/api/getUserInfo")
		}
		`,
	)
}

func TestParseThriftFile(t *testing.T) {
	ParseThriftFile("../idl/service.thrift")
}
