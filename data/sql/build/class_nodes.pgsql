INSERT INTO story_nodes (id, type, value, paths)
VALUES (17, 'terminus', 'Rogue', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (18, 'terminus', 'Monk', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (19, 'terminus', 'Ranger', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (20, 'terminus', 'Druid', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (21, 'terminus', 'Wizard', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (22, 'terminus', 'Warlock', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (23, 'terminus', 'Sorcerer', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (24, 'terminus', 'Fighter', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (25, 'terminus', 'Paladin', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (26, 'terminus', 'Cleric', null);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (27, 'terminus', 'Bard', null);

INSERT INTO story_nodes (id, type, value, paths)
VALUES (28, 'choice', 'What is their preferred stealth?', array[
  row('They move in the dark', 17)::choice,
  row('Go through the chapel for santuary', 18)::choice,
  row('Find a vacant wall and scall to a window', 19)::choice
]);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (29, 'choice', 'And they would use their magic too...', array[
  row('Levitate up to the top of the wall', 21)::choice,
  row('Break down the wall', 22)::choice,
  row('Distract the guards to make an opening', 23)::choice
]);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (30, 'choice', 'What is their go-to guise?', array[
  row('A courier with a message for the ruler', 17)::choice,
  row('An unassuming beggar', 21)::choice,
  row('Foreign royalty on a visit', 22)::choice,
  row('A member of the guard', 23)::choice
]);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (31, 'choice', 'On what grounds?', array[
  row('A challenge to the guard prove their worth and honor', 24)::choice,
  row('Call on the priest of the chapel to vouch for their character', 25)::choice
]);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (32, 'choice', 'How could would they convice the guard?', array[
  row('Befriend them with charm and wit', 27)::choice,
  row('By detailing their divine mission and right', 26)::choice
]);

INSERT INTO story_nodes (id, type, value, paths)
VALUES (33, 'choice', 'What means would they use to discover another way?', array[
  row('They know magic, that should work', 29)::choice,
  row('They are very sneaky', 28)::choice,
  row('Turn into an awesome eagle and fly in, of course', 20)::choice
]);
INSERT INTO story_nodes (id, type, value, paths)
VALUES (34, 'choice', 'How would they get through?', array[
  row('They are a master of disguise', 30)::choice,
  row('They demand entry, why should they be kept out?', 31)::choice,
  row('Convice the guards to let them in', 32)::choice
]);

INSERT INTO story_nodes (id, type, value, paths)
VALUES (35, 'choice', 'Your character must infiltrate a well-guarded castle, how do they get in?', array[
  row('Through the front door!', 34)::choice,
  row('...I don''t know, but not the front door', 33)::choice
]);

