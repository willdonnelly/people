package people

import (
	"fmt"
	"math/rand"
)

// Select a random word and append some numbers to it to make something username-looking.
func Username() string {
	word := Words[rand.Int31n(int32(len(Words)))]
	digits := rand.Int31n(1000)
	return fmt.Sprintf("%s%d", word, digits)
}

var Words = []string{
	"acquirable",
	"bestsellers",
	"farther",
	"prizer",
	"shasta",
	"evaporate",
	"auspices",
	"garments",
	"partnership",
	"blocs",
	"forestalling",
	"razors",
	"extensibility",
	"unavoidably",
	"logician",
	"embroidered",
	"crippling",
	"supranational",
	"milton",
	"healthily",
	"spiraling",
	"coolies",
	"bartend",
	"precondition",
	"reflectors",
	"judged",
	"rinser",
	"amplify",
	"casseroles",
	"physics",
	"raider",
	"whippet",
	"expulsion",
	"enzyme",
	"prohibit",
	"gazers",
	"unchangeable",
	"matching",
	"mouthe",
	"millihenry",
	"plowshare",
	"quicken",
	"blackmailing",
	"chatham",
	"jobbing",
	"augustly",
	"constitutionality",
	"cathodes",
	"inspirations",
	"seniority",
	"staging",
	"figuratively",
	"beckon",
	"rankle",
	"buzzwords",
	"mccullough",
	"justifying",
	"antiquities",
	"ardency",
	"tribunals",
	"laughs",
	"shakes",
	"feedback",
	"balustrade",
	"mattress",
	"seduces",
	"attainments",
	"counterattack",
	"sweeter",
	"deforestation",
	"digests",
	"sacrificed",
	"scripts",
	"philharmonic",
	"legerdemain",
	"advancements",
	"disburse",
	"bottles",
	"scatterbrain",
	"conceptions",
	"planer",
	"fishpond",
	"tidying",
	"illustration",
	"dishonoring",
	"impostors",
	"aspect",
	"summations",
	"steering",
	"cheesy",
	"hamlets",
	"cryptanalyst",
	"ensued",
	"upholsterer",
	"detaining",
	"penned",
	"robbers",
	"contingency",
	"effectively",
	"soybean",
	"clockings",
	"pappas",
	"jellies",
	"formulae",
	"routines",
	"savoyard",
	"redefining",
	"insistently",
	"macroscopic",
	"taster",
	"phosphates",
	"midsts",
	"invertebrates",
	"vices",
	"vacancy",
	"predominated",
	"timeshare",
	"convincing",
	"paralleling",
	"conceived",
	"guggenheim",
	"paintings",
	"dispells",
	"incapacitating",
	"nostrand",
	"pliant",
	"sleuth",
	"grammar",
	"wallows",
	"dismisses",
	"wilhelm",
	"exiling",
	"checkers",
	"proceedings",
	"hoarsely",
	"stretches",
	"purport",
	"limousine",
	"inheritresses",
	"company",
	"thruway",
	"hopkinsian",
	"downcast",
	"dangers",
	"anatomically",
	"allure",
	"stampers",
	"executive",
	"postmaster",
	"depressing",
	"dragons",
	"countys",
	"harriet",
	"attire",
	"runway",
	"bubbled",
	"waterman",
	"gerhardt",
	"honorableness",
	"flurry",
	"refract",
	"bacteria",
	"antiques",
	"provide",
	"mysteriously",
	"interrogation",
	"discontinuous",
	"victrola",
	"replications",
	"passion",
	"thawed",
	"alligator",
	"documentaries",
	"nakedness",
	"veining",
	"durability",
	"corrosion",
	"laterally",
	"winnipeg",
	"federally",
	"divest",
	"gasped",
	"unselfishly",
	"disclosing",
	"nurturing",
	"tramway",
	"palmed",
	"disruptions",
	"footman",
	"senators",
	"cleave",
	"effected",
	"ceramic",
	"leathery",
	"nicely",
	"frustrater",
	"warning",
	"lexicons",
	"exactions",
	"prover",
	"recreates",
	"puddling",
	"diabolic",
	"spatula",
	"herons",
	"blobs",
	"fibrosity",
	"cabinetmake",
	"phobic",
	"jingling",
	"double",
	"proving",
	"taipei",
	"skims",
	"prophesied",
	"hastily",
	"parasitics",
	"landings",
	"taxicabs",
	"subway",
	"recount",
	"noisemake",
	"induce",
	"mountaineer",
	"achieved",
	"celebrities",
	"fluffy",
	"bimini",
	"briefcases",
	"devote",
	"stylishly",
	"cleansing",
	"disclaimed",
	"phonemes",
	"impertinent",
	"connecting",
	"lentil",
	"revelations",
	"phoned",
	"lading",
	"lengthens",
	"nobles",
	"despairing",
	"hatchets",
	"livably",
	"lodger",
	"tokens",
	"ensurers",
	"interconnects",
	"passionate",
	"peppergrass",
	"bookkeep",
	"humerus",
	"thanklessness",
	"shamed",
	"choreography",
	"swimmers",
	"authors",
	"football",
	"auditions",
	"greener",
	"deflater",
	"tariff",
	"banjos",
	"packages",
	"gambit",
	"heated",
	"interfere",
	"collectors",
	"sideboards",
	"shoreline",
	"rutherford",
	"ethnology",
	"persecuting",
	"operatives",
	"demark",
	"curtate",
	"inheritress",
	"economizer",
	"pleural",
	"broiling",
	"minting",
	"ricochet",
	"lookup",
	"biases",
	"auctioneers",
	"formula",
	"morphism",
	"outstripped",
	"falsifying",
	"fealty",
	"homesteads",
	"dilate",
	"councilmen",
	"cornea",
	"intercept",
	"adjoins",
	"medals",
	"autonomic",
	"monologue",
	"cruisers",
	"psychoanalyst",
	"registrations",
	"agnostics",
	"ambivalently",
	"punishable",
	"philosophically",
	"storages",
	"wistful",
	"loveland",
	"preferential",
	"armchairs",
	"washington",
	"accretions",
	"interchangeable",
	"ambitions",
	"hostesss",
	"heading",
	"crucifies",
	"venturesome",
	"mullion",
	"fueling",
	"bedposts",
	"soapstone",
	"garland",
	"heaved",
	"instrumentalists",
	"patristic",
	"tableau",
	"plagiarist",
	"disambiguate",
	"autopilot",
	"anointing",
	"retypes",
	"pirates",
	"obfuscatory",
	"octennial",
	"indeterminately",
	"defended",
	"childbirth",
	"liberation",
	"kilograms",
	"elaborates",
	"snyaptic",
	"granitic",
	"carthage",
	"deteriorate",
	"matilda",
	"antislavery",
	"batter",
	"cringes",
	"aerosolize",
	"floppily",
	"caribbean",
	"woodbury",
	"wrapper",
	"capistrano",
	"meats",
	"overdrafts",
	"gnats",
	"sympathetic",
	"pritchard",
	"subscripted",
	"chinquapin",
	"skater",
	"counterfeiter",
	"leathern",
	"tabula",
	"bowled",
	"reagan",
	"appropriators",
	"curing",
	"pacific",
	"scandalous",
	"anesthetized",
	"reinforcements",
	"conner",
	"complains",
	"conjugal",
	"enumerator",
	"inconclusive",
	"pipelines",
	"synthesizer",
	"intimate",
	"saturater",
	"splintered",
	"taxonomy",
	"roaring",
	"transduction",
	"collegial",
	"breakdown",
	"adducing",
	"debenture",
	"jeopardy",
	"intoxicant",
	"rescue",
	"phrased",
	"cartwheel",
	"remedies",
	"penguin",
	"shined",
	"codification",
	"impugn",
	"doorbell",
	"ludlow",
	"visibility",
	"agglutinins",
	"apposition",
	"pathogenic",
	"bestial",
	"present",
	"encyclopedic",
	"qualifiers",
	"realists",
	"baptism",
	"plasticity",
	"transitioned",
	"atalanta",
	"crucially",
	"trackers",
	"identities",
	"cursors",
	"backspace",
	"airships",
	"multilevel",
	"concretely",
	"gazette",
	"intelligibility",
	"cottager",
	"denigrated",
	"unimpeded",
	"matisse",
	"thrashed",
	"impious",
	"ceaseless",
	"callisto",
	"lollipop",
	"defenestrated",
	"reredos",
	"chemic",
	"foulest",
	"solemn",
	"staley",
	"ballfield",
	"alameda",
	"panaceas",
	"nabisco",
	"strainer",
	"hackmatack",
	"hemispheric",
	"cogitated",
	"customizing",
	"pushbutton",
	"dressmaker",
	"amending",
	"penance",
	"seasonal",
	"chromium",
	"offsaddle",
	"atrophy",
	"souffle",
	"platforms",
	"wrangle",
	"clearness",
	"anecdotes",
	"hurting",
	"tooled",
	"angora",
	"narrate",
	"statistician",
	"philosoph",
	"assertions",
	"indefinitely",
	"parsimonious",
	"bribing",
	"tolerant",
	"lilies",
	"sulfate",
	"righteously",
	"stereotypical",
	"degeneracy",
	"similarity",
	"pastimes",
	"informed",
	"polypropylene",
	"backlog",
	"typography",
	"survivors",
	"reconfiguring",
	"gadding",
	"caryatid",
	"scuttling",
	"semaphores",
	"debugged",
	"pacification",
	"carbone",
	"firearms",
	"neurophysiology",
	"blazing",
	"ballrooms",
	"thunderbolts",
	"forefather",
	"rachel",
	"collision",
	"reticulately",
	"resignations",
	"interactions",
	"conspirator",
	"basilar",
	"climaxes",
	"draining",
	"cabinets",
	"checksumming",
	"suicide",
	"coffees",
	"mescaline",
	"tininess",
	"tinder",
	"binomial",
	"berates",
	"cashed",
	"bellwethers",
	"carbonation",
	"kalamazoo",
	"thyroglobulin",
	"kidnappers",
	"numbed",
	"shiftiness",
	"presuming",
	"achievements",
	"amplifiers",
	"lurches",
	"cataclysmic",
	"subvert",
	"paragon",
	"hoppers",
	"lapels",
	"recast",
	"pitilessly",
	"coffins",
	"outstretched",
	"perceiving",
	"thoughtfully",
	"taking",
	"stems",
	"favors",
	"streets",
	"quieting",
	"monoid",
	"delectable",
	"encoding",
	"jejune",
	"sincere",
	"goober",
	"testes",
	"lexicon",
	"richter",
	"covenants",
	"pitiers",
	"quintessence",
	"yellower",
	"equitably",
	"dickens",
	"contentment",
	"bessemer",
	"metabole",
	"timetables",
	"solemnity",
	"report",
	"indiana",
	"fortunate",
	"sweepstake",
	"lapelled",
	"arduousness",
	"blunts",
	"anorthosite",
	"acclimatized",
	"potters",
	"babysitter",
	"graveyard",
	"forthcoming",
	"glimmer",
	"knaves",
	"purposed",
	"entice",
	"amorality",
	"poetics",
	"frightened",
	"dilution",
	"erastus",
	"anabaptists",
	"carport",
	"whatre",
	"harpsichord",
	"marvin",
	"triers",
	"dumbbells",
	"hopefulness",
	"sorting",
	"continentally",
	"asynchronism",
	"illustratively",
	"afforestation",
	"constitutional",
	"arcsin",
	"darlings",
	"removes",
	"incompletion",
	"bitterroot",
	"blissfully",
	"splash",
	"manfred",
	"rashly",
	"bustling",
	"hathaway",
	"lacerating",
	"underplayed",
	"roundhead",
	"purposefully",
	"baldly",
	"steadily",
	"syndromes",
	"subversion",
	"lunchtime",
	"congressman",
	"mouses",
	"valences",
	"perhaps",
	"sawmills",
	"pinehurst",
	"comparison",
	"expansive",
	"kidnappers",
	"occasioned",
	"transferable",
	"transducer",
	"synchrotron",
	"rutile",
	"effete",
	"savaged",
	"dearths",
	"reading",
	"horizons",
	"scabbards",
	"presences",
	"trinity",
	"isinglass",
	"abusive",
	"critics",
	"recalculates",
	"loitering",
	"atrium",
	"terrorizing",
	"unblocked",
	"trickery",
	"accomplices",
	"bleedings",
	"wholesomeness",
	"opaquely",
	"epitaph",
	"escorted",
	"automatically",
	"interdependence",
	"ludicrously",
	"necessary",
	"chastising",
	"beneficences",
	"mbabane",
	"floury",
	"iniquity",
	"craftsmen",
	"feasted",
	"inviting",
	"gasoline",
	"balsam",
	"finality",
	"howling",
	"largesse",
	"docile",
	"conveniences",
	"sensors",
	"dinners",
	"alterable",
	"overhangs",
	"satisfaction",
	"semicolons",
	"eclipses",
	"languish",
	"symbol",
	"praecox",
	"muskox",
	"inheritrices",
	"invade",
	"measurable",
	"converses",
	"enlivened",
	"tantrums",
	"stopped",
	"wavefront",
	"constance",
	"humanness",
	"postscripts",
	"troublesomely",
	"unfavorable",
	"tarbell",
	"antagonistically",
	"unsure",
	"bridgehead",
	"imprints",
	"neurons",
	"volatilities",
	"renowned",
	"midwestern",
	"mistakable",
	"recover",
	"stoichiometry",
	"hating",
	"feuds",
	"superscripts",
	"irvine",
	"vertebrates",
	"fifteenth",
	"parthenon",
	"southbound",
	"nineveh",
	"twitched",
	"flooded",
	"sectors",
	"tractable",
	"subtrees",
	"suffering",
	"doggedness",
	"shoveled",
	"poisons",
	"cinders",
	"payoff",
	"enjoyably",
	"applicative",
	"aspersion",
	"analytic",
	"tactician",
	"forgiveness",
	"shibboleth",
	"equalizing",
	"deluxe",
	"frankest",
	"epistemology",
	"donated",
	"armpits",
	"corpsmen",
	"pyhrric",
	"backers",
	"clearings",
	"patience",
	"bijectively",
	"screwbean",
	"careers",
	"delaney",
	"applications",
	"autocrats",
	"homeopath",
	"abysmal",
	"ribbon",
	"bimodal",
	"relaxations",
	"fascism",
	"chatterer",
	"inaugurated",
	"uncaught",
	"hardhat",
	"voiceband",
	"wetland",
	"additions",
	"yugoslav",
	"atrophying",
	"plebeian",
	"duplication",
	"birthplaces",
	"farmers",
	"suspends",
	"retrograde",
	"transgress",
	"harassing",
	"yonder",
	"formulators",
	"alhambra",
	"attesting",
	"fairport",
	"osseous",
	"oratorical",
	"songbook",
	"wingspan",
	"loggers",
	"julies",
	"highlighted",
	"granter",
	"louver",
	"seersucker",
	"parley",
	"beheading",
	"soothes",
	"crises",
	"rounded",
	"lesser",
	"streamline",
	"measure",
	"rusticate",
	"ordained",
	"gabrielle",
	"haplessness",
	"raccoon",
	"explosions",
	"freakish",
	"soother",
	"pravda",
	"icebox",
	"homing",
	"leeway",
	"handsomeness",
	"commodity",
	"liberated",
	"playwrights",
	"rebroadcast",
	"instincts",
	"bumblers",
	"variously",
	"martyrs",
	"wednesdays",
	"thresholds",
	"azimuths",
	"aspire",
	"touchdown",
	"drudge",
	"teared",
	"exposing",
	"downside",
	"lofts",
	"lebensraum",
	"nether",
	"nozzle",
	"lawsuits",
	"gustafson",
	"predates",
	"meritorious",
	"searchings",
	"bundling",
	"scenarios",
	"schmidt",
	"fischbein",
	"bewitched",
	"grassiest",
	"disbursed",
	"indestructible",
	"burning",
	"swearer",
	"spirits",
	"cushman",
	"timbre",
	"seconding",
	"giggle",
	"ordered",
	"bibliophile",
	"affianced",
	"camera",
	"bernet",
	"facilitated",
	"highways",
	"grandly",
	"inkings",
	"shelves",
	"staggering",
	"bloomington",
	"magnesite",
	"colored",
	"sunned",
	"audiometer",
	"affectation",
	"maiden",
	"subfields",
	"axiomatically",
	"respire",
	"palmate",
	"domino",
	"depressed",
	"industrials",
	"expellable",
	"justine",
	"nuclei",
	"balkanize",
	"refreshments",
	"santiago",
	"organizers",
	"prospections",
	"antagonizing",
	"resort",
	"incoherent",
	"nonsense",
	"highwayman",
	"members",
	"sightings",
	"unique",
	"relentlessness",
	"chorused",
	"recipes",
	"strongest",
	"airlocks",
	"brethren",
	"legate",
	"capitalist",
	"digestive",
	"bookcases",
	"controller",
	"deficit",
	"masochism",
	"unlocks",
	"recreated",
	"bulbs",
	"avertive",
	"penetrable",
	"mynheer",
	"attracts",
	"axiomatizes",
	"typical",
	"resuming",
	"remarkable",
	"cookery",
	"excused",
	"meals",
	"bodybuilders",
	"canoes",
	"harmonize",
	"uncertainly",
	"blowup",
	"former",
	"squaring",
	"alphabetical",
	"disciplines",
	"genera",
	"melody",
	"cossack",
	"varyings",
	"abridge",
	"harvested",
	"tenney",
	"zodiacal",
	"labrador",
	"honoraries",
	"heuristics",
	"billeted",
	"outlays",
	"scrimmage",
	"folksy",
	"resented",
	"audits",
	"pilgrims",
	"compromisers",
	"niggardly",
	"patchwork",
	"briefest",
	"comatose",
	"faceplate",
	"variations",
	"assertively",
	"tutelage",
	"precedences",
	"bleating",
	"suspensor",
	"divisors",
	"youngest",
	"ruthlessly",
	"breadth",
	"abundance",
	"penalized",
	"promotes",
	"thermometers",
	"retired",
	"circumlocutions",
	"recites",
	"nettlesome",
	"woodside",
	"differentials",
	"handicaps",
	"lawmen",
	"budging",
	"balkanized",
	"instantiated",
	"contradict",
	"insinuations",
	"assigned",
	"budgeter",
	"hailing",
	"strolls",
	"ultrasound",
	"liberties",
	"pantomime",
	"triplex",
	"inroad",
	"befouling",
	"foregoing",
	"gullys",
}
