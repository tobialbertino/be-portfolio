-- public.notes foreign keys
ALTER TABLE IF EXISTS notes 
    DROP CONSTRAINT IF EXISTS "fk_notes.owner_users.id";

-- public.collaborations foreign keys
ALTER TABLE IF EXISTS collaborations 
    DROP CONSTRAINT IF EXISTS "fk_collaborations.note_id_notes.id";
ALTER TABLE IF EXISTS collaborations 
    DROP CONSTRAINT IF EXISTS "fk_collaborations.user_id_users.id";
