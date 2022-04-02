package userrepository

import usersqldb "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/user_sql_db"

type UserRepository struct {
	db usersqldb.UserDber
}
