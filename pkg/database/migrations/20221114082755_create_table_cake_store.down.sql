DROP TABLE IF EXISTS cake_store;
-- $ migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/privy_cake-store" -path pkg/database/migrations up