include "./base.thrift"

struct UserReq {
  1: required string openId
  2: optional string email
}

/**
 * comment
 */
service TestService {
  void ping(), /* hehehe */
  base.UserInfo1 getUserInfo(1: UserReq req),
  list<base.UserInfo> getAllUserInfo(),
} (
  api.uri_prefix = "/api"
)
