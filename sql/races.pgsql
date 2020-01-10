SELECT * from races;

INSERT INTO races
  (name, ability, size, speed, age, align, size_desc, traits)
VALUES
  (
    'Dwarf (Hill)',
    array[
    ROW('Str', 2)::ability_mod, 
    ROW('Wis', 1)::ability_mod
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