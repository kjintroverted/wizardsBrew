DROP TYPE IF EXISTS node_type CASCADE;
CREATE TYPE node_type AS ENUM ('choice', 'terminus');

DROP TYPE IF EXISTS choice CASCADE;
CREATE TYPE choice AS (
   prompt text,
   value int
);

DROP TABLE IF EXISTS story_nodes;
CREATE TABLE story_nodes
(
   id serial PRIMARY KEY,
   type node_type NOT NULL,
   value text NOT NULL,
   paths choice[]
);