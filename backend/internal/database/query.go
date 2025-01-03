package database

var (
	CreateUserQuery = `
	INSERT INTO user(first_name, last_name, email, password) VALUES (?, ?, ?, ?)
	`
	GetUserByEmailQuery = `
	SELECT id, password FROM user WHERE email = ?
	`
	GetUserByIDQuery = `
	SELECT first_name, last_name, email, CASE WHEN COUNT(s.subscription_id) = 1 THEN true ELSE false END AS subscribed FROM user u JOIN subscription s ON u.id = s.user_id WHERE id=? and status = "authorized" GROUP BY u.id, u.first_name ;
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
	UpdatePlanQuery = `
	UPDATE plan SET reason=?, transaction_amount=? WHERE id=?
	`

	CreateSubscriptionQuery = `
	INSERT INTO subscription(subscription_id, user_id, date_created, next_payment, plan_id, status) VALUES (?, ?, ?, ?, ?, ?)
	`
	GetSubscriptionByUserIDQuery = `
	SELECT reason, status, date_created, next_payment FROM subscription s JOIN plan p ON s.plan_id = p.id WHERE user_id=? and status="authorized";
	`
	UpdateSubscriptionQuery = `
	UPDATE subscription SET next_payment=?, status=?, last_modified=? WHERE subscription_id=?
	`

	CreateSessionQuery = `
	INSERT INTO session(user_id, refresh_token) VALUES (?, ?)
	`
	GetSessionQuery = `
	SELECT refresh_token, is_valid FROM session WHERE refresh_token = ?
	`
	UpdateSessionQuery = `
	UPDATE session SET is_valid=? WHERE refresh_token=?
	`
)
