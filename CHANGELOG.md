## Changelog
**2.0.0**
- REST API integration
- Response meta block has next and previous links
- Added possibility to get all and get categories by id and the same for attributes

**1.2.1**
- Moved repository initialising part to the main file
- Created repository collection
- adjusted Makefile

**1.2.0**
- Added cobra comands to main file
- Adjusted migrations to the cobra commands
- Adjusted parsing script to the cobra commands

**1.1.1**
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