package main

var Ex_flags= 
[]string{ 
"isdoor", "closed", "locked", "secret", "swim", "pickproof", "fly", "climb",
"dig", "eatkey", "nopassdoor", "hidden", "passage", "portal", "r1", "r2",
"can_climb", "can_enter", "can_leave", "auto", "noflee", "searchable", 
"bashed", "bashproof", "nomob", "window", "can_look", "isbolt", "bolted"};

var Sec_flags=
[]string{
"inside", "city", "field", "forest", "hills", "mountain", "water_swim",
"water_noswim", "underwater", "air", "desert", "dunno", "oceanfloor",
"underground", "lava", "swamp", "r1", "r2", "r3", "r4", "r5", "r6", "r7",
"r8", "r9", "r10", "r11", "r12", "r13", "r14", "r15", "r16"};

var R_flags=
[]string{
"dark", "death", "nomob", "indoors", "lawful", "neutral", "chaotic",
"nomagic", "tunnel", "private", "safe", "solitary", "petshop", "norecall",
"donation", "nodropall", "silence", "logspeech", "nodrop", "clanstoreroom",
"nosummon", "noastral", "teleport", "teleshowdesc", "nofloor",
"nosupplicate", "arena", "nomissile", "frog", "track", "prototype", "dnd"};

var R_flags2=
[]string{
 "library", "river", "water1", "water2", "water3", "hospital", 
"canfish", "noquit", "tree", "kori",
"wipe", "noanki", "bank", "dummy", "mist1", "mist2", "mist3", "fog", "sand1", "sand2", 
"sand3", "sandstorm", "light", "bindspot", "nomission", "home_vacant", "home_room",
"home_free", "home_bindspot", "morgue", "r30", "r31", "r32"};

var O_flags =
[]string{
"glow", "hum", "dark", "loyal", "evil", "invis", "magic", "nodrop", "bless",
"antigood", "antievil", "antineutral", "noremove", "inventory",
"antimage", "antithief", "antiwarrior", "anticleric", "organic", "metal",
"donation", "clanobject", "clancorpse", "antivampire", "antidruid", 
"hidden", "poisoned", "covering", "deathrot", "buried", "prototype",  
"nolocate", "groundrot", "lootable", "quest", "legendary", "binding",
"rename", "throwable", "nothrow", "house_buy", "house_retire",
};

var Mag_flags=
[]string{
"returning", "backstabber", "bane", "loyal", "haste", "drain", 
"lightning_blade",
};

var W_flags=
[]string{
"take", "finger", "neck", "body", "head", "legs", "feet", "hands", "arms",
"shield", "about", "waist", "wrist", "wield", "hold", "_dual_", "ears", "eyes",
"missile", "back", "face", "ankle", "hip","r5","r6",
"r7","r8","r9","r10","r11","r12","r13",
};

var Item_w_flags=
[]string{
"take", "finger", "finger", "neck", "neck",  "body", "head", "legs", "feet",
"hands", "arms", "shield", "about", "waist", "wrist", "wrist", "wield",
"hold", "dual", "ears", "eyes", "missile", "back", "face", "ankle", "ankle",
"hip","r5","r6","r7","r8","r9","r10","r11","r12","r13",
};

/* Don't touch r4!!! it is "" for a reason! */
var Area_flags=
[]string {
"nopkill", "freekill", "noteleport", "spelllimit", "prototype", "hidden", "nomission", "r7", "r8",
"r9", "r10", "r11", "r12", "r13", "r14", "r15", "r16", "r17",
"r18", "r19","r20","r21","r22","r23","r24",
"r25","r26","r27","r28","r29","r30","r31",
};

var O_types=
[]string{
"none", "light", "scroll", "wand", "staff", "weapon", "_fireweapon", "_missile",
"treasure", "armor", "potion", "_worn", "furniture", "trash", "_oldtrap",
"container", "_note", "drinkcon", "key", "food", "money", "pen", "boat",
"corpse", "corpse_pc", "fountain", "pill", "blood", "bloodstain",
"scraps", "pipe", "herbcon", "herb", "incense", "fire", "book", "switch",
"lever", "pullchain", "button", "dial", "rune", "runepouch", "match", "trap",
"map", "portal", "paper", "tinder", "lockpick", "spike", "disease", "oil",
"fuel", "_empty1", "_empty2", "missileweapon", "projectile", "quiver", "shovel",
"salve", "cook", "keyring", "odor", "chance", "rod", "transmittor", "fish",
"mission", "skill_scroll", "raid_shrine",
};

