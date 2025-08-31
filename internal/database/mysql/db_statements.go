package mysql

var (
	create_user = "insert into users(first_name, last_name, email, user_name, password, phone, address, nationality, d_o_b) values(?, ?, ?, ?, ?, ?, ?, ?, ?)"
)
