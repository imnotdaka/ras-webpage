package database

var (
	CreateUserQuery = `
	INSERT INTO user(firstname, lastname, email, password) VALUES (?, ?, ?, ?)
	`
	GetUserByEmailQuery = `
	SELECT id, password FROM user WHERE email = ?
	`
	GetUserByIDQuery = `
	SELECT firstname, lastname, email FROM user WHERE id = ?
	`
	UpdateUserByIdQuery = `
	UPDATE user SET name=?, dni=?, email=?, bday=? WHERE id=?
	`
	DeleteUserByIDQuery = `
	DELETE FROM user WHERE id=?
	`
)
