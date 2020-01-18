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
      ROW('Languages', array['You can speak, read, and write Common and Draconic. Draconic is thought to be one of the oldest languages and is often used in the study of magic. The language sounds harsh to most other creatures and includes numerous hard consonants and sibilants.'])::trait
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

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    4,
    'Elf (Drow)',
    array[
    ROW('DEX', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    'Although elves reach physical maturity at about the same age as humans, the elven understanding of adulthood goes beyond physical growth to encompass worldly experience. An elf typically claims adulthood and an adult name around the age of 100 and can live to be 750 years old.',
    'Elves love freedom, variety, and self-expression, so they lean strongly toward the gentler aspects of chaos. They value and protect others'' freedom as well as their own, and they are more often good than not.',
    'Elves range from under 5 to over 6 feet tall and have slender builds. Your size is Medium.',
    array[
      ROW('Superior Darkvision', array['Accustomed to the depths of the Underdark, you have superior vision in dark and dim conditions. You can see in dim light within 120 feet of you as if it were bright light, and in darkness as if it were dim light. You can''t discern color in darkness, only shades of gray.'])::trait,
      ROW('Keen Senses', array['You have proficiency in the Perception skill.'])::trait,
      ROW('Fey Ancestry', array['You have advantage on saving throws against being charmed, and magic can''t put you to sleep.'])::trait,
      ROW('Trance', array[
        'Elves don''t need to sleep. Instead, they meditate deeply, remaining semiconscious, for 4 hours a day. (The Common word for such meditation is "trance.") While meditating, you can dream after a fashion; such dreams are actually mental exercises that have become reflexive through years of practice. After resting in this way, you gain the same benefit that a human does from 8 hours of sleep.',
        'If you meditate during a long rest, you finish the rest after only 4 hours. You otherwise obey all the rules for a long rest; only the duration is changed.'
        ])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Elvish. Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their language so they can add Elvish ballads to their repertoires.'])::trait,
      ROW('Sunlight Sensitivity', array['You have disadvantage on attack rolls and on Wisdom (Perception) checks that rely on sight when you, the target of your attack, or whatever you are trying to perceive is in direct sunlight.'])::trait,
      ROW('Drow Magic', array['You know the dancing lights cantrip. When you reach 3rd level, you can cast the faerie fire spell once per day; you must finish a long rest in order to cast the spell again using this trait. When you reach 5th level, you can also cast the darkness spell once per day; you must finish a long rest in order to cast the spell again using this trait. Charisma is your spellcasting ability for these spells.'])::trait,
      ROW('Drow Weapon Training', array['You have proficiency with rapiers, shortswords, and hand crossbows.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    5,
    'Elf (High)',
    array[
    ROW('DEX', 2)::ability_mod, 
    ROW('INT', 1)::ability_mod
    ],
    'Medium',
    30, 
    'Although elves reach physical maturity at about the same age as humans, the elven understanding of adulthood goes beyond physical growth to encompass worldly experience. An elf typically claims adulthood and an adult name around the age of 100 and can live to be 750 years old.',
    'Elves love freedom, variety, and self-expression, so they lean strongly toward the gentler aspects of chaos. They value and protect others'' freedom as well as their own, and they are more often good than not.',
    'Elves range from under 5 to over 6 feet tall and have slender builds. Your size is Medium.',
    array[
      ROW('Darkvision', array['Accustomed to twilit forests and the night sky, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can''t discern color in darkness, only shades of gray.'])::trait,
      ROW('Keen Senses', array['You have proficiency in the Perception skill.'])::trait,
      ROW('Trance', array[
        'Elves don''t need to sleep. Instead, they meditate deeply, remaining semiconscious, for 4 hours a day. (The Common word for such meditation is "trance.") While meditating, you can dream after a fashion; such dreams are actually mental exercises that have become reflexive through years of practice. After resting in this way, you gain the same benefit that a human does from 8 hours of sleep.',
        'If you meditate during a long rest, you finish the rest after only 4 hours. You otherwise obey all the rules for a long rest; only the duration is changed.'
        ])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Elvish. Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their language so they can add Elvish ballads to their repertoires.'])::trait,
      ROW('Elf Weapon Training', array['You have proficiency with the longsword, shortsword, shortbow, and longbow.'])::trait,
      ROW('Cantrip', array['You know one cantrip of your choice from the wizard spell list. Intelligence is your spellcasting ability for it.'])::trait,
      ROW('Extra Language', array['You can speak, read, and write one extra language of your choosing.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    6,
    'Elf (Wood)',
    array[
    ROW('DEX', 2)::ability_mod, 
    ROW('WIS', 1)::ability_mod
    ],
    'Medium',
    35, 
    'Although elves reach physical maturity at about the same age as humans, the elven understanding of adulthood goes beyond physical growth to encompass worldly experience. An elf typically claims adulthood and an adult name around the age of 100 and can live to be 750 years old.',
    'Elves love freedom, variety, and self-expression, so they lean strongly toward the gentler aspects of chaos. They value and protect others'' freedom as well as their own, and they are more often good than not.',
    'Elves range from under 5 to over 6 feet tall and have slender builds. Your size is Medium.',
    array[
      ROW('Darkvision', array['Accustomed to twilit forests and the night sky, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can''t discern color in darkness, only shades of gray.'])::trait,
      ROW('Keen Senses', array['You have proficiency in the Perception skill.'])::trait,
      ROW('Trance', array[
        'Elves don''t need to sleep. Instead, they meditate deeply, remaining semiconscious, for 4 hours a day. (The Common word for such meditation is "trance.") While meditating, you can dream after a fashion; such dreams are actually mental exercises that have become reflexive through years of practice. After resting in this way, you gain the same benefit that a human does from 8 hours of sleep.',
        'If you meditate during a long rest, you finish the rest after only 4 hours. You otherwise obey all the rules for a long rest; only the duration is changed.'
        ])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Elvish. Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their language so they can add Elvish ballads to their repertoires.'])::trait,
      ROW('Elf Weapon Training', array['You have proficiency with the longsword, shortsword, shortbow, and longbow.'])::trait,
      ROW('Fleet of Foot', array['Your base walking speed increases to 35 feet.'])::trait,
      ROW('Mask of the Wild', array['You can attempt to hide even when you are only lightly obscured by foliage, heavy rain, falling snow, mist, and other natural phenomena.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    7,
    'Gnome (Forest)',
    array[
    ROW('INT', 2)::ability_mod, 
    ROW('DEX', 1)::ability_mod
    ],
    'Small',
    25, 
    'Gnomes mature at the same rate humans do, and most are expected to settle down into an adult life by around age 40. They can live 350 to almost 500 years.',
    'Gnomes are most often good. Those who tend toward law are sages, engineers, researchers, scholars, investigators, or inventors. Those who tend toward chaos are minstrels, tricksters, wanderers, or fanciful jewelers. Gnomes are good-hearted, and even the tricksters among them are more playful than vicious.',
    'Gnomes are between 3 and 4 feet tall and average about 40 pounds. Your size is Small.',
    array[
      ROW('Darkvision', array['Accustomed to life underground, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can''t discern color in darkness, only shades of gray.'])::trait,
      ROW('Gnome Cunning', array['You have advantage on all Intelligence, Wisdom, and Charisma saving throws against magic.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Gnomish. The Gnomish language, which uses the Dwarvish script, is renowned for its technical treatises and its catalogs of knowledge about the natural world.'])::trait,
      ROW('Natural Illusionist', array['You know the minor illusion cantrip. Intelligence is your spellcasting ability for it.'])::trait,
      ROW('Speak with Small Beasts', array['Through sounds and gestures, you can communicate simple ideas with Small or smaller beasts. Forest gnomes love animals and often keep squirrels, badgers, rabbits, moles, woodpeckers, and other creatures as beloved pets.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    8,
    'Gnome (Rock)',
    array[
    ROW('INT', 2)::ability_mod, 
    ROW('CON', 1)::ability_mod
    ],
    'Small',
    25, 
    'Gnomes mature at the same rate humans do, and most are expected to settle down into an adult life by around age 40. They can live 350 to almost 500 years.',
    'Gnomes are most often good. Those who tend toward law are sages, engineers, researchers, scholars, investigators, or inventors. Those who tend toward chaos are minstrels, tricksters, wanderers, or fanciful jewelers. Gnomes are good-hearted, and even the tricksters among them are more playful than vicious.',
    'Gnomes are between 3 and 4 feet tall and average about 40 pounds. Your size is Small.',
    array[
      ROW('Darkvision', array['Accustomed to life underground, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can''t discern color in darkness, only shades of gray.'])::trait,
      ROW('Gnome Cunning', array['You have advantage on all Intelligence, Wisdom, and Charisma saving throws against magic.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Gnomish. The Gnomish language, which uses the Dwarvish script, is renowned for its technical treatises and its catalogs of knowledge about the natural world.'])::trait,
      ROW('Artificer''s Lore', array['Whenever you make an Intelligence (History) check related to magic items, alchemical objects, or technological devices, you can add twice your proficiency bonus, instead of any proficiency bonus you normally apply.'])::trait,
      ROW('Tinker', array[
        'You have proficiency with artisan''s tools (tinker''s tools). Using those tools, you can spend 1 hour and 10 gp worth of materials to construct a Tiny clockwork device (AC 5, 1 hp). The device ceases to function after 24 hours (unless you spend 1 hour repairing it to keep the device functioning), or when you use your action to dismantle it; at that time, you can reclaim the materials used to create it. You can have up to three such devices active at a time.',
        'When you create a device, choose one of the following options:'
        ])::trait,
      ROW('choices', array[
        'Clockwork Toy: This toy is a clockwork animal, monster, or person, such as a frog, mouse, bird, dragon, or soldier. When placed on the ground, the toy moves 5 feet across the ground on each of your turns in a random direction. It makes noises as appropriate to the creature it represents.',
        'Fire Starter: The device produces a miniature flame, which you can use to light a candle, torch, or campfire. Using the device requires your action.',
        'Music Box: When opened, this music box plays a single song at a moderate volume. The box stops playing when it reaches the song''s end or when it is closed.'
        ])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    9,
    'Half-Elf',
    array[
    ROW('CHA', 2)::ability_mod, 
    ROW('UNI', 2)::ability_mod
    ],
    'Medium',
    30, 
    'Half-elves mature at the same rate humans do and reach adulthood around the age of 20. They live much longer than humans, however, often exceeding 180 years.',
    'Half-elves share the chaotic bent of their elven heritage. They value both personal freedom and creative expression, demonstrating neither love of leaders nor desire for followers. They chafe at rules, resent others'' demands, and sometimes prove unreliable, or at least unpredictable.',
    'Half-elves are about the same size as humans, ranging from 5 to 6 feet tall. Your size is Medium.',
    array[
      ROW('Darkvision', array['Thanks to your elf blood, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can''t discern color in darkness, only shades of gray.'])::trait,
      ROW('Fey Ancestry', array['You have advantage on saving throws against being charmed, and magic can''t put you to sleep.'])::trait,
      ROW('Skill Versatility', array['You gain proficiency in two skills of your choice.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common, Elvish, and one extra language of your choice.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    10,
    'Half-Orc',
    array[
    ROW('STR', 2)::ability_mod, 
    ROW('CON', 1)::ability_mod
    ],
    'Medium',
    30, 
    'Half-orcs mature a little faster than humans, reaching adulthood around age 14. They age noticeably faster and rarely live longer than 75 years.',
    'Half-orcs inherit a tendency toward chaos from their orc parents and are not strongly inclined toward good. Half-orcs raised among orcs and willing to live out their lives among them are usually evil.',
    'Half-orcs are somewhat larger and bulkier than humans, and they range from 5 to well over 6 feet tall. Your size is Medium.',
    array[
      ROW('Darkvision', array['Thanks to your orc blood, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can''t discern color in darkness, only shades of gray.'])::trait,
      ROW('Menacing', array['You gain proficiency in the Intimidation skill.'])::trait,
      ROW('Relentless Endurance', array['When you are reduced to 0 hit points but not killed outright, you can drop to 1 hit point instead. You can''t use this feature again until you finish a long rest.'])::trait,
      ROW('Savage Attacks', array['When you score a critical hit with a melee weapon attack, you can roll one of the weapon''s damage dice one additional time and add it to the extra damage of the critical hit.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Orc. Orc is a harsh, grating language with hard consonants. It has no script of its own but is written in the Dwarvish script.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    11,
    'Halfling (Lightfoot)',
    array[
    ROW('DEX', 2)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Small',
    25, 
    'A halfling reaches adulthood at the age of 20 and generally lives into the middle of his or her second century.',
    'Most halflings are lawful good. As a rule, they are good-hearted and kind, hate to see others in pain, and have no tolerance for oppression. They are also very orderly and traditional, leaning heavily on the support of their community and the comfort of their old ways.',
    'Halflings average about 3 feet tall and weigh about 40 pounds. Your size is Small.',
    array[
      ROW('Lucky', array['When you roll a 1 on an attack roll, ability check, or saving throw, you can reroll the die and must use the new roll.'])::trait,
      ROW('Brave', array['You have advantage on saving throws against being frightened.'])::trait,
      ROW('Halfling Nimbleness', array['You can move through the space of any creature that is of a size larger than yours.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Halfling. The Halfling language isn''t secret, but halflings are loath to share it with others. They write very little, so they don''t have a rich body of literature. Their oral tradition, however, is very strong. Almost all halflings speak Common to converse with the people in whose lands they dwell or through which they are traveling.'])::trait,
      ROW('Naturally Stealthy', array['You can attempt to hide even when you are obscured only by a creature that is at least one size larger than you.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    12,
    'Halfling (Stout)',
    array[
    ROW('DEX', 2)::ability_mod, 
    ROW('CON', 1)::ability_mod
    ],
    'Small',
    25, 
    'A halfling reaches adulthood at the age of 20 and generally lives into the middle of his or her second century.',
    'Most halflings are lawful good. As a rule, they are good-hearted and kind, hate to see others in pain, and have no tolerance for oppression. They are also very orderly and traditional, leaning heavily on the support of their community and the comfort of their old ways.',
    'Halflings average about 3 feet tall and weigh about 40 pounds. Your size is Small.',
    array[
      ROW('Lucky', array['When you roll a 1 on an attack roll, ability check, or saving throw, you can reroll the die and must use the new roll.'])::trait,
      ROW('Brave', array['You have advantage on saving throws against being frightened.'])::trait,
      ROW('Halfling Nimbleness', array['You can move through the space of any creature that is of a size larger than yours.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Halfling. The Halfling language isn''t secret, but halflings are loath to share it with others. They write very little, so they don''t have a rich body of literature. Their oral tradition, however, is very strong. Almost all halflings speak Common to converse with the people in whose lands they dwell or through which they are traveling.'])::trait,
      ROW('Stout Resilience', array['You have advantage on saving throws against poison, and you have resistance against poison damage.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    13,
    'Human',
    array[
    ROW('STR', 1)::ability_mod, 
    ROW('DEX', 1)::ability_mod, 
    ROW('CON', 1)::ability_mod, 
    ROW('INT', 1)::ability_mod, 
    ROW('WIS', 1)::ability_mod, 
    ROW('CHA', 1)::ability_mod
    ],
    'Medium',
    30, 
    'Humans reach adulthood in their late teens and live less than a century.',
    'Humans tend toward no particular alignment. The best and the worst are found among them.',
    'Humans vary widely in height and build, from barely 5 feet to well over 6 feet tall. Regardless of your position in that range, your size is Medium.',
    array[
      ROW('Languages', array['You can speak, read, and write Common and one extra language of your choice. Humans typically learn the languages of other peoples they deal with, including obscure dialects. They are fond of sprinkling their speech with words borrowed from other tongues: Orc curses, Elvish musical expressions, Dwarvish military phrases, and so on.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    14,
    'Human (Variant)',
    array[
    ROW('UNI', 2)::ability_mod
    ],
    'Medium',
    30, 
    'Humans reach adulthood in their late teens and live less than a century.',
    'Humans tend toward no particular alignment. The best and the worst are found among them.',
    'Humans vary widely in height and build, from barely 5 feet to well over 6 feet tall. Regardless of your position in that range, your size is Medium.',
    array[
      ROW('Languages', array['You can speak, read, and write Common and one extra language of your choice. Humans typically learn the languages of other peoples they deal with, including obscure dialects. They are fond of sprinkling their speech with words borrowed from other tongues: Orc curses, Elvish musical expressions, Dwarvish military phrases, and so on.'])::trait,
      ROW('Skills', array['You gain proficiency in one skill of your choice.'])::trait,
      ROW('Feat', array['You gain one feat of your choice.'])::trait
      ]
);

INSERT INTO races
  (id, name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    15,
    'Tiefling',
    array[
    ROW('CHA', 2)::ability_mod, 
    ROW('INT', 1)::ability_mod
    ],
    'Medium',
    30, 
    'Tieflings mature at the same rate as humans but live a few years longer.',
    'Tieflings might not have an innate tendency toward evil, but many of them end up there. Evil or not, an independent nature inclines many tieflings toward a chaotic alignment.',
    'Tieflings are about the same size and build as humans. Your size is Medium.',
    array[
      ROW('Darkvision', array['Thanks to your infernal heritage, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can''t discern color in darkness, only shades of gray.'])::trait,
      ROW('Hellish Resistance', array['You have resistance to fire damage.'])::trait,
      ROW('Languages', array['You can speak, read, and write Common and Infernal.'])::trait,
      ROW('Infernal Legacy', array['You know the thaumaturgy cantrip. Once you reach 3rd level, you can cast the hellish rebuke spell as a 2nd-level spell; you must finish a long rest in order to cast the spell again using this trait. Once you reach 5th level, you can also cast the darkness spell; you must finish a long rest in order to cast the spell again using this trait. Charisma is your spellcasting ability for these spells.'])::trait
      ]
);