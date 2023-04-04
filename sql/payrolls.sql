DROP TABLE IF EXISTS "public"."payrolls";
CREATE TABLE IF NOT EXISTS "public"."payrolls"(
  "id"                BIGSERIAL NOT NULL,
  "basic_salary"      BIGINT NOT NULL,
  "pay_cut"           BIGINT NOT NULL,
  "additional_salary" BIGINT NOT NULL,
  "employee_id"       BIGSERIAL NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_payrolls_employee" FOREIGN KEY ("employee_id") REFERENCES "public"."employees" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
