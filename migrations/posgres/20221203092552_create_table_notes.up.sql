CREATE TABLE notes (
	id VARCHAR(50) NOT NULL,
	title TEXT NOT NULL,
	body TEXT NOT NULL,
	tags TEXT[] NOT NULL,
	created_at INT8 NOT NULL,
	updated_at INT8 NOT NULL,
	"owner" VARCHAR(50) NULL,
    
	-- CONSTRAINT notes_pkey PRIMARY KEY (id)
	PRIMARY KEY (id)
);