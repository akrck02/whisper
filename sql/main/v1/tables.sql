CREATE TABLE database_metadata(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	version INTEGER NOT NULL
);

CREATE TABLE user(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid TEXT NOT NULL,
	email TEXT NOT NULL,
	username TEXT NOT NULL,
	profile_pic TEXT,
	password TEXT NOT NULL,
	insert_date INTEGER NOT NULL
);

CREATE INDEX idx_user_uuid ON user(uuid);
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

CREATE TABLE user_server(
	user_id INTEGER NOT NULL,
	server_uuid TEXT NOT NULL,
	PRIMARY KEY (user_id, server_uuid),
	FOREIGN KEY (user_id) REFERENCES user(id)
);
