## Changelog
**1.1.0**
- cleaned main.go and prepeared for future modification

**1.1.0**
- Improved categories DB storage adding categories_parent_categories table Which gives the possibility to store the same categories for many different parent categories
- Added unique key constraints to the categories, attributes, attributes_values tables
- Created storing categories logick
- Created Bol parser

**1.0.0**
- Created Migration script and sql files
- Added config.example.json file which can be used like a template
- Created Makefile
- Added CI script
- Created Docker-compose file