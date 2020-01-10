DROP TYPE IF EXISTS trait CASCADE;
CREATE TYPE trait AS (
  name varchar(50),
  description text[]
  );

DROP TYPE IF EXISTS section CASCADE;
CREATE TYPE section AS (
  title varchar(50),
  body text[]
  );

DROP TYPE IF EXISTS ability_mod CASCADE;
CREATE TYPE ability_mod AS (
  name varchar(10),
  mod int
  );

DROP TYPE IF EXISTS size CASCADE;
CREATE TYPE size AS ENUM ('Small', 'Medium', 'Large');

DROP TYPE IF EXISTS item_rarity CASCADE;
CREATE TYPE item_rarity AS ENUM (
   'Common',
   'Uncommon',
   'Rare',
   'Very Rare',
   'Legendary',
   'Artifact',
   'Varies',
   'Unknown'
   );

DROP TYPE IF EXISTS item_type CASCADE;
CREATE TYPE item_type AS ENUM (
  'Vehicle',
  'Treasure',
  'Trade Good',
  'Tool',
  'Tack and Harness',
  'Staff',
  'Simple Weapon',
  'Melee Weapon',
  'Spellcasting Focus',
  'Ranged Weapon',
  'Shield',
  'Renaissance',
  'Explosive',
  'Ammunition',
  'Mount',
  'Modern',
  'Medium Armor',
  'Martial Weapon',
  'Light Armor',
  'Instrument',
  'Heavy Armor',
  'Generic Variant',
  'Gaming Set',
  'Futuristic',
  'Firearm',
  'Artisan Tool',
  'Adventuring Gear',
  'Poison',
  'Wondrous Item',
  'Wand',
  'Scroll',
  'Rod',
  'Ring',
  'Potion'
  );

DROP TABLE IF EXISTS races;
CREATE TABLE races
(
   id serial PRIMARY KEY,
   name varchar(50) UNIQUE NOT NULL,
   ability ability_mod[],
   size size NOT NULL DEFAULT 'Medium',
   speed int NOT NULL DEFAULT 30,
   age text,
   align text,
   size_desc text,
   traits trait[]
);

DROP TABLE IF EXISTS classes;
CREATE TABLE classes
(
   id serial PRIMARY KEY,
   name varchar(50) UNIQUE NOT NULL,
   hit_dice varchar(25) NOT NULL,
   pro_armor varchar(25)[],
   pro_weapon varchar(25)[],
   pro_tool varchar(25)[],
   pro_save varchar(25)[],
   skills text,
   init_hp varchar(50),
   init_equip text,
   description section[],
   progress text[][]
);


DROP TABLE IF EXISTS feats;
CREATE TABLE feats
(
   id   serial PRIMARY KEY,
   name varchar(50) UNIQUE NOT NULL,
   class_id int,
   ability ability_mod[],
   description text[],
   level int,
   path varchar(50),
   prereq varchar(50)
);

DROP TABLE IF EXISTS items;
CREATE TABLE items
(
   id   serial PRIMARY KEY,
   name varchar(50) UNIQUE NOT NULL,
   type item_type NOT NULL,
   cost numeric,
   weight numeric,
   attune boolean default false,
   rarity item_type,
   description text
);

DROP TABLE IF EXISTS spells;
CREATE TABLE spells
(
   id   serial PRIMARY KEY,
   name varchar(50) UNIQUE NOT NULL,
   level int NOT NULL,
   time varchar(50),
   duration varchar(100),
   comp text[],
   school varchar(50),
   concentrate bool default false,
   range varchar(50),
   class_id int[],
   description text[]
);