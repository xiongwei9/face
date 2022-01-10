namespace go face.test

enum Sex {
	Male = 1,
	Female = 2,
}

struct LoginRequest {
	1: required string userName
	2: required string password
	3: required i32 age
}

struct UserInfo {
	1: required string openId
	2: optional string userName
	3: optional Sex sex
	4: optional list<string> favorite
	5: optional map<string, string> fri
}

service TestService {
	string login(1: LoginRequest req)(api.post="/api/login")
	UserInfo getUserInfo(1: string openId)(api.get="/api/getUserInfo")
}
