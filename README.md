# CHAR-GEN

Char-Gen is a barebones character generator for *most* tabletop RPG systems, written in pure Go. 

Take a look at how it works in action over [HERE!](rpg.austinmorales.dev)

Want to work on your own front-end dev skills? Well you're in luck because *you* can self host this API if you want 🤠

To get this running on your own hardware, just make sure you have the latest version of Go installed and run `go build` in the same directory as `main.go`.

# Endpoints

`/character` - returns a JSON response with character stats and name. 

`/weapon` - this will generate a new enchanted/magic weapon.

`/npc` - returns a monster/NPC statblock


# Environment Variables
`$DB_URL` = the URL to the PostgreSQL DB

`$PORTNO` = the port number for the server to run on


# Features roadmap:

## 2.0:
- Export generated characters to a character sheet via LaTeX
- ~~Generate random enchanted weapons~~
- Generate random NPCs/Monsters from DB
- Move over to new domain name (dungeongenius.com)


## To Infinity (And Beyond...)
- Comission pixel art for generated items (funding needed)
