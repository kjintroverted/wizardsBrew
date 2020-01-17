SELECT * from races;

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    1,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    'Young dragonborn grow quickly. They walk hours after hatching, attain the size and development of a 10-year-old human child by the age of 3, and reach adulthood by 15. They live to be around 80.',
    'Dragonborn tend to extremes, making a conscious choice for one side or the other in the cosmic war between good and evil (represented by Bahamut and Tiamat, respectively). Most dragonborn are good, but those who side with Tiamat can be terrible villains.',
    'Dragonborn are taller and heavier than humans, standing well over 6 feet tall and averaging almost 250 pounds. Your size is Medium.',
    array[
      ROW('Draconic Ancestry', array['You have draconic ancestry. Choose one type of dragon from the Draconic Ancestry table. Your breath weapon and damage resistance are determined by the dragon type, as shown in the table.'])::trait,
      ROW('Dragon|Damage Type|Breath Weapon', array[
        'Black|Acid|5 by 30 ft. line (Dex. save)',
        'Blue|Lightning|5 by 30 ft. line (Dex. save)',
        'Brass|Fire|5 by 30 ft. line (Dex. save)',
        'Bronze|Lightning|5 by 30 ft. line (Dex. save)',
        'Copper|Acid|5 by 30 ft. line (Dex. save)',
        'Gold|Fire|15 ft. cone (Dex. save)',
        'Green|Poison|15 ft. cone (Con. save)',
        'Red|Fire|15 ft. cone (Dex. save)',
        'Silver|Cold|15 ft. cone (Con. save)',
        'White|Cold|15 ft. cone (Con. save)'
        ])::trait,
      ROW('Breath Weapon', array[
        'You can use your action to exhale destructive energy. Your draconic ancestry determines the size, shape, and damage type of the exhalation.',
        'When you use your breath weapon, each creature in the area of the exhalation must make a saving throw, the type of which is determined by your draconic ancestry. The DC for this saving throw equals 8 + your Constitution modifier + your proficiency bonus. A creature takes 2d6 damage on a failed save, and half as much damage on a successful one. The damage increases to 3d6 at 6th level, 4d6 at 11th level, and 5d6 at 16th level.',
        'After you use your breath weapon, you can''t use it again until you complete a short or long rest.'
        ])::trait,
      ROW('Damage Resistance', array['You have resistance to the damage type associated with your draconic ancestry.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Draconic. Draconic is thought to be one of the oldest languages and is often used in the study of magic. The language sounds harsh to most other creatures and includes numerous hard consonants and sibilants.'])::trait,
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    2,
    'Dwarf (Hill)',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('WIS', 1)::ability_mod
    ],
    'Medium',
    25, 
    'Dwarves mature at the same rate as humans, but they''re considered young until they reach the age of 50. On average, they live about 350 years.',
    'Most dwarves are lawful, believing firmly in the benefits of a well-ordered society. They tend toward good as well, with a strong sense of fair play and a belief that everyone deserves to share in the benefits of a just order.',
    'Dwarves stand between 4 and 5 feet tall and average about 150 pounds. Your size is Medium.',
    array[
      ROW('Speed', array['Your speed is not reduced by wearing heavy armor.'])::trait,
      ROW('Darkvision', array['Accustomed to life underground, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can not discern color in darkness, only shades of gray.'])::trait,
      ROW('Dwarven Resilience', array['You have advantage on saving throws against poison, and you have resistance against poison damage.'])::trait,
      ROW('Dwarven Combat Training', array['You have proficiency with the battleaxe, handaxe, light hammer, and warhammer.'])::trait,
      ROW('Tool Proficiency', array['You gain proficiency with the artisan''s tools of your choice: Smith''s tools, brewer''s supplies, or mason''s tools.'])::trait,
      ROW('Stone Cunning', array['Whenever you make an Intelligence (History) check related to the origin of stonework, you are considered proficient in the History skill and add double your proficiency bonus to the check, instead of your normal proficiency bonus.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Dwarvish. Dwarvish is full of hard consonants and guttural sounds, and those characteristics spill over into whatever other language a dwarf might speak.'])::trait,
      ROW('Dwarven Toughness', array['Your hit point maximum increases by 1, and it increases by 1 every time you gain a level.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    3,
    'Dwarf (Mountain)',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CON', 2)::ability_mod
    ],
    'Medium',
    25, 
    'Dwarves mature at the same rate as humans, but they''re considered young until they reach the age of 50. On average, they live about 350 years.',
    'Most dwarves are lawful, believing firmly in the benefits of a well-ordered society. They tend toward good as well, with a strong sense of fair play and a belief that everyone deserves to share in the benefits of a just order.',
    'Dwarves stand between 4 and 5 feet tall and average about 150 pounds. Your size is Medium.',
    array[
      ROW('Speed', array['Your speed is not reduced by wearing heavy armor.'])::trait,
      ROW('Darkvision', array['Accustomed to life underground, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can''t discern color in darkness, only shades of gray.'])::trait,
      ROW('Dwarven Resilience', array['You have advantage on saving throws against poison, and you have resistance against poison damage.'])::trait,
      ROW('Dwarven Combat Training', array['You have proficiency with the battleaxe, handaxe, light hammer, and warhammer.'])::trait,
      ROW('Tool Proficiency', array['You gain proficiency with the artisan''s tools of your choice: Smith''s tools, brewer''s supplies, or mason''s tools.'])::trait,
      ROW('Stonecunning', array['Whenever you make an Intelligence (History) check related to the origin of stonework, you are considered proficient in the History skill and add double your proficiency bonus to the check, instead of your normal proficiency bonus.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Dwarvish. Dwarvish is full of hard consonants and guttural sounds, and those characteristics spill over into whatever other language a dwarf might speak.'])::trait,
      ROW('Dwarven Armor Training', array['You have proficiency with light and medium armor.'])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    4,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    5,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    6,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    7,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    8,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    9,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    10,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    11,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    12,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    13,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    14,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    15,
    'Dragonborn',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    '',
    '',
    '',
    array[
      ROW('', array[''])::trait
      ]
);