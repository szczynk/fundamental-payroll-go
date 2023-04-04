DROP TABLE IF EXISTS "public"."salaries";
CREATE TABLE IF NOT EXISTS "public"."salaries"(
  "id"             BIGSERIAL NOT NULL, 
  "grade"          INT NOT NULL, 
  "basic_salary"   BIGINT NOT NULL,
  "pay_cut"        BIGINT NOT NULL,
  "allowance"      BIGINT NOT NULL,
  "head_of_family" BIGINT NOT NULL,
  PRIMARY KEY ("id")
);

INSERT INTO "salaries"(id,grade,basic_salary,pay_cut,allowance,head_of_family) VALUES (1,1,5000000,100000,150000,1000000);
INSERT INTO "salaries"(id,grade,basic_salary,pay_cut,allowance,head_of_family) VALUES (2,2,9000000,200000,300000,2000000);
INSERT INTO "salaries"(id,grade,basic_salary,pay_cut,allowance,head_of_family) VALUES (3,3,15000000,400000,600000,3000000);