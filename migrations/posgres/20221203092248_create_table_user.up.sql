CREATE TABLE users (
	id varchar(50) NOT NULL,
	username varchar(50) NOT NULL,
	"password" text NOT NULL,
	fullname text NOT NULL,
    
	-- CONSTRAINT users_pkey PRIMARY KEY (id),
	-- CONSTRAINT users_username_key UNIQUE (username)
	PRIMARY KEY (id),
	UNIQUE (username)
);