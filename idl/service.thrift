include "./base.thrift"

/**
 * comment
 */
service TestService {
  void ping(), /* hehehe */
  base.UserInfo1 getUserInfo(),
  list<base.UserInfo> getAllUserInfo(),
} (
  api.uri_prefix = "/api"
)