var A_types=
[]string{
"none", "strength", "dexterity", "intelligence", "wisdom", "constitution",
"sex", "class", "level", "age", "height", "weight", "mana", "hit", "move",
"gold", "experience", "armor", "hitroll", "damroll", "save_poison", "save_rod",
"save_para", "save_breath", "save_spell", "charisma", "affected", "resistant",
"immune", "susceptible", "weaponspell", "luck", "backstab", "pick", "track",
"steal", "sneak", "hide", "palm", "detrap", "dodge", "peek", "scan", "gouge",
"search", "mount", "disarm", "kick", "parry", "bash", "stun", "punch", "climb",
"grip", "scribe", "brew", "wearspell", "removespell", "emotion", "mentalstate",
"stripsn", "remove", "dig", "full", "thirst", "drunk", "blood", "cook",
"recurringspell", "contagious", "xaffected", "odor", "roomflag", "sectortype",
"roomlight", "televnum", "teledelay", "drain", "navigation", "houkou",
"meiki", "chouraku", "onibi", "futsuriai", "meiun_hippaku", "bunshin",
"kushin_katsu", "wabigoe", "death_timer", "buff",
};

var A_flags=
[]string{
"blind", "invisible", "detect_evil", "detect_invis", "detect_magic",
"detect_hidden", "hold", "sanctuary", "faerie_fire", "infrared", "curse",
"_flaming", "poison", "protect", "_paralysis", "sneak", "hide", "sleep",
"charm", "flying", "pass_door", "floating", "truesight", "detect_traps",
"scrying", "fireshield", "shockshield", "r1", "iceshield", "possess", 
"berserk", "aqua_breath", "recurringspell", "contagious", "acidmist",
"venomshield", "sharingan", "implantsharingan", "frogstomach",
"byakugan", "monster", "sand", "baikyuu_kesson", "dizzy",
"kyouji_senbotsusha", "shoumei", "drain_hp", "drain_chakra",
"kori_shinchu", "houkou_taijutsu", "houkou_ninjutsu", "meiki",
"chouraku", "onibi", "futsuriai_chakra", "futsuriai_stamina",
"meiun_hippaku", "bunshin", "kushin_katsu", "kaiten_shuriken",
"suna_no_yoroi", "wabigoe", "death_timer", "implant_byakugan",
"skill_cooldown", "numb", "raidan", "amaterasu", "kukyo kogeki",
"room_water", "hachimon", "weak", "byakugou no in", "hisoukan",
"lucky", "enka", "max",
};

var Act_flags=
[]string{
"npc", "sentinel", "scavenger", "day", "night", "aggressive", "stayarea",
"wimpy", "pet", "train", "practice", "immortal", "deadly", "polyself",
"meta_aggr", "guardian", "running", "nowander", "mountable", "mounted",
"scholar", "secretive", "hardhat", "mobinvis", "noassist", "autonomous",
"pacifist", "noattack", "annoying", "statshield", "prototype", "clone",
"waterclone", "sandclone","nocorpse", "assist", "notrashtalk", "kageclone",
"bloodtype", "handseal" , "animal", "rumor", "messege", "kage",
"undertaker", "combo", "corpseclone", "powerstruggle", "in_prog", "dai",
"leap", "ignoretactic", "shunshin", "raider", "raidboss", "mission",
"noexp", "lightningclone", "nomission", "nohistorybonus", "raidhalloween",
"raidexpexempt", "raiduber", "raidhostage",
};

var Pc_flags=
[]string{
"r1", "deadly", "unauthed", "norecall", "nointro", "gag", "retired", "guest",
/* changed "r8" to "" so players on watch can't see it  -- Gorog */
"nosummon", "pager", "notitled", "groupwho", "diagnose", "highgag", "",
"nstart", "dnd", "idle", "noeiyorei", "notused",
"privacy", "toggle", "semi", "handseal", "grownup", "exitnames", 
"fishing", "botcheck", "tprompt", "helper", "silent", "immforce",
};

var Plr_flags=
[]string{
"npc", "boughtpet", "shovedrag", "autoexits", "autoloot", "autosac", "blank", 
"outcast", "brief", "combine", "prompt", "telnet_ga", "holylight", 
"wizinvis", "roomvnum","silence", "noemote", "attacker", "notell", "log", 
"deny", "freeze", "thief","killer", "litterbug", "ansi", "rip", "nice", 
"flee" ,"autogold", "automap", "afk", "invisprompt", "showexp",
"monster", "byakugan", "sand", "hardcore", "notused1", "notused2",
"notused3", "notused4",
"hiwarn", "notalk", "oocshow", "callexhaust", "exempt", "nowho", "notify", 
"rpauth", "unused", "olddesc", "nocompass", "english", "tally", "gaeshi", 
"doryuuheki", "warn", "colorflags", "customtrashtalk", "details", 
"mip", "automap", "localecho", "combo", "rpon", "powerstruggle",
"override_save", "noemail", "dai", "leap", "ignoretactic", "256color",
"shunshin", "autodroploot", "clonefail", "autodoor", "clonedetails",
};

