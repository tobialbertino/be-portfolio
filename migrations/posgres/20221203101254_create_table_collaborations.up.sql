CREATE TABLE collaborations (
	id varchar(50) NOT NULL,
	note_id varchar(50) NOT NULL,
	user_id varchar(50) NOT NULL,
    
	-- CONSTRAINT collaborations_pkey PRIMARY KEY (id),
	-- CONSTRAINT unique_note_id_and_user_id UNIQUE (note_id, user_id)
	PRIMARY KEY (id),
	UNIQUE (note_id, user_id)
);