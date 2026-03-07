CREATE TABLE database_metadata(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	version INTEGER NOT NULL
);

CREATE TABLE server(
	uuid TEXT PRIMARY KEY,
	owner_id INTEGER NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	profile_pic TEXT,
	insert_date INTEGER NOT NULL,
	FOREIGN KEY (owner_id) REFERENCES user(id)
);

CREATE TABLE channel(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	server_uuid TEXT NOT NULL,
	type INTEGER NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	insert_date INTEGER NOT NULL,
	FOREIGN KEY (server_uuid) REFERENCES server(uuid)
)
