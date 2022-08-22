package usecases

const CREATE_READERS_TABLE = `CREATE TABLE readers (
	id TEXT PRIMARY KEY,
	created_at TIMESTAMP
  );`

const CREATE_COMICS_TABLE = `CREATE TABLE comics (
	id TEXT PRIMARY KEY,
	name TEXT,
	latest_volume TEXT,
	updated_at TIMESTAMP
  );`

const CREATE_FAVORITE_TABLE = `CREATE TABLE favorite (
	id SERIAL PRIMARY KEY,
	reader_id TEXT,
	comic_id TEXT,

	FOREIGN KEY(reader_id)
	  REFERENCES readers(id),
	FOREIGN KEY(comic_id)
	  REFERENCES comics(id)
  );`

const CREATE_HISTORY_TABLE = `CREATE TABLE history (
	id SERIAL PRIMARY KEY,
	reader_id TEXT NOT NULL,
	comic_id TEXT UNIQUE NOT NULL,
	volume TEXT,
	page INT,
	read_at TIMESTAMP,

	FOREIGN KEY(reader_id)
	  REFERENCES readers(id),
	FOREIGN KEY(comic_id)
	  REFERENCES comics(id),
	CONSTRAINT reader_comic_constraint UNIQUE (reader_id, comic_id)
  );`
