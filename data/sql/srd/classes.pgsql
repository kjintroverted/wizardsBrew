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
    array['Rages', '2', '2', '3', '3', '3', '4', '4', '4', '4', '4', '4', '5', '5', '5', '5', '5', '6', '6', '6', null],
    array['Damage Mod', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4']
  ]
);

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
VALUES (
  2,
  'Bard',
  '1d8',
  array['light armor'],
  array['simple', 'hand crossbow', 'longsword', 'rapier', 'shortsword'],
  array['three musical instruments of your choice'],
  array[
    'DEX',
    'CHA'
  ],
  'Choose any 3',
  '8 + your Constitution modifier',
  'You start with the following items, plus anything provided by your background: (a) a rapier, (b) a longsword, or (c) any simple weapon; (a) a diplomat''s pack or (b) an entertainer''s pack; (a) a lute or (b) any other musical instrument; Leather armor, and a dagger',
  array[
    row('', 
      array[
        'Humming as she traces her fingers over an ancient monument in a long-forgotten ruin, a half-elf in rugged leathers finds knowledge springing into her mind, conjured forth by the magic of her song—knowledge of the people who constructed the monument and the mythic saga it depicts.',
        'A stern human warrior bangs his sword rhythmically against his scale mail, setting the tempo for his war chant and exhorting his companions to bravery and heroism. The magic of his song fortifies and emboldens them.',
        'Laughing as she tunes her cittern, a gnome weaves her subtle magic over the assembled nobles, ensuring that her companions'' words will be well received.',
        'Whether scholar, skald, or scoundrel, a bard weaves magic through words and music to inspire allies, demoralize foes, manipulate minds, create illusions, and even heal wounds.'
        ])::section,
    row('Music and Magic', 
      array[
        'In the worlds of D&D, words and music are not just vibrations of air, but vocalizations with power all their own. The bard is a master of song, speech, and the magic they contain. Bards say that the multiverse was spoken into existence, that the words of the gods gave it shape, and that echoes of these primordial Words of Creation still resound throughout the cosmos. The music of bards is an attempt to snatch and harness those echoes, subtly woven into their spells and powers.',
        'The greatest strength of bards is their sheer versatility. Many bards prefer to stick to the sidelines in combat, using their magic to inspire their allies and hinder their foes from a distance. But bards are capable of defending themselves in melee if necessary, using their magic to bolster their swords and armor. Their spells lean toward charms and illusions rather than blatantly destructive spells. They have a wide-ranging knowledge of many subjects and a natural aptitude that lets them do almost anything well. Bards become masters of the talents they set their minds to perfecting, from musical performance to esoteric knowledge.'
        ])::section,
    row('Learning from Experience', 
      array[
        'True bards are not common in the world. Not every minstrel singing in a tavern or jester cavorting in a royal court is a bard. Discovering the magic hidden in music requires hard study and some measure of natural talent that most troubadours and jongleurs lack. It can be hard to spot the difference between these performers and true bards, though. A bard''s life is spent wandering across the land gathering lore, telling stories, and living on the gratitude of audiences, much like any other entertainer. But a depth of knowledge, a level of musical skill, and a touch of magic set bards apart from their fellows.',
        'Only rarely do bards settle in one place for long, and their natural desire to travel—to find new tales to tell, new skills to learn, and new discoveries beyond the horizon—makes an adventuring career a natural calling. Every adventure is an opportunity to learn, practice a variety of skills, enter long-forgotten tombs, discover lost works of magic, decipher old tomes, travel to strange places, or encounter exotic creatures. Bards love to accompany heroes to witness their deeds firsthand. A bard who can tell an awe-inspiring story from personal experience earns renown among other bards. Indeed, after telling so many stories about heroes accomplishing mighty deeds, many bards take these themes to heart and assume heroic roles themselves.'
        ])::section,
    row('Creating a Bard', 
      array[
        'Bards thrive on stories, whether those stories are true or not. Your character''s background and motivations are not as important as the stories that he or she tells about them. Perhaps you had a secure and mundane childhood. There''s no good story to be told about that, so you might paint yourself as an orphan raised by a hag in a dismal swamp. Or your childhood might be worthy of a story. Some bards acquire their magical music through extraordinary means, including the inspiration of fey or other supernatural creatures.',
        'Did you serve an apprenticeship, studying under a master, following the more experienced bard until you were ready to strike out on your own? Or did you attend a college where you studied bardic lore and practiced your musical magic? Perhaps you were a young runaway or orphan, befriended by a wandering bard who became your mentor. Or you might have been a spoiled noble child tutored by a master. Perhaps you stumbled into the clutches of a hag, making a bargain for a musical gift in addition to your life and freedom, but at what cost?',
        'You can make a bard quickly by following these suggestions. First, Charisma should be your highest ability score, followed by Dexterity. Second, choose the entertainer background. Third, choose the dancing lights and vicious mockery cantrips, along with the following 1st-level spells: charm person, detect magic, healing word, and thunderwave.'
        ])::section
  ],
  array[
    array['Cantrips Known', '2', '2', '2', '3', '3', '3', '3', '3', '3', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4'],
    array['Spells Known', '4', '5', '6', '7', '8', '9', '10', '11', '12', '14', '15', '15', '16', '18', '19', '19', '20', '22', '22', '22'],
    array['1st', '2', '3', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4', '4'],
    array['2nd', '0', '0', '2', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3'],
    array['3rd', '0', '0', '0', '0', '2', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3'],
    array['4th', '0', '0', '0', '0', '0', '0', '1', '2', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3', '3'],
    array['5th', '0', '0', '0', '0', '0', '0', '0', '0', '1', '2', '2', '2', '2', '2', '2', '2', '2', '3', '3', '3'],
    array['6th', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '2', '2'],
    array['7th', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '1', '1', '1', '1', '2'],
    array['8th', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '1', '1', '1'],
    array['9th', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '1']
  ]
);

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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

INSERT into classes (id, name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_hp, init_equip, description, progress)
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
