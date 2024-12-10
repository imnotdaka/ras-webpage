package database

var (
	CreateUserQuery = `
	INSERT INTO user(first_name, last_name, email, password) VALUES (?, ?, ?, ?)
	`
	GetUserByEmailQuery = `
	SELECT id, password FROM user WHERE email = ?
	`
	GetUserByIDQuery = `
	SELECT first_name, last_name, email FROM user WHERE id = ?
	`
	UpdateUserByIdQuery = `
	UPDATE user SET first_name=?, last_name=?, email=?, password=? WHERE id=?
	`
	DeleteUserByIDQuery = `
	DELETE FROM user WHERE id=?
	`
	CreatePlanQuery = `
	INSERT INTO plan(id, reason, frequency, frequency_type, transaction_amount) VALUES (?, ?, ?, ?, ?)
	`
	GetPlanQuery = `
	SELECT * FROM plan
	`
	GetPlanByIdQuery = `
	SELECT * FROM plan WHERE id=?
	`
)
