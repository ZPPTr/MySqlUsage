# MySqlUsage
- **up containers** `make up`
- **init data** `make init` - creates the table of users and inserts 40 000 000 rows with random data
- **insert data** `make insert` - inserts 10000 rows and measures time execution (ips=500)
- **to insert data with parameters use** `docker-compose run go /dist/main -ips=1000 -count=1000`

**to change** `innodb_flush_log_at_trx_commit` set it in the docker-compose file and run `make restart`


## 40 000 000 rows query
`SELECT * FROM user WHERE true ORDER BY birthday DESC LIMIT 40000000`

|        | NO INDEX | BTREE | HASH  |
|--------|----------|-------|-------|
| Seconds | 30-36    | 30-32 | 30-31 |

## insert 10000 rows
| innodb_flush_log_at_trx_commit | 0        | 1        | 2        |
|--------------------------------|----------|----------|----------|
| IPS 100   (time/per second)    | 1m 40s   | 1m 40s   | 1m 40s   |
| IPS 500   (time/per second)    | 26s/~400 | 42s/~240 | 21s/~460 |
| IPS 1000  (time/per second)    | 24s/~430 | 34s/~330 | 23s/~460 |
| IPS 5000  (time/per second)    | 23s/~450 | 42s/~240 | 23s/~450 |
