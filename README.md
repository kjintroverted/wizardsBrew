- [Wizard's Brew](#wizards-brew)
  - [Data](#data)
  - [Development](#development)
    - [Setup](#setup)
    - [Startup](#startup)
  - [API](#api)
    - [Endpoints](#endpoints)
      - [Races](#races)
        - [`/api/races`](#apiraces)
        - [`/api/races/{id}`](#apiracesid)
      - [Classes](#classes)
        - [`/api/classes`](#apiclasses)
        - [`/api/classes/{id}`](#apiclassesid)
      - [Backgrounds](#backgrounds)
        - [`/api/bg`](#apibg)
        - [`/api/bg/{id}`](#apibgid)
      - [Features](#features)
        - [`/api/feats`](#apifeats)
        - [`/api/feats/{id}`](#apifeatsid)
      - [Spells](#spells)
        - [`/api/spells`](#apispells)
        - [`/api/spells/{id}`](#apispellsid)
      - [Items](#items)
        - [`/api/items`](#apiitems)
        - [`/api/items/{id}`](#apiitemsid)
      - [Characters](#characters)
        - [`/api/pc`](#apipc)
        - [`/api/pc/{id}`](#apipcid)

# Wizard's Brew

A server to handle data for managing a D&D party. 

## Data

The project includes raw json data that acts as the source for much of the base SRD data. The data is then converted into SQL inserts and entered into a PSQL database.

## Development

### Setup

To setup the DB run `go run . sql`

### Startup

To start the server run `go run .`

If you are trying to connect to an external front end you need to set up CORS like 

```
ALLOW_ORIGIN=http://localhost:3000 go run .
```

## API

The api interacts with the PSQL database and serves the data out to the requester.

### Endpoints

#### Races

##### `/api/races` 
<details><summary>lists all races</summary><p>

```json
[
  {
    "id": "1",
    "name": "Dragonborn",
    "ability": [
      {
        "name": "STR",
        "mod": 2
      },
      {
        "name": "CHA",
        "mod": 1
      }
    ],
    "size": "Medium",
    "speed": 30,
    "age": "Young dragonborn grow quickly. They walk hours after hatching, attain the size and development of a 10-year-old human child by the age of 3, and reach adulthood by 15. They live to be around 80.",
    "align": "Dragonborn tend to extremes, making a conscious choice for one side or the other in the cosmic war between good and evil (represented by Bahamut and Tiamat, respectively). Most dragonborn are good, but those who side with Tiamat can be terrible villains.",
    "sizeDesc": "Dragonborn are taller and heavier than humans, standing well over 6 feet tall and averaging almost 250 pounds. Your size is Medium.",
    "traits": [
      {
        "title": "Draconic Ancestry",
        "body": [
          "You have draconic ancestry. Choose one type of dragon from the Draconic Ancestry table. Your breath weapon and damage resistance are determined by the dragon type, as shown in the table."
        ]
      },
      {
        "title": "Dragon|Damage Type|Breath Weapon",
        "body": [
          "Black|Acid|5 by 30 ft. line (Dex. save)",
          "Blue|Lightning|5 by 30 ft. line (Dex. save)",
          "Brass|Fire|5 by 30 ft. line (Dex. save)",
          "Bronze|Lightning|5 by 30 ft. line (Dex. save)",
          "Copper|Acid|5 by 30 ft. line (Dex. save)",
          "Gold|Fire|15 ft. cone (Dex. save)",
          "Green|Poison|15 ft. cone (Con. save)",
          "Red|Fire|15 ft. cone (Dex. save)",
          "Silver|Cold|15 ft. cone (Con. save)",
          "White|Cold|15 ft. cone (Con. save)"
        ]
      },
      {
        "title": "Breath Weapon",
        "body": [
          "You can use your action to exhale destructive energy. Your draconic ancestry determines the size, shape, and damage type of the exhalation.",
          "When you use your breath weapon, each creature in the area of the exhalation must make a saving throw, the type of which is determined by your draconic ancestry. The DC for this saving throw equals 8 + your Constitution modifier + your proficiency bonus. A creature takes 2d6 damage on a failed save, and half as much damage on a successful one. The damage increases to 3d6 at 6th level, 4d6 at 11th level, and 5d6 at 16th level.",
          "After you use your breath weapon, you can't use it again until you complete a short or long rest."
        ]
      },
      {
        "title": "Damage Resistance",
        "body": [
          "You have resistance to the damage type associated with your draconic ancestry."
        ]
      },
      {
        "title": "Languages",
        "body": [
          "You can speak, read, and write Common and Draconic. Draconic is thought to be one of the oldest languages and is often used in the study of magic. The language sounds harsh to most other creatures and includes numerous hard consonants and sibilants."
        ]
      }
    ]
  }
  ...
]
```
</p></details>

##### `/api/races/{id}` 
<details><summary>returns one race</summary><p>

```json
{
  "id": "5",
  "name": "Elf (High)",
  "ability": [
    {
      "name": "DEX",
      "mod": 2
    },
    {
      "name": "INT",
      "mod": 1
    }
  ],
  "size": "Medium",
  "speed": 30,
  "age": "Although elves reach physical maturity at about the same age as humans, the elven understanding of adulthood goes beyond physical growth to encompass worldly experience. An elf typically claims adulthood and an adult name around the age of 100 and can live to be 750 years old.",
  "align": "Elves love freedom, variety, and self-expression, so they lean strongly toward the gentler aspects of chaos. They value and protect others' freedom as well as their own, and they are more often good than not.",
  "sizeDesc": "Elves range from under 5 to over 6 feet tall and have slender builds. Your size is Medium.",
  "traits": [
    {
      "title": "Darkvision",
      "body": [
        "Accustomed to twilit forests and the night sky, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can't discern color in darkness, only shades of gray."
      ]
    },
    {
      "title": "Keen Senses",
      "body": [
        "You have proficiency in the Perception skill."
      ]
    },
    {
      "title": "Trance",
      "body": [
        "Elves don't need to sleep. Instead, they meditate deeply, remaining semiconscious, for 4 hours a day. (The Common word for such meditation is \\\\\"\"trance.\\\\\"\") While meditating, you can dream after a fashion; such dreams are actually mental exercises that have become reflexive through years of practice. After resting in this way, you gain the same benefit that a human does from 8 hours of sleep.",
        "If you meditate during a long rest, you finish the rest after only 4 hours. You otherwise obey all the rules for a long rest; only the duration is changed."
      ]
    },
    {
      "title": "Languages",
      "body": [
        "You can speak, read, and write Common and Elvish. Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their language so they can add Elvish ballads to their repertoires."
      ]
    },
    {
      "title": "Elf Weapon Training",
      "body": [
        "You have proficiency with the longsword, shortsword, shortbow, and longbow."
      ]
    },
    {
      "title": "Cantrip",
      "body": [
        "You know one cantrip of your choice from the wizard spell list. Intelligence is your spellcasting ability for it."
      ]
    },
    {
      "title": "Extra Language",
      "body": [
        "You can speak, read, and write one extra language of your choosing."
      ]
    }
  ]
}
```
</p></details>

#### Classes

##### `/api/classes` 
<details><summary>lists all classes</summary><p>

```json
[
  {
    "id": "7",
    "name": "Monk",
    "hitDice": "1d8",
    "proArmor": null,
    "proWeapon": [
      "simple",
      "shortswords"
    ],
    "proTool": "any one type of artisan's tools or any one musical instrument of your choice",
    "proSave": [
      "STR",
      "DEX"
    ],
    "skills": "Choose 2: acrobatics, athletics, history, insight, religion, stealth",
    "startEquip": [
      "(a) a shortsword or (b) any simple weapon",
      "(a) a dungeoneer's pack or (b) an explorer's pack",
      "10 dart"
    ],
    "description": [
      {
        "title": "",
        "body": [
          "Her fists a blur as they deflect an incoming hail of arrows, a half-elf springs over a barricade and throws herself into the massed ranks of hobgoblins on the other side. She whirls among them, knocking their blows aside and sending them reeling, until at last she stands alone.",
          "Taking a deep breath, a human covered in tattoos settles into a battle stance. As the first charging orcs reach him, he exhales and a blast of fire roars from his mouth, engulfing his foes.",
          "Moving with the silence of the night, a black-clad halfling steps into a shadow beneath an arch and emerges from another inky shadow on a balcony a stone's throw away. She slides her blade free of its cloth-wrapped scabbard and peers through the open window at the tyrant prince, so vulnerable in the grip of sleep.",
          "Whatever their discipline, monks are united in their ability to magically harness the energy that flows in their bodies. Whether channeled as a striking display of combat prowess or a subtler focus of defensive ability and speed, this energy infuses all that a monk does."
        ]
      },
      {
        "title": "The Magic of Ki",
        "body": [
          "Monks make careful study of a magical energy that most monastic traditions call ki. This energy is an element of the magic that suffuses the multiverse—specifically, the element that flows through living bodies. Monks harness this power within themselves to create magical effects and exceed their bodies' physical capabilities, and some of their special attacks can hinder the flow of ki in their opponents. Using this energy, monks channel uncanny speed and strength into their unarmed strikes. As they gain experience, their martial training and their mastery of ki gives them more power over their bodies and the bodies of their foes."
        ]
      },
      {
        "title": "Training and Asceticism",
        "body": [
          "Small walled cloisters dot the landscapes of the worlds of D\u0026D, tiny refuges from the flow of ordinary life, where time seems to stand still. The monks who live there seek personal perfection through contemplation and rigorous training. Many entered the monastery as children, sent to live there when their parents died, when food couldn't be found to support them, or in return for some kindness that the monks had performed for their families.",
          "Some monks live entirely apart from the surrounding population, secluded from anything that might impede their spiritual progress. Others are sworn to isolation, emerging only to serve as spies or assassins at the command of their leader, a noble patron, or some other mortal or divine power.",
          "The majority of monks don't shun their neighbors, making frequent visits to nearby towns or villages and exchanging their service for food and other goods. As versatile warriors, monks often end up protecting their neighbors from monsters or tyrants.",
          "For a monk, becoming an adventurer means leaving a structured, communal lifestyle to become a wanderer. This can be a harsh transition, and monks don't undertake it lightly. Those who leave their cloisters take their work seriously, approaching their adventures as personal tests of their physical and spiritual growth. As a rule, monks care little for material wealth and are driven by a desire to accomplish a greater mission than merely slaying monsters and plundering their treasure."
        ]
      },
      {
        "title": "Creating a Monk",
        "body": [
          "As you make your monk character, think about your connection to the monastery where you learned your skills and spent your formative years. Were you an orphan or a child left on the monastery's threshold? Did your parents promise you to the monastery in gratitude for a service performed by the monks? Did you enter this secluded life to hide from a crime you committed? Or did you choose the monastic life for yourself?",
          "Consider why you left. Did the head of your monastery choose you for a particularly important mission beyond the cloister? Perhaps you were cast out because of some violation of the community's rules. Did you dread leaving, or were you happy to go? Is there something you hope to accomplish outside the monastery? Are you eager to return to your home?",
          "As a result of the structured life of a monastic community and the discipline required to harness ki, monks are almost always lawful in alignment."
        ]
      }
    ],
    "progress": [
      [
        "Martial Arts",
        "1d4",
        "1d4",
        "1d4",
        "1d4",
        "1d6",
        "1d6",
        "1d6",
        "1d6",
        "1d6",
        "1d6",
        "1d8",
        "1d8",
        "1d8",
        "1d8",
        "1d8",
        "1d8",
        "1d10",
        "1d10",
        "1d10",
        "1d10"
      ],
      [
        "Ki Points",
        "0",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "10",
        "11",
        "12",
        "13",
        "14",
        "15",
        "16",
        "17",
        "18",
        "19",
        "20"
      ],
      [
        "Unarmored Movement",
        "0",
        "10",
        "10",
        "10",
        "10",
        "15",
        "15",
        "15",
        "15",
        "20",
        "20",
        "20",
        "20",
        "25",
        "25",
        "25",
        "25",
        "30",
        "30",
        "30"
      ]
    ]
  },
  ...
]
```
</p></details>

##### `/api/classes/{id}` 
<details><summary>returns one class</summary><p>


```json
{
  "id": "10",
  "name": "Rogue",
  "hitDice": "1d8",
  "proArmor": [
    "light"
  ],
  "proWeapon": [
    "simple",
    "hand crossbows",
    "longswords",
    "rapiers",
    "shortswords"
  ],
  "proTool": "thieves' tools",
  "proSave": [
    "DEX",
    "INT"
  ],
  "skills": "Choose 4: acrobatics, athletics, deception, insight, intimidation, investigation, perception, performance, persuasion, sleight of hand, stealth",
  "startEquip": [
    "(a) a rapier or (b) a shortsword",
    "(a) a shortbow and quiver of arrows (20) or (b) a shortsword",
    "(a) a burglar's pack, (b) a dungeoneer's pack, or (c) an explorer's pack",
    "Leather armor, two dagger, and thieves' tools"
  ],
  "description": [
    {
      "title": "",
      "body": [
        "Signaling for her companions to wait, a halfling creeps forward through the dungeon hall. She presses an ear to the door, then pulls out a set of tools and picks the lock in the blink of an eye. Then she disappears into the shadows as her fighter friend moves forward to kick the door open.",
        "A human lurks in the shadows of an alley while his accomplice prepares for her part in the ambush. When their target—a notorious slaver—passes the alleyway, the accomplice cries out, the slaver comes to investigate, and the assassin's blade cuts his throat before he can make a sound.",
        "Suppressing a giggle, a gnome waggles her fingers and magically lifts the key ring from the guard's belt. In a moment, the keys are in her hand, the cell door is open, and she and her companions are free to make their escape.",
        "Rogues rely on skill, stealth, and their foes' vulnerabilities to get the upper hand in any situation. They have a knack for finding the solution to just about any problem, demonstrating a resourcefulness and versatility that is the cornerstone of any successful adventuring party."
      ]
    },
    {
      "title": "Skill and Precision",
      "body": [
        "Rogues devote as much effort to mastering the use of a variety of skills as they do to perfecting their combat abilities, giving them a broad expertise that few other characters can match. Many rogues focus on stealth and deception, while others refine the skills that help them in a dungeon environment, such as climbing, finding and disarming traps, and opening locks.",
        "When it comes to combat, rogues prioritize cunning over brute strength. A rogue would rather make one precise strike, placing it exactly where the attack will hurt the target most, than wear an opponent down with a barrage of attacks. Rogues have an almost supernatural knack for avoiding danger, and a few learn magical tricks to supplement their other abilities."
      ]
    },
    {
      "title": "A Shady Living",
      "body": [
        "Every town and city has its share of rogues. Most of them live up to the worst stereotypes of the class, making a living as burglars, assassins, cutpurses, and con artists. Often, these scoundrels are organized into thieves' guilds or crime families. Plenty of rogues operate independently, but even they sometimes recruit apprentices to help them in their scams and heists. A few rogues make an honest living as locksmiths, investigators, or exterminators, which can be a dangerous job in a world where dire rats—and wererats—haunt the sewers.",
        "As adventurers, rogues fall on both sides of the law. Some are hardened criminals who decide to seek their fortune in treasure hoards, while others take up a life of adventure to escape from the law. Some have learned and perfected their skills with the explicit purpose of infiltrating ancient ruins and hidden crypts in search of treasure."
      ]
    },
    {
      "title": "Creating a Rogue",
      "body": [
        "As you create your rogue character, consider the character's relationship to the law. Do you have a criminal past—or present? Are you on the run from the law or from an angry thieves' guild master? Or did you leave your guild in search of bigger risks and bigger rewards? Is it greed that drives you in your adventures, or some other desire or ideal?",
        "What was the trigger that led you away from your previous life? Did a great con or heist gone terribly wrong cause you to reevaluate your career? Maybe you were lucky and a successful robbery gave you the coin you needed to escape the squalor of your life. Did wanderlust finally call you away from your home? Perhaps you suddenly found yourself cut off from your family or your mentor, and you had to find a new means of support. Or maybe you made a new friend—another member of your adventuring party—who showed you new possibilities for earning a living and employing your particular talents."
      ]
    }
  ],
  "progress": [
    [
      "Sneak Attack",
      "1d6",
      "1d6",
      "2d6",
      "2d6",
      "3d6",
      "3d6",
      "4d6",
      "4d6",
      "5d6",
      "5d6",
      "6d6",
      "6d6",
      "7d6",
      "7d6",
      "8d6",
      "8d6",
      "9d6",
      "9d6",
      "10d6",
      "10d6"
    ]
  ]
}
```
</p></details>

#### Backgrounds

##### `/api/bg` 
<details><summary>lists all backgrounds</summary><p>


```json
[
  {
    "id": "1",
    "name": "Acolyte",
    "proSkill": [
      "religion",
      "insight"
    ],
    "proTool": null,
    "language": [
      "Choose 2"
    ],
    "equipment": [
      "A holy symbol (a gift to you when you entered the priesthood)",
      "a prayer book or prayer wheel",
      "5 sticks of incense",
      "vestments",
      "a set of common clothes",
      "and a belt pouch containing 15 gp"
    ],
    "specialOpts": null,
    "characterOpts": null
  },
  {
    "id": "2",
    "name": "Anthropologist",
    "proSkill": [
      "insight",
      "religion"
    ],
    "proTool": null,
    "language": [
      "Choose 2"
    ],
    "equipment": [
      "A leather-bound diary",
      "a bottle of ink",
      "an ink pen",
      "a set of traveler's clothes",
      "one trinket of special significance",
      "and a pouch containing 10 gp"
    ],
    "specialOpts": ["string array of specialty options where applicable"],
    "characterOpts": null
  },
  ...
 ]
```
</p></details>

##### `/api/bg/{id}` 
<details><summary>returns one background</summary><p>


```json
{
  "id": "5",
  "name": "City Watch",
  "proSkill": [
    "athletics",
    "insight"
  ],
  "proTool": null,
  "language": [
    "Choose 2"
  ],
  "equipment": [
    "A uniform in the style of your unit and indicative of your rank",
    "a horn with which to summon help",
    "a set of manacles",
    "and a pouch containing 10 gp"
  ],
  "specialOpts": null,
  "characterOpts": null
}
```
</p></details>

#### Features

##### `/api/feats` 
<details><summary>lists features </summary><p>

  - *Query Param Options*
  - class={class name}
  - subclass={subclass name}
  - level={level number}
  - background={background name}

```json 
[
  {
    "id": "231",
    "name": "Second Wind",
    "ability": null,
    "description": [
      {
        "title": "",
        "body": [
          "You have a limited well of stamina that you can draw on to protect yourself from harm. On your turn, you can use a bonus action to regain hit points equal to 1d10 + your fighter level.",
          "Once you use this feature, you must finish a short or long rest before you can use it again."
        ]
      }
    ],
    "class": "Fighter",
    "subclass": null,
    "background": null,
    "level": 1,
    "prereq": null
  },
  ...
 ]
```
</p></details>

##### `/api/feats/{id}` 
<details><summary>returns one feature</summary><p>


```json
{
  "id": "5",
  "name": "Artificer Specialist",
  "ability": null,
  "description": [
    {
      "title": "",
      "body": [
        "At 3rd level, you choose the type of specialist you are: Alchemist, Artillerist, or Battle Smith, each of which is detailed at the end of the class's description. Your choice grants you features at 5th level and again at 9th and 15th level."
      ]
    }
  ],
  "class": "Artificer",
  "subclass": null,
  "background": null,
  "level": 3,
  "prereq": null
}
```
</p></details>

#### Spells

##### `/api/spells`
<details><summary>returns a list of spells</summary><p>
  - *Query Param Options*
  - class={class name}
  - level={level number}
  - school={magic school}
  
  ```json
  [
  {
    "id": "51",
    "name": "Clairvoyance",
    "school": "divination",
    "time": {
      "time": "10 minute"
    },
    "duration": "10 minute",
    "comp": [
      {
        "name": "verbal",
        "description": "A verbal component is a spoken incantation. To provide a verbal component, you must be able to speak in a strong voice."
      },
      {
        "name": "somatic",
        "description": "A somatic component is a measured and precise movement of the hand. You must have at least one hand free to provide a somatic component."
      },
      {
        "name": "material",
        "description": "a focus worth at least 100 gp, either a jeweled horn for hearing or a glass eye for seeing",
        "cost": 100
      }
    ],
    "concentrate": true,
    "range": "1 miles",
    "level": 3,
    "class": [
      "Bard",
      "Cleric",
      "Sorcerer",
      "Wizard"
    ],
    "description": [
      {
        "title": "",
        "body": [
          "You create an invisible sensor within range in a location familiar to you (a place you have visited or seen before) or in an obvious location that is unfamiliar to you (such as behind a door, around a corner, or in a grove of trees). The sensor remains in place for the duration, and it can't be attacked or otherwise interacted with.",
          "When you cast the spell, you choose seeing or hearing. You can use the chosen sense through the sensor as if you were in its space. As your action, you can switch between seeing and hearing.",
          "A creature that can see the sensor (such as a creature benefiting from see invisibility or truesight) sees a luminous, intangible orb about the size of your fist."
        ]
      }
    ]
  },
  ...
 ]
  ```
</p></details>

##### `/api/spells/{id}`
<details><summary>returns one spell</summary><p>

```json
{
  "id": "164",
  "name": "Greater Invisibility",
  "school": "illusion",
  "time": {
    "time": "1 action"
  },
  "duration": "1 minute",
  "comp": [
    {
      "name": "verbal",
      "description": "A verbal component is a spoken incantation. To provide a verbal component, you must be able to speak in a strong voice."
    },
    {
      "name": "somatic",
      "description": "A somatic component is a measured and precise movement of the hand. You must have at least one hand free to provide a somatic component."
    }
  ],
  "concentrate": true,
  "range": "touch",
  "level": 4,
  "class": [
    "Bard",
    "Sorcerer",
    "Wizard"
  ],
  "description": [
    {
      "title": "",
      "body": [
        "You or a creature you touch becomes invisible until the spell ends. Anything the target is wearing or carrying is invisible as long as it is on the target's person."
      ]
    }
  ]
}
```
</p></details>

#### Items

##### `/api/items` 
<details><summary>lists items</summary><p>

  - *Query Param Options* 
  - type={weapon/armor}
  
```json
[
  {
    "id": "1",
    "name": "Antimatter Rifle",
    "type": "Ranged Weapon",
    "cost": null,
    "weight": 10,
    "attune": "false",
    "rarity": "Common",
    "weapon": {
      "category": "Martial",
      "damage": "6d8",
      "damageType": "necrotic"
    },
    "ac": null
  },
  {
    "id": "10",
    "name": "Breastplate",
    "type": "Medium Armor",
    "cost": 400,
    "weight": 20,
    "attune": "false",
    "rarity": "Common",
    "weapon": {
      "category": null,
      "damage": null,
      "damageType": null
    },
    "ac": 14,
    "info": [
      {
        "title": "",
        "body": [
          "This armor consists of a fitted metal chest piece worn with supple leather. Although it leaves the legs and arms relatively unprotected, this armor provides good protection for the wearer's vital organs while leaving the wearer relatively unencumbered."
        ]
      }
    ]
  },
  ...
]
```
</p></details>  

##### `/api/items/{id}` 
<details><summary>returns one item</summary><p>


```json
{
  "id": "675",
  "name": "Ring of Fire Elemental Command",
  "type": "Ring",
  "cost": null,
  "weight": null,
  "attune": "true",
  "rarity": "Legendary",
  "weapon": {
    "category": null,
    "damage": null,
    "damageType": null
  },
  "ac": null,
  "info": [
    {
      "title": "",
      "body": [
        "While wearing this ring, you have advantage on attack rolls against elementals from the Elemental Plane of Fire and they have disadvantage on attack rolls against you. In addition, you have access to properties based on the Elemental Plane of Fire.",
        "The ring has 5 charges. It regains 1d4+1 expended charges daily at dawn. Spells cast from the ring have a save DC of 17.",
        "You can expend 2 of the ring's charges to cast dominate monster on a fire elemental. In addition, you have resistance to fire damage. You can also speak and understand Ignan.",
        "If you help slay a fire elemental while attuned to the ring, you gain access to the following additional properties:"
      ]
    },
    {
      "title": "choices",
      "body": [
        "You are immune to fire damage.",
        "You can cast the following spells from the ring, expending the necessary number of charges: burning hands (1 charge), fireball (2 charges), and wall of fire (3 charges)."
      ]
    }
  ]
}
```
</p></details>

#### Characters

##### `/api/pc` 
<details><summary>GET</summary><p>returns an array of [characters](#apipcid) that the user is authorized for</p></details>
<details><summary>POST</summary><p>
Performs and Upsert on a character. 
If the character object sent has an ID it will attempt to update at that ID, otherwise it will insert.
Request body should look like:

```json
{
  "id": 2,
  "owner": "t.pratchett@unseen.edu",
  "name": "Jerry",
  "authUsers": ["thenobbiestnobbs@watch.com"],
  "raceID": 1,
  "classID": 5,
  "backgroundID": 5,
  "stats": {
    "str": 10,
    "dex": 15,
    "con": 13,
    "int": 13,
    "wis": 8,
    "cha": 16
  },
  "xp": 300,
  "hp": 8,
  "maxHP": 8,
  "proTools": [
    "Poisoner's Tools"
  ],
  "proWeapons": [
    "Dagger",
    "Dart",
    "Sling",
    "Quarterstaff",
    "Crossbow, light"
  ],
  "languages": [
    "Common",
    "Infernal",
    "Orc",
    "Draconic"
  ],
  "equipmentIDs": [
    31
  ],
  "weaponIDs": [
    17,
    37
  ],
  "inventoryIDs": [
    15
  ],
  "gold": 10.2,
  "spellIDs": [
    46,
    37
  ]
}
```
</p></details>

##### `/api/pc/{id}`
<details><summary>DELETE</summary><p>deletes a character by id</p></details>
<details><summary>GET</summary><p> 
Returns character data, including details on all relevant field ids

```json
{
  "background": {
    "id": "3",
    "name": "Archaeologist",
    "proSkill": [
      "history",
      "survival"
    ],
    "proTool": [
      "navigator's tools or cartographer's tools"
    ],
    "language": [
      "Choose 1"
    ],
    "equipment": [
      "A wooden case containing a map to a ruin or dungeon",
      "a bullseye lantern",
      "a miner's pick",
      "a set of traveler's clothes",
      "a shovel",
      "a two-person tent",
      "a trinket recovered from a dig site",
      "and a pouch containing 25 gp"
    ],
    "specialOpts": null,
    "characterOpts": null
  },
  "class": {
    "id": "10",
    "name": "Rogue",
    "hitDice": "1d8",
    "proArmor": [
      "light"
    ],
    "proWeapon": [
      "simple",
      "hand crossbows",
      "longswords",
      "rapiers",
      "shortswords"
    ],
    "proTool": "thieves' tools",
    "proSave": [
      "DEX",
      "INT"
    ],
    "skills": "Choose 4: acrobatics, athletics, deception, insight, intimidation, investigation, perception, performance, persuasion, sleight of hand, stealth",
    "startEquip": [
      "(a) a rapier or (b) a shortsword",
      "(a) a shortbow and quiver of arrows (20) or (b) a shortsword",
      "(a) a burglar's pack, (b) a dungeoneer's pack, or (c) an explorer's pack",
      "Leather armor, two dagger, and thieves' tools"
    ],
    "description": [
      {
        "title": "",
        "body": [
          "Signaling for her companions to wait, a halfling creeps forward through the dungeon hall. She presses an ear to the door, then pulls out a set of tools and picks the lock in the blink of an eye. Then she disappears into the shadows as her fighter friend moves forward to kick the door open.",
          "A human lurks in the shadows of an alley while his accomplice prepares for her part in the ambush. When their target—a notorious slaver—passes the alleyway, the accomplice cries out, the slaver comes to investigate, and the assassin's blade cuts his throat before he can make a sound.",
          "Suppressing a giggle, a gnome waggles her fingers and magically lifts the key ring from the guard's belt. In a moment, the keys are in her hand, the cell door is open, and she and her companions are free to make their escape.",
          "Rogues rely on skill, stealth, and their foes' vulnerabilities to get the upper hand in any situation. They have a knack for finding the solution to just about any problem, demonstrating a resourcefulness and versatility that is the cornerstone of any successful adventuring party."
        ]
      },
      {
        "title": "Skill and Precision",
        "body": [
          "Rogues devote as much effort to mastering the use of a variety of skills as they do to perfecting their combat abilities, giving them a broad expertise that few other characters can match. Many rogues focus on stealth and deception, while others refine the skills that help them in a dungeon environment, such as climbing, finding and disarming traps, and opening locks.",
          "When it comes to combat, rogues prioritize cunning over brute strength. A rogue would rather make one precise strike, placing it exactly where the attack will hurt the target most, than wear an opponent down with a barrage of attacks. Rogues have an almost supernatural knack for avoiding danger, and a few learn magical tricks to supplement their other abilities."
        ]
      },
      {
        "title": "A Shady Living",
        "body": [
          "Every town and city has its share of rogues. Most of them live up to the worst stereotypes of the class, making a living as burglars, assassins, cutpurses, and con artists. Often, these scoundrels are organized into thieves' guilds or crime families. Plenty of rogues operate independently, but even they sometimes recruit apprentices to help them in their scams and heists. A few rogues make an honest living as locksmiths, investigators, or exterminators, which can be a dangerous job in a world where dire rats—and wererats—haunt the sewers.",
          "As adventurers, rogues fall on both sides of the law. Some are hardened criminals who decide to seek their fortune in treasure hoards, while others take up a life of adventure to escape from the law. Some have learned and perfected their skills with the explicit purpose of infiltrating ancient ruins and hidden crypts in search of treasure."
        ]
      },
      {
        "title": "Creating a Rogue",
        "body": [
          "As you create your rogue character, consider the character's relationship to the law. Do you have a criminal past—or present? Are you on the run from the law or from an angry thieves' guild master? Or did you leave your guild in search of bigger risks and bigger rewards? Is it greed that drives you in your adventures, or some other desire or ideal?",
          "What was the trigger that led you away from your previous life? Did a great con or heist gone terribly wrong cause you to reevaluate your career? Maybe you were lucky and a successful robbery gave you the coin you needed to escape the squalor of your life. Did wanderlust finally call you away from your home? Perhaps you suddenly found yourself cut off from your family or your mentor, and you had to find a new means of support. Or maybe you made a new friend—another member of your adventuring party—who showed you new possibilities for earning a living and employing your particular talents."
        ]
      }
    ],
    "progress": [
      [
        "Sneak Attack",
        "1d6",
        "1d6",
        "2d6",
        "2d6",
        "3d6",
        "3d6",
        "4d6",
        "4d6",
        "5d6",
        "5d6",
        "6d6",
        "6d6",
        "7d6",
        "7d6",
        "8d6",
        "8d6",
        "9d6",
        "9d6",
        "10d6",
        "10d6"
      ]
    ]
  },
  "equipment": null,
  "info": {
    "id": 1,
    "name": "Carrion",
    "owner": "wkgreen13@gmail.com",
    "authUsers": null,
    "readUsers": null,
    "raceID": 15,
    "classID": 10,
    "subclass": "",
    "backgroundID": 3,
    "stats": {
      "str": 10,
      "dex": 15,
      "con": 13,
      "int": 13,
      "wis": 8,
      "cha": 16
    },
    "xp": 300,
    "hp": 8,
    "maxHP": 8,
    "initiative": 0,
    "proSkills": null,
    "proTools": [
      "Poisoner's Tools"
    ],
    "proWeapons": null,
    "languages": [
      "Common",
      "Infernal",
      "Orc",
      "Draconic"
    ],
    "equipmentIDs": null,
    "weaponIDs": [
      17,
      37
    ],
    "inventoryIDs": [
      535,
      215,
      486,
      57,
      15,
      750,
      547,
      346,
      711,
      335,
      102,
      227,
      718
    ],
    "gold": 0,
    "spellIDs": [
      46,
      137,
      298,
      340,
      127,
      214
    ],
    "specFeatIDs": null
  },
  "inventory": [
    {
      "id": "15",
      "name": "Crossbow Bolts (20)",
      "type": "Ammunition",
      "cost": 1,
      "weight": 1.5,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null
    },
    {
      "id": "57",
      "name": "Renaissance Bullets (10)",
      "type": "Futuristic",
      "cost": 3,
      "weight": 2,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null
    },
    {
      "id": "102",
      "name": "Armor of Invulnerability",
      "type": "Heavy Armor",
      "cost": null,
      "weight": 65,
      "attune": "true",
      "rarity": "Legendary",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": 18,
      "info": [
        {
          "title": "",
          "body": [
            "You have resistance to nonmagical damage while you wear this armor. Additionally, you can use an action to make yourself immune to nonmagical damage for 10 minutes or until you are no longer wearing the armor. Once this special action is used, it can't be used again until the next dawn."
          ]
        }
      ]
    },
    {
      "id": "215",
      "name": "Carved ivory statuette",
      "type": "Treasure",
      "cost": 250,
      "weight": null,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null
    },
    {
      "id": "227",
      "name": "Chrysoprase",
      "type": "Treasure",
      "cost": 50,
      "weight": null,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null,
      "info": [
        {
          "title": "",
          "body": [
            "A translucent green gemstone."
          ]
        }
      ]
    },
    {
      "id": "335",
      "name": "Feed (per day)",
      "type": "Tack and Harness",
      "cost": 0.05,
      "weight": 10,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null
    },
    {
      "id": "346",
      "name": "Fine gold chain set with a fire opal",
      "type": "Treasure",
      "cost": 2500,
      "weight": null,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null
    },
    {
      "id": "486",
      "name": "Lapis Lazuli",
      "type": "Treasure",
      "cost": 10,
      "weight": null,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null,
      "info": [
        {
          "title": "",
          "body": [
            "An opaque light and dark blue with yellow flecks gemstone."
          ]
        }
      ]
    },
    {
      "id": "535",
      "name": "Mule",
      "type": "Mount",
      "cost": 8,
      "weight": null,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null
    },
    {
      "id": "547",
      "name": "Obsidian statuette with gold fittings and inlay",
      "type": "Treasure",
      "cost": 750,
      "weight": null,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null
    },
    {
      "id": "711",
      "name": "Rod of Absorption",
      "type": "Rod",
      "cost": null,
      "weight": 2,
      "attune": "true",
      "rarity": "Very Rare",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null,
      "info": [
        {
          "title": "",
          "body": [
            "While holding this rod, you can use your reaction to absorb a spell that is targeting only you and not with an area of effect. The absorbed spell's effect is canceled, and the spell's energy—not the spell itself—is stored in the rod. The energy has the same level as the spell when it was cast. The rod can absorb and store up to 50 levels of energy over the course of its existence. Once the rod absorbs 50 levels of energy, it can't absorb more. If you are targeted by a spell that the rod can't store, the rod has no effect on that spell.",
            "When you become attuned to the rod, you know how many levels of energy the rod has absorbed over the course of its existence, and how many levels of spell energy it currently has stored.",
            "If you are a spellcaster holding the rod, you can convert energy stored in it into spell slots to cast spells you have prepared or know. You can create spell slots only of a level equal to or lower than your own spell slots, up to a maximum of 5th level. You use the stored levels in place of your slots, but otherwise cast the spell as normal. For example, you can use 3 levels stored in the rod as a 3rd-level spell slot.",
            "A newly found rod has 1d10 levels of spell energy stored in it already. A rod that can no longer absorb spell energy and has no energy remaining becomes nonmagical."
          ]
        }
      ]
    },
    {
      "id": "718",
      "name": "Rod of the Pact Keeper, +2",
      "type": "Rod",
      "cost": null,
      "weight": 2,
      "attune": "by a Warlock",
      "rarity": "Rare",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null,
      "info": [
        {
          "title": "",
          "body": [
            "While holding this rod, you gain a +2 bonus to spell attack rolls and to the saving throw DCs of your warlock spells.",
            "In addition, you can regain one warlock spell slot as an action while holding the rod. You can't use this property again until you finish a long rest."
          ]
        }
      ]
    },
    {
      "id": "750",
      "name": "Sending Stones",
      "type": "Wondrous Item",
      "cost": null,
      "weight": null,
      "attune": "false",
      "rarity": "Uncommon",
      "weapon": {
        "category": null,
        "damage": null,
        "damageType": null
      },
      "ac": null,
      "info": [
        {
          "title": "",
          "body": [
            "Sending stones come in pairs, with each smooth stone carved to match the other so the pairing is easily recognized. While you touch one stone, you can use an action to cast the sending spell from it. The target is the bearer of the other stone. If no creature bears the other stone, you know that fact as soon as you use the stone and don't cast the spell.",
            "Once sending is cast through the stones, they can't be used again until the next dawn. If one of the stones in a pair is destroyed, the other one becomes nonmagical."
          ]
        }
      ]
    }
  ],
  "level": {
    "level": 2,
    "proBonus": 2,
    "next": 900
  },
  "race": {
    "id": "15",
    "name": "Tiefling",
    "ability": [
      {
        "name": "CHA",
        "mod": 2
      },
      {
        "name": "INT",
        "mod": 1
      }
    ],
    "size": "Medium",
    "speed": 30,
    "age": "Tieflings mature at the same rate as humans but live a few years longer.",
    "align": "Tieflings might not have an innate tendency toward evil, but many of them end up there. Evil or not, an independent nature inclines many tieflings toward a chaotic alignment.",
    "sizeDesc": "Tieflings are about the same size and build as humans. Your size is Medium.",
    "traits": [
      {
        "title": "Darkvision",
        "body": [
          "Thanks to your infernal heritage, you have superior vision in dark and dim conditions. You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can't discern color in darkness, only shades of gray."
        ]
      },
      {
        "title": "Hellish Resistance",
        "body": [
          "You have resistance to fire damage."
        ]
      },
      {
        "title": "Languages",
        "body": [
          "You can speak, read, and write Common and Infernal."
        ]
      },
      {
        "title": "Infernal Legacy",
        "body": [
          "You know the thaumaturgy cantrip. Once you reach 3rd level, you can cast the hellish rebuke spell as a 2nd-level spell; you must finish a long rest in order to cast the spell again using this trait. Once you reach 5th level, you can also cast the darkness spell; you must finish a long rest in order to cast the spell again using this trait. Charisma is your spellcasting ability for these spells."
        ]
      }
    ]
  },
  "spells": [
    {
      "id": "46",
      "name": "Charm Person",
      "school": "enchantment",
      "time": {
        "time": "1 action"
      },
      "duration": "1 hour",
      "comp": [
        {
          "name": "verbal",
          "description": "A verbal component is a spoken incantation. To provide a verbal component, you must be able to speak in a strong voice."
        },
        {
          "name": "somatic",
          "description": "A somatic component is a measured and precise movement of the hand. You must have at least one hand free to provide a somatic component."
        }
      ],
      "concentrate": false,
      "range": "30 feet",
      "level": 1,
      "class": [
        "Bard",
        "Druid",
        "Sorcerer",
        "Warlock",
        "Wizard"
      ],
      "description": [
        {
          "title": "",
          "body": [
            "You attempt to charm a humanoid you can see within range. It must make a Wisdom saving throw, and does so with advantage if you or your companions are fighting it. If it fails the saving throw, it is charmed by you until the spell ends or until you or your companions do anything harmful to it. The charmed creature regards you as a friendly acquaintance. When the spell ends, the creature knows it was charmed by you."
          ]
        },
        {
          "title": "At Higher Levels",
          "body": [
            "When you cast this spell using a spell slot of 2nd level or higher, you can target one additional creature for each slot level above 1st. The creatures must be within 30 feet of each other when you target them."
          ]
        }
      ]
    },
    {
      "id": "127",
      "name": "Faerie Fire",
      "school": "evocation",
      "time": {
        "time": "1 action"
      },
      "duration": "1 minute",
      "comp": [
        {
          "name": "verbal",
          "description": "A verbal component is a spoken incantation. To provide a verbal component, you must be able to speak in a strong voice."
        }
      ],
      "concentrate": true,
      "range": "60 feet",
      "level": 1,
      "class": [
        "Bard",
        "Druid",
        "Artificer"
      ],
      "description": [
        {
          "title": "",
          "body": [
            "Each object in a 20-foot cube within range is outlined in blue, green, or violet light (your choice). Any creature in the area when the spell is cast is also outlined in light if it fails a Dexterity saving throw. For the duration, objects and affected creatures shed dim light in a 10-foot radius.",
            "Any attack roll against an affected creature or object has advantage if the attacker can see it, and the affected creature or object can't benefit from being invisible."
          ]
        }
      ]
    },
    {
      "id": "137",
      "name": "Finger of Death",
      "school": "necromancy",
      "time": {
        "time": "1 action"
      },
      "duration": "instant",
      "comp": [
        {
          "name": "verbal",
          "description": "A verbal component is a spoken incantation. To provide a verbal component, you must be able to speak in a strong voice."
        },
        {
          "name": "somatic",
          "description": "A somatic component is a measured and precise movement of the hand. You must have at least one hand free to provide a somatic component."
        }
      ],
      "concentrate": false,
      "range": "60 feet",
      "level": 7,
      "class": [
        "Sorcerer",
        "Warlock",
        "Wizard"
      ],
      "description": [
        {
          "title": "",
          "body": [
            "You send negative energy coursing through a creature that you can see within range, causing it searing pain. The target must make a Constitution saving throw. It takes 7d8 + 30 necrotic damage on a failed save, or half as much damage on a successful one.",
            "A humanoid killed by this spell rises at the start of your next turn as a zombie that is permanently under your command, following your verbal orders to the best of its ability."
          ]
        }
      ]
    },
    {
      "id": "214",
      "name": "Magic Jar",
      "school": "necromancy",
      "time": {
        "time": "1 minute"
      },
      "duration": "permanent",
      "comp": [
        {
          "name": "verbal",
          "description": "A verbal component is a spoken incantation. To provide a verbal component, you must be able to speak in a strong voice."
        },
        {
          "name": "somatic",
          "description": "A somatic component is a measured and precise movement of the hand. You must have at least one hand free to provide a somatic component."
        },
        {
          "name": "material",
          "description": "a gem, crystal, reliquary, or some other ornamental container worth at least 500 gp",
          "cost": 500
        }
      ],
      "concentrate": false,
      "range": "self",
      "level": 6,
      "class": [
        "Wizard"
      ],
      "description": [
        {
          "title": "",
          "body": [
            "Your body falls into a catatonic state as your soul leaves it and enters the container you used for the spell's material component. While your soul inhabits the container, you are aware of your surroundings as if you were in the container's space. You can't move or use reactions. The only action you can take is to project your soul up to 100 feet out of the container, either returning to your living body (and ending the spell) or attempting to possess a humanoids body.",
            "You can attempt to possess any humanoid within 100 feet of you that you can see (creatures warded by a protection from evil and good or magic circle spell can't be possessed). The target must make a Charisma saving throw. On a failure, your soul moves into the target's body, and the target's soul becomes trapped in the container. On a success, the target resists your efforts to possess it, and you can't attempt to possess it again for 24 hours.",
            "Once you possess a creature's body, you control it. Your game statistics are replaced by the statistics of the creature, though you retain your alignment and your Intelligence, Wisdom, and Charisma scores. You retain the benefit of your own class features. If the target has any class levels, you can't use any of its class features.",
            "Meanwhile, the possessed creature's soul can perceive from the container using its own senses, but it can't move or take actions at all.",
            "While possessing a body, you can use your action to return from the host body to the container if it is within 100 feet of you, returning the host creature's soul to its body. If the host body dies while you're in it, the creature dies, and you must make a Charisma saving throw against your own spellcasting DC. On a success, you return to the container if it is within 100 feet of you. Otherwise, you die.",
            "If the container is destroyed or the spell ends, your soul immediately returns to your body. If your body is more than 100 feet away from you or if your body is dead when you attempt to return to it, you die. If another creature's soul is in the container when it is destroyed, the creature's soul returns to its body if the body is alive and within 100 feet. Otherwise, that creature dies.",
            "When the spell ends, the container is destroyed."
          ]
        }
      ]
    },
    {
      "id": "298",
      "name": "Shillelagh",
      "school": "transmutation",
      "time": {
        "time": "1 bonus"
      },
      "duration": "1 minute",
      "comp": [
        {
          "name": "verbal",
          "description": "A verbal component is a spoken incantation. To provide a verbal component, you must be able to speak in a strong voice."
        },
        {
          "name": "somatic",
          "description": "A somatic component is a measured and precise movement of the hand. You must have at least one hand free to provide a somatic component."
        },
        {
          "name": "material",
          "description": "mistletoe, a shamrock leaf, and a club or quarterstaff"
        }
      ],
      "concentrate": false,
      "range": "touch",
      "level": 0,
      "class": [
        "Druid"
      ],
      "description": [
        {
          "title": "",
          "body": [
            "The wood of a club or quarterstaff you are holding is imbued with nature's power. For the duration, you can use your spellcasting ability instead of Strength for the attack and damage rolls of melee attacks using that weapon, and the weapon's damage die becomes a d8. The weapon also becomes magical, if it isn't already. The spell ends if you cast it again or if you let go of the weapon."
          ]
        }
      ]
    },
    {
      "id": "340",
      "name": "True Seeing",
      "school": "divination",
      "time": {
        "time": "1 action"
      },
      "duration": "1 hour",
      "comp": [
        {
          "name": "verbal",
          "description": "A verbal component is a spoken incantation. To provide a verbal component, you must be able to speak in a strong voice."
        },
        {
          "name": "somatic",
          "description": "A somatic component is a measured and precise movement of the hand. You must have at least one hand free to provide a somatic component."
        },
        {
          "name": "material",
          "description": "an ointment for the eyes that costs 25 gp; is made from mushroom powder, saffron, and fat; and is consumed by the spell",
          "cost": 25,
          "consume": true
        }
      ],
      "concentrate": false,
      "range": "touch",
      "level": 6,
      "class": [
        "Bard",
        "Cleric",
        "Sorcerer",
        "Warlock",
        "Wizard"
      ],
      "description": [
        {
          "title": "",
          "body": [
            "This spell gives the willing creature you touch the ability to see things as they actually are. For the duration, the creature has truesight, notices secret doors hidden by magic, and can see into the Ethereal Plane, all out to a range of 120 feet."
          ]
        }
      ]
    }
  ],
  "weapons": [
    {
      "id": "17",
      "name": "Dagger",
      "type": "Melee Weapon",
      "cost": 2,
      "weight": 1,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": "Simple",
        "damage": "1d4",
        "damageType": "piercing"
      },
      "ac": null
    },
    {
      "id": "37",
      "name": "Laser Rifle",
      "type": "Ranged Weapon",
      "cost": null,
      "weight": 7,
      "attune": "false",
      "rarity": "Common",
      "weapon": {
        "category": "Martial",
        "damage": "3d8",
        "damageType": "radiant"
      },
      "ac": null
    }
  ]
}
```
</p></details>
