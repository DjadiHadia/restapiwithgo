CREATE TABLE "Agency" (
  "id" integer PRIMARY KEY,
  "Name" varchar,
  "Phone" varchar,
  "Email" varchar
);

CREATE TABLE "Car" (
  "id" integer PRIMARY KEY,
  "Registration_number" varchar,
  "Type" varchar,
  "Brand" varchar,
  "Color" varchar,
  "Places" integer,
  "Daily_rate" money,
  "Fuel" varchar,
  "Chassis_number" varchar,
  "Year" timestamp,
  "Actual_Km" integer,
  "Air_Conditioner" boolean,
  "Technical_inspection_start_date" timestamp,
  "Technical_inspection_end_date" timestamp,
  "insurance_start_date" timestamp,
  "Observation" varchar,
  "insurance_end_date" timestamp,
  "Agency" integer
);

CREATE TABLE "Employee" (
  "id" integer PRIMARY KEY,
  "type" varchar,
  "salary" varchar,
  "Agency" integer,
  "person" integer

);

CREATE TABLE "Client" (
  "id" integer PRIMARY KEY,
  "Document" integer,
  "person" integer
);

CREATE TABLE "Agency_Document" (
  "id" integer PRIMARY KEY,
  "Agency" integer,
  "Document" integer
);

CREATE TABLE "Employee_Document" (
  "id" integer PRIMARY KEY,
  "Agency" integer,
  "Document" integer
);

CREATE TABLE "Client_Document" (
  "id" integer PRIMARY KEY,
  "Agency" integer,
  "Document" integer
);

CREATE TABLE "Document" (
  "id" integer PRIMARY KEY,
  "label" varchar,
  "type" varchar,
  "size" integer
);

CREATE TABLE "Reservation" (
  "id" integer PRIMARY KEY,
  "date" date,
  "duration" integer,
  "start_date" date,
  "end_date" date,
  "car" integer
);

CREATE TABLE "Contract" (
  "id" integer PRIMARY KEY,
  "date" date,
  "duration" integer,
  "start_date" date,
  "end_date" date,
  "amount" money,
  "Observation" varchar,
  "Reservation" integer,
  "Bill" integer
);

CREATE TABLE "Bill" (
  "id" integer PRIMARY KEY,
  "date" date,
  "amount" money,
  "Observation" varchar
);

CREATE TABLE "person" (
  "id" integer PRIMARY KEY,
  "fullname" varchar,
  "adress" varchar,
  "email" varchar,
  "Phone" varchar
);

CREATE TABLE "user_account" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "password" varchar,
  "Employee" integer
);

COMMENT ON COLUMN "Car"."Agency" IS 'ON DELETE NO ACTION';

COMMENT ON COLUMN "Employee"."Agency" IS 'ON DELETE NO ACTION ';

ALTER TABLE "Car" ADD FOREIGN KEY ("Agency") REFERENCES "Agency" ("id") ON DELETE NO ACTION;

ALTER TABLE "Employee" ADD FOREIGN KEY ("Agency") REFERENCES "Agency" ("id");

ALTER TABLE "user_account" ADD FOREIGN KEY ("Employee") REFERENCES "Employee" ("id");

ALTER TABLE "Reservation" ADD FOREIGN KEY ("car") REFERENCES "Car" ("id");

ALTER TABLE "Agency_Document" ADD FOREIGN KEY ("Agency") REFERENCES "Agency" ("id");

ALTER TABLE "Employee_Document" ADD FOREIGN KEY ("Agency") REFERENCES "Agency" ("id");

ALTER TABLE "Client_Document" ADD FOREIGN KEY ("Agency") REFERENCES "Agency" ("id");

ALTER TABLE "Agency_Document" ADD FOREIGN KEY ("Document") REFERENCES "Document" ("id");

ALTER TABLE "Employee_Document" ADD FOREIGN KEY ("Document") REFERENCES "Document" ("id");

ALTER TABLE "Client_Document" ADD FOREIGN KEY ("Document") REFERENCES "Document" ("id");

ALTER TABLE "Contract" ADD FOREIGN KEY ("Reservation") REFERENCES "Reservation" ("id");

ALTER TABLE "Contract" ADD FOREIGN KEY ("Bill") REFERENCES "Bill" ("id");

ALTER TABLE "Client" ADD FOREIGN KEY ("Document") REFERENCES "Client_Document" ("id");

ALTER TABLE "Client" ADD FOREIGN KEY ("person") REFERENCES "person" ("id");

