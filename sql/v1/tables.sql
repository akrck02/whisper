CREATE TABLE database_metadata(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	version INTEGER NOT NULL
);

CREATE TABLE user(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL,
	username TEXT NOT NULL,
	profile_pic TEXT,
	password TEXT NOT NULL,
	insert_date INTEGER NOT NULL
);

CREATE INDEX idx_user_email ON user(email);
CREATE INDEX idx_user_username ON user(username);

CREATE TABLE device(
	user_id INTEGER NOT NULL,
	address TEXT NOT NULL,
	user_agent TEXT NOT NULL,
	token TEXT NOT NULL,
	insert_date INTEGER NOT NULL,
	update_date INTEGER NOT NULL,
	PRIMARY KEY (user_id, address, user_agent),
 	FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE server(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	owner_id INTEGER NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	profile_pic TEXT,
	insert_date INTEGER NOT NULL,
	FOREIGN KEY (owner_id) REFERENCES user(id)
);

CREATE TABLE channel(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	server_id INTEGER NOT NULL,
	type INTEGER NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	insert_date INTEGER NOT NULL,
	FOREIGN KEY (server_id) REFERENCES server(id)
)
