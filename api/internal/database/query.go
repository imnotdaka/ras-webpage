package database

var (
	CreateUserQuery = `
	INSERT INTO usersdb(name, dni, bday) VALUES (?, ?, ?)
	`
	GetUserByIDQuery = `
	SELECT * FROM usersdb WHERE id = ?
	`
	UpdateUserByIdQuery = `
	UPDATE usersdb SET name=?, dni=?, bday=? WHERE id=?
	`
	DeleteUserByIDQuery = `
	DELETE FROM usersdb WHERE id=?
	`
)
