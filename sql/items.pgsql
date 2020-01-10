select * from items;

INSERT INTO items (name, type, cost, weight, attune, rarity, damage, damage_type, description)
VALUES ('Abacus', array['Adventuring Gear'::item_type], 2, 2, false, 'Common', null, null, null);

INSERT INTO items (name, type, cost, weight, attune, rarity, damage, damage_type, description)
VALUES ('Greataxe', array['Martial Weapon'::item_type, 'Melee Weapon'::item_type], 30, 7, false, 'Common', '1d12', 'slashing', 
array[
  'Heavy. Small creatures have disadvantage on attack rolls with heavy weapons. A heavy weapon''s size and bulk make it too large for a Small creature to use effectively.',
  'Two-Handed. This weapon requires two hands to use. This property is relevant only when you attack with the weapon, not when you simply hold it.'
]);