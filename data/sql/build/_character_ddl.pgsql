DROP TYPE IF EXISTS raw_stat CASCADE;
CREATE TYPE raw_stat AS (
   str int,
   dex int,
   con int,
   intel int,
   wis int,
   cha int
);

DROP TYPE IF EXISTS skill CASCADE;
CREATE TYPE skill AS (
   name varchar(100),
   mult int
);

DROP TABLE IF EXISTS characters;
CREATE TABLE characters
(
   id varchar(100) PRIMARY KEY,
   name varchar(150) NOT NULL,
   photo_url text,
   owner varchar(100) NOT NULL,
   auth_users varchar(100)[],
   auth_req varchar(100)[],
   party_inv varchar(100)[],
   race_id int NOT NULL REFERENCES races,
   class_id int NOT NULL REFERENCES classes,
   subclass varchar(100),
   background varchar(100),
   stats raw_stat NOT NULL,
   xp int NOT NULL,
   max_hp int NOT NULL,
   hp int NOT NULL,
   init int NOT NULL,
   pro_skill skill[],
   pro_tool varchar(100)[],
   pro_weapon varchar(100)[],
   lang varchar(100)[],
   equip_ids int[],
   weapon_ids int[],
   inventory_ids int[],
   gold numeric,
   spell_ids int[],
   feat_ids int[]
);

DROP TABLE IF EXISTS parties;
CREATE TABLE parties
(
   id varchar(100) PRIMARY KEY,
   name varchar(150) NOT NULL,
   photo_url text,
   admin varchar(100),
   members varchar(100)[]
)