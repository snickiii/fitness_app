-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS dailyration_id_seq;

-- Table Definition
CREATE TABLE "public"."daily_ration" (
    "id" int4 NOT NULL DEFAULT nextval('dailyration_id_seq'::regclass),
    "user_id" int4 NOT NULL,
    "ration" varchar,
    "created_at" varchar(120),
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS untitled_table_193_id_seq;

-- Table Definition
CREATE TABLE "public"."equipment" (
    "id" int4 NOT NULL DEFAULT nextval('untitled_table_193_id_seq'::regclass),
    "name" varchar
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS exercises_id_seq;

-- Table Definition
CREATE TABLE "public"."exercises" (
    "id" int4 NOT NULL DEFAULT nextval('exercises_id_seq'::regclass),
    "name" varchar,
    "description" varchar,
    "upper_strength_impact" float4,
    "lower_strength_impact" float4,
    "flexibility_impact" float4,
    "endurance_impact" float4,
    "equipment_id" int4,
    "kcal" int4
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS usercharacteristics_id_seq;

-- Table Definition
CREATE TABLE "public"."user_characteristics" (
    "id" int4 NOT NULL DEFAULT nextval('usercharacteristics_id_seq'::regclass),
    "user_id" int4 NOT NULL,
    "upper_strength" float4,
    "lower_strength" float4,
    "flexibility" float4,
    "endurance" float4,
    "height" float4,
    "weight" float4,
    "imt" float4,
    "created_at" varchar(120),
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "username" varchar,
    "password" varchar,
    "fisrtname" varchar,
    "lastname" varchar,
    "email" varchar
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS plan_exercises_id_seq;

-- Table Definition
CREATE TABLE "public"."workout_exercises" (
    "id" int4 NOT NULL DEFAULT nextval('plan_exercises_id_seq'::regclass),
    "user_id" int4,
    "exercise_id" int4,
    "date" varchar,
    "status" varchar
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS workouts_id_seq;

-- Table Definition
CREATE TABLE "public"."workouts" (
    "id" int4 NOT NULL DEFAULT nextval('workouts_id_seq'::regclass),
    "user_id" int4,
    "name" int4,
    "status" varchar
);

