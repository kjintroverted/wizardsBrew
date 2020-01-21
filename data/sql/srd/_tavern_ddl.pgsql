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

DROP TYPE IF EXISTS spell_time CASCADE;
CREATE TYPE spell_time AS (
  time varchar(50),
  condition text
  );

DROP TYPE IF EXISTS ability_mod CASCADE;
CREATE TYPE ability_mod AS (
  name varchar(10),
  mod int
  );

DROP TYPE IF EXISTS spell_comp CASCADE;
CREATE TYPE spell_comp AS (
  name varchar(15),
  description text,
  cost int,
  consume bool
  );

DROP TYPE IF EXISTS size CASCADE;
CREATE TYPE size AS ENUM ('Small', 'Medium', 'Large');

DROP TYPE IF EXISTS damage_type CASCADE;
CREATE TYPE damage_type AS ENUM (
   'acid', 
   'bludgeoning', 
   'cold', 
   'fire', 
   'force', 
   'lightning', 
   'necrotic', 
   'piercing', 
   'poison', 
   'psychic', 
   'radiant', 
   'slashing', 
   'thunder'
   );

DROP TYPE IF EXISTS magic_school CASCADE;
CREATE TYPE magic_school AS ENUM (
   'transmutation', 
   'necromancy', 
   'illusion',
   'evocation',
   'enchantment',
   'divination',
   'conjuration',
   'abjuration'
   );

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
  'Water Vehicle',
  'Air Vehicle',
  'Land Vehicle',
  'Treasure',
  'Trade Good',
  'Tool',
  'Tack and Harness',
  'Staff',
  'Simple',
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
  'Martial',
  'Light Armor',
  'Instrument',
  'Heavy Armor',
  'Generic Variant',
  'Gaming Set',
  'Futuristic',
  'Firearm',
  'Artisan Tools',
  'Adventuring Gear',
  'Poison',
  'Wondrous Item',
  'Wand',
  'Scroll',
  'Rod',
  'Ring',
  'Potion'
  );

DROP TYPE IF EXISTS weapon_info CASCADE;
CREATE TYPE weapon_info AS (
   category item_type,
   damage varchar(10),
   damage_type damage_type
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
   pro_armor varchar(100)[],
   pro_weapon varchar(25)[],
   pro_tool varchar(100),
   pro_save varchar(10)[],
   skills text,
   init_equip text[],
   description section[],
   progress text[]
);

DROP TABLE IF EXISTS backgrounds;
CREATE TABLE backgrounds
(
   id   serial PRIMARY KEY,
   name varchar(50) UNIQUE NOT NULL,
   pro_skill varchar(25)[],
   pro_tool text,
   language varchar(100),
   equipment text,
   special_opts trait,
   character_opts trait[]
);

DROP TABLE IF EXISTS feats;
CREATE TABLE feats
(
   id   serial PRIMARY KEY,
   name varchar(50) UNIQUE NOT NULL,
   ability ability_mod[],
   description section[],
   class varchar(50),
   subclass varchar(50),
   background_req varchar(50),
   level int,
   prereq varchar(50)
);

DROP TABLE IF EXISTS items;
CREATE TABLE items
(
   id   serial PRIMARY KEY,
   name varchar(100) UNIQUE NOT NULL,
   type varchar(50),
   cost numeric,
   weight numeric,
   attune varchar(100),
   rarity item_rarity,
   weapon weapon_info,
   armor_class int,
   info section[]
);

DROP TABLE IF EXISTS spells;
CREATE TABLE spells
(
   id   serial PRIMARY KEY,
   name varchar(50) UNIQUE NOT NULL,
   school varchar(50),
   time spell_time,
   duration varchar(100),
   comp spell_comp[],
   concentrate bool default false,
   range varchar(50),
   level int NOT NULL,
   class varchar(50)[],
   description section[]
);