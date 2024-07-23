goctl api go -api *.api -dir ./  --style=goZero

go mod tidy

go run user.go -f etc/user.yaml

goctl model mysql ddl -src todo.sql -dir . -c