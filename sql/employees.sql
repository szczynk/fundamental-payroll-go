DROP TABLE IF EXISTS "public"."employees";
CREATE TABLE IF NOT EXISTS "public"."employees"(
  "id"      BIGSERIAL NOT NULL, 
  "name"    VARCHAR(255) NOT NULL, 
  "gender"  VARCHAR(255) NOT NULL, 
  "grade"   INT NOT NULL, 
  "married" BOOLEAN NOT NULL, 
  PRIMARY KEY ("id")
);
