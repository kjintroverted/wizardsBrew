select * from spells;

INSERT INTO spells (name, school, time, duration, comp, concentrate, range, level, class, description)
VALUES (
  'Acid Splash',
  'conjuration',
  '1 action',
  'Instantaneous',
  array['v', 's'],
  false,
  '60 feet',
  0, 
  array['Artificer', 'Wizard', 'Sorcerer'],
  array[
    'You hurl a bubble of acid. Choose one creature you can see within range, or choose two creatures you can see within range that are within 5 feet of each other. A target must succeed on a Dexterity saving throw or take 1d6 acid damage.',
    'This spell''s damage increases by 1d6 when you reach 5th level (2d6), 11th level (3d6), and 17th level (4d6).'
  ]
)