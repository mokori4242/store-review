DB_FILE := sqlite.db
MIGRATION_DIR := migration

help:
	@echo "Available commands:"
	@echo "  init-db      - Initialize sqlite.db (remove existing and create new)"
	@echo "  create-tables - Create tables from migration files"
	@echo "  refresh      - Initialize DB and create all tables"
	@echo "  clean-db     - Remove sqlite.db file"
	@echo "  reset-db     - Clean and recreate database"

# Initialize database (create empty sqlite.db)
init-db:
	@echo "Initializing SQLite database..."
	@rm -f $(DB_FILE)
	@touch $(DB_FILE)
	@echo "Database $(DB_FILE) initialized."

# Create tables from migration files
create-tables:
	@echo "Creating tables from migration files..."
	@if [ ! -f $(DB_FILE) ]; then \
		echo "Database file not found. Running init-db first..."; \
		$(MAKE) init-db; \
	fi
	@for sql_file in $(MIGRATION_DIR)/*.sql; do \
		if [ -f "$$sql_file" ]; then \
			echo "Executing $$sql_file..."; \
			sqlite3 $(DB_FILE) < "$$sql_file"; \
		fi; \
	done
	@echo "All tables created successfully."

# Full migration (init + create tables)
refresh: init-db create-tables
	@echo "Database migration completed."

# Clean database
clean-db:
	@echo "Removing database file..."
	@rm -f $(DB_FILE)
	@echo "Database cleaned."

# Reset database (clean + migrate)
reset-db: clean-db refresh
	@echo "Database reset completed."
