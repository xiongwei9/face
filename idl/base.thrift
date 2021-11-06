// Enum
enum Sex {
  female = 0
  male = 1
}

struct UserInfo {
  1: required string name
  2: required Sex sex
  3: optional i64 age
}
