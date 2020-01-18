select * from classes;

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  1,
  'Barbarian',
  '1d12',
  array[
    'light armor',
    'medium armor',
    'shields'
  ],
  array[
    'simple',
    'martial'
  ],
  null,
  array[
    'STR',
    'CON'
  ],
  'Choose 2 from Animal Handling, Athletics, Intimidation, Nature, Perception, and Survival.',
  '12 + your Constitution modifier',
  'You start with the following items, plus anything provided by your background: (a) a greataxe or (b) any martial melee weapon, (a) two handaxes or (b) any simple weapon, An explorer''s pack, and four javelins.',
  array[
    row('Insight', 
      array[
        'A tall human tribesman strides through a blizzard, draped in fur and hefting his axe. He laughs as he charges toward the frost giant who dared poach his people''s elk herd.',
        'A half-orc snarls at the latest challenger to her authority over their savage tribe, ready to break his neck with her bare hands as she did to the last six rivals.',
        'Frothing at the mouth, a dwarf slams his helmet into the face of his drow foe, then turns to drive his armored elbow into the gut of another.',
        'These barbarians, different as they might be, are defined by their rage: unbridled, unquenchable, and unthinking fury. More than a mere emotion, their anger is the ferocity of a cornered predator, the unrelenting assault of a storm, the churning turmoil of the sea.',
        'For some, their rage springs from a communion with fierce animal spirits. Others draw from a roiling reservoir of anger at a world full of pain. For every barbarian, rage is a power that fuels not just a battle frenzy but also uncanny reflexes, resilience, and feats of strength.'
        ])::section,
    row('Primal Instinct', 
      array[
        'People of towns and cities take pride in how their civilized ways set them apart from animals, as if denying one''s own nature was a mark of superiority. To a barbarian, though, civilization is no virtue, but a sign of weakness. The strong embrace their animal nature—keen instincts, primal physicality, and ferocious rage. Barbarians are uncomfortable when hedged in by walls and crowds. They thrive in the wilds of their homelands: the tundra, jungle, or grasslands where their tribes live and hunt.',
        'Barbarians come alive in the chaos of combat. They can enter a berserk state where rage takes over, giving them superhuman strength and resilience. A barbarian can draw on this reservoir of fury only a few times without resting, but those few rages are usually sufficient to defeat whatever threats arise.'
        ])::section,
    row('A Life of Danger', 
      array[
        'Not every member of the tribes deemed "barbarians" by scions of civilized society has the barbarian class. A true barbarian among these people is as uncommon as a skilled fighter in a town, and he or she plays a similar role as a protector of the people and a leader in times of war. Life in the wild places of the world is fraught with peril: rival tribes, deadly weather, and terrifying monsters. Barbarians charge headlong into that danger so that their people don''t have to.',
        'Their courage in the face of danger makes barbarians perfectly suited for adventuring. Wandering is often a way of life for their native tribes, and the rootless life of the adventurer is little hardship for a barbarian. Some barbarians miss the close-knit family structures of the tribe, but eventually find them replaced by the bonds formed among the members of their adventuring parties.'
        ])::section,
    row('Creating a Barbarian', 
      array[
        'When creating a barbarian character, think about where your character comes from and his or her place in the world. Talk with your DM about an appropriate origin for your barbarian. Did you come from a distant land, making you a stranger in the area of the campaign? Or is the campaign set in a rough-and-tumble frontier where barbarians are common?',
        'What led you to take up the adventuring life? Were you lured to settled lands by the promise of riches? Did you join forces with soldiers of those lands to face a shared threat? Did monsters or an invading horde drive you out of your homeland, making you a rootless refugee? Perhaps you were a prisoner of war, brought in chains to "civilized" lands and only now able to win your freedom. Or you might have been cast out from your people because of a crime you committed, a taboo you violated, or a coup that removed you from a position of authority.',
        'You can make a barbarian quickly by following these suggestions. First, put your highest ability score in Strength, followed by Constitution. Second, choose the outlander background.'
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  2,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  3,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  4,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  5,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  6,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  7,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  8,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  9,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  10,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
 11,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  12,
  '',
  '1d12',
  array[''],
  array[''],
  array[''],
  array[
    'STR',
    'CON'
  ],
  '',
  '',
  '',
  array[
    row('', 
      array[
        '',
        ])::section
  ],
  array[
    array['rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['damage_mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);