var Trap_flags=
[]string{
"room", "obj", "enter", "leave", "open", "close", "get", "put", "pick",
"unlock", "north", "south", "east", "west", "up", "down", "examine",
"northeast", "northwest", "southeast", "southwest", "r6", "r7", 
"r8", "r9", "r10", "r11", "r12", "r13","r14", "r15",
};

var Cmd_flags =
[]string{
 "possessed", "polymorphed", "watch", "fullname", "r3", "r4", "r5", "r6", "r7", "r8", 
 "r9", "r10", "r11", "r12", "r13", "r14", "r15", "r16", "r17", "r18", "r19", 
 "r20", "r21", "r22", "r23", "r24", "r25", "r26", "r27", "r28", "r29", "r30",
};

var Wear_locs=
[]string{
"light", "finger1", "finger2", "neck1", "neck2", "body", "head", "legs",
"feet", "hands", "arms", "shield", "about", "waist", "wrist1", "wrist2",
"wield", "hold", "dual_wield", "ears", "eyes", "missile_wield", "back",
"face", "ankle1", "ankle2", "hip",
};

var Ris_flags=
[]string{
"fire", "cold", "electricity", "energy", "blunt", "pierce", "slash", "acid",
"poison", "drain", "sleep", "charm", "hold", "nonmagic", "plus1", "plus2",
"plus3", "plus4", "plus5", "plus6", "magic", "paralysis", "r1", "r2", "r3",
"r4", "r5", "r6", "r7", "r8", "r9", "r10",
};

var Trig_flags=
[]string{
"up", "unlock", "lock", "d_north", "d_south", "d_east", "d_west", "d_up",
"d_down", "door", "container", "open", "close", "passage", "oload", "mload",
"teleport", "teleportall", "teleportplus", "death", "cast", "fakeblade",
"rand4", "rand6", "trapdoor", "anotherroom", "usedial", "absolutevnum",
"showroomdesc", "autoreturn", "r2", "r3",
};

var Part_flags=
[]string{
"head", "arms", "legs", "heart", "brains", "guts", "hands", "feet", "fingers",
"ear", "eye", "long_tongue", "eyestalks", "tentacles", "fins", "wings",
"tail", "scales", "claws", "fangs", "horns", "tusks", "tailattack",
"sharpscales", "beak", "haunches", "hooves", "paws", "forelegs", "feathers",
"r1", "r2",
};

var Attack_flags=
[]string{
"bite", "claws", "tail", "sting", "punch", "kick", "trip", "bash", "stun",
"gouge", "backstab", "feed", "drain", "firebreath", "frostbreath",
"acidbreath", "lightnbreath", "gasbreath", "poison", "nastypoison", "gaze",
"blindness", "causeserious", "earthquake", "causecritical", "curse",
"flamestrike", "harm", "fireball", "colorspray", "weaken", "spiralblast",
};

var Defense_flags=
[]string{
"parry", "dodge", "heal", "curelight", "cureserious", "curecritical",
"dispelmagic", "dispelevil", "sanctuary", "fireshield", "shockshield",
"shield", "bless", "stoneskin", "teleport", "monsum1", "monsum2", "monsum3",
"monsum4", "disarm", "iceshield", "grip", "truesight", "acidmist", "venomshield",
};

var Element_flags=
[]string{
"fire", "water", "wind", "earth", "lightning", "ice", "wood", "metal",
"max_element_type",
};

/*
 * Note: I put them all in one big set of flags since almost all of these
 * can be shared between mobs, objs and rooms for the exception of
 * bribe and hitprcnt, which will probably only be used on mobs.
 * ie: drop -- for an object, it would be triggered when that object is
 * dropped; -- for a room, it would be triggered when anything is dropped
 *          -- for a mob, it would be triggered when anything is dropped
 *
 * Something to consider: some of these triggers can be grouped together,
 * and differentiated by different arguments... for example:
 *  hour and time, rand and randiw, speech and speechiw
 * 
 */
var Mprog_flags=
[]string{
"act", "speech", "rand", "fight", "death", "hitprcnt", "entry", "greet",
"allgreet", "give", "bribe", "hour", "time", "wear", "remove", "sac",
"look", "exa", "zap", "get", "drop", "damage", "repair", "randiw",
"speechiw", "pull", "push", "sleep", "rest", "leave", "script", "use",
};

