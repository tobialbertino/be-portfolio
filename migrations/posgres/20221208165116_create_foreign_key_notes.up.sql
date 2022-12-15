-- public.notes foreign keys

ALTER TABLE notes 
ADD CONSTRAINT "fk_notes.owner_users.id" 
FOREIGN KEY ("owner") 
REFERENCES users(id) 
ON DELETE CASCADE;

-- public.collaborations foreign keys
ALTER TABLE collaborations 
ADD CONSTRAINT "fk_collaborations.note_id_notes.id" 
FOREIGN KEY (note_id) 
REFERENCES notes(id) 
ON DELETE CASCADE;

ALTER TABLE collaborations 
ADD CONSTRAINT "fk_collaborations.user_id_users.id" 
FOREIGN KEY (user_id) 
REFERENCES users(id) 
ON DELETE CASCADE;