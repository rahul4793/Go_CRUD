## Curl for testing

1) For all users
`curl --location --request GET 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--header 'Cookie: Cookie_1=value' \
--data '{
  "page": 1,
  "limit": 10
}
'`
2) For a specific users
`curl --location --request GET 'http://localhost:8080/users/686cc62568bef2601cef514a' \
--header 'Content-Type: application/json' \
--header 'Cookie: Cookie_1=value' \
--data '{
  "page": 1,
  "limit": 2
}
'`
3) For adding a user
`curl --location 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--header 'Cookie: Cookie_1=value' \
--data-raw '{"name":"mohan","email":"mddtescdcdctan@gmail.com","age":19}'
'`
4) For updating user deatils
`
curl --location --request PUT 'http://localhost:8080/users/686bb0b2abc2333f945ec59b' \
--header 'Content-Type: application/json' \
--header 'Cookie: Cookie_1=value' \
--data-raw '{"name":"rohiteeeee","email":"rashussll@example.com","age":30}'`

5) For deleting a user
`
curl --location --request DELETE 'http://localhost:8080/users/686bb0b2abc2333f945ec59b' \
--header 'Cookie: Cookie_1=value'
`
