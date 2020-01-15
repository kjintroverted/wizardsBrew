INSERT INTO story_nodes (id, type, value, paths)
VALUES (1, 'terminus', 'Gnome', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (2, 'terminus', 'Dwarf', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (3, 'terminus', 'Elf', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (4, 'terminus', 'Half-Elf', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (5, 'terminus', 'Half-Orc', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (6, 'terminus', 'Dragonborn', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (7, 'terminus', 'Tiefling', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (8, 'terminus', 'Halfling', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (9, 'terminus', 'Human', null);

INSERT INTO story_nodes (id, type, value, paths)
VALUES (10, 'choice', 'Where does your character come from?', array[
  row('Deep in the woods', 1)::choice,
  row('High in the mountains', 2)::choice
]);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (11, 'choice', 'Where does your character come from?', array[
  row('A far away land', 3)::choice,
  row('A local town', 4)::choice
]);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (12, 'choice', 'What matters more to you?', array[
  row('Being accepted and valued', 5)::choice,
  row('POWER!', 6)::choice
]);

INSERT INTO story_nodes (id, type, value, paths)
VALUES (13, 'choice', 'But how tall is your character?', array[
  row('Pretty short actually', 10)::choice,
  row('Tall', 11)::choice
]);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (14, 'choice', 'Would you rather be successful or differnt?', array[
  row('I wanna be a winner', 12)::choice,
  row('I wanna be a werido', 7)::choice
]);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (15, 'choice', 'What makes your character differnt from you?', array[
  row('They''re shorter...', 8)::choice,
  row('They have swords and stuff', 9)::choice,
  row('Actually, I wanna be diiferent', 14)::choice,
  row('Maybe I meant fantasy character', 13)::choice
]);

INSERT INTO story_nodes (id, type, value, paths)
VALUES (16, 'choice', 'Why do you wanna play a fantasy RPG?', array[
  row('I love classic fanasy characters', 13)::choice,
  row('I want to do something out of the norm', 14)::choice,
  row('I want to go on adventures as myself', 15)::choice
]);

