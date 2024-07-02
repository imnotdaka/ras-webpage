package database

var (
	CreateUserQuery = `
	INSERT INTO user(name, dni, bday) VALUES (?, ?, ?)
	`
	GetUserByIDQuery = `
	SELECT * FROM user WHERE id = ?
	`
	UpdateUserByIdQuery = `
	UPDATE user SET name=?, dni=?, bday=? WHERE id=?
	`
	DeleteUserByIDQuery = `
	DELETE FROM user WHERE id=?
	`
)
