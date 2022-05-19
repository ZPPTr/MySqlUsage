# MySqlUsage
## run app `make up`

### 40 000 000 items query
`SELECT * FROM user WHERE true ORDER BY birthday DESC LIMIT 40000000`

|        | NO INDEX | BTREE | HASH  |
|--------|----------|-------|-------|
| Seconds | 30-36    | 30-32 | 30-31 |