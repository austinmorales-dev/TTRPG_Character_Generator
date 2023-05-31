# CHAR-GEN

Char-Gen is a barebones character generator for *most* tabletop RPG systems, written in pure Go. 

Take a look at how it works in action over [HERE!](rpg.austinmorales.dev)

Want to work on your own front-end dev skills? Well you're in luck because *you* can self host this API if you want 🤠

To get this running on your own hardware, just make sure you have the latest version of Go installed and run `go build` in the same directory as `main.go`.

It runs on port 9001 by default, but feel free to change it! 

As time goes on, expect more features! 

# Endpoints

`/character` - returns a JSON response with character stats and name. 

`/weapon` - this will generate a new enchanted/magic weapon.

# Features roadmap:
- Export generated characters to a character sheet via LaTeX (2.0)
- ~~Generate random enchanted weapons (2.0)~~
- Comission pixel art for generated items (funding needed)
- Move over to new domain name (dungeongenius.com)