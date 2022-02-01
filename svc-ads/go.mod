module github.com/u-shylianok/ad-service/svc-ads

go 1.17

require (
	github.com/u-shylianok/ad-service/svc-ads/client v0.0.1
	github.com/u-shylianok/ad-service/svc-auth/client v0.0.1
)

replace (
	github.com/u-shylianok/ad-service/svc-ads/client v0.0.1 => ../svc-ads/client
	github.com/u-shylianok/ad-service/svc-auth/client v0.0.1 => ../svc-auth/client
)

require (
	github.com/golang/protobuf v1.5.0
	github.com/jackc/pgx/v4 v4.14.1
	github.com/jmoiron/sqlx v1.3.4
	github.com/maxbrunsfeld/counterfeiter/v6 v6.4.1
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.9.1 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/net v0.0.0-20210331212208-0fccb6fa2b5c // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.1.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
