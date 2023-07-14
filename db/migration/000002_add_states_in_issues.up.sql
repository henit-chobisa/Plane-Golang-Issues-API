ALTER TABLE "issues"
ADD COLUMN "state_id" uuid,
ADD FOREIGN KEY ("state_id") REFERENCES "states"("id");