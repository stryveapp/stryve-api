DROP INDEX communities_owner_id_idx;
ALTER TABLE communities DROP CONSTRAINT communities_owner_id_fkey;
DROP TABLE IF EXISTS communities;
