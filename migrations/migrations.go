package config

import (
	"log"

	"gorm.io/gorm"
)

// Set up triggers for alias updates
func CreateTriggers(db *gorm.DB) {
	// Drop existing triggers
	err := db.Exec(`
        DROP TRIGGER IF EXISTS alias_insert_trigger ON aliases;
        DROP TRIGGER IF EXISTS alias_update_trigger ON aliases;
        DROP TRIGGER IF EXISTS alias_delete_trigger ON aliases;
    `).Error
	if err != nil {
		log.Fatalf("failed to drop existing triggers: %v", err)
	}

	// Drop existing function
	err = db.Exec(`
        DROP FUNCTION IF EXISTS update_user_alias_counts;
    `).Error
	if err != nil {
		log.Fatalf("failed to drop existing function: %v", err)
	}

	err = db.Exec(`
        CREATE OR REPLACE FUNCTION update_user_alias_counts()
        RETURNS TRIGGER AS $$
        BEGIN
            UPDATE users
            SET num_aliases = (SELECT COUNT(*) FROM aliases WHERE user_id = NEW.user_id)
            WHERE id = NEW.user_id;

            UPDATE users
            SET num_undeleted_aliases = (SELECT COUNT(*) FROM aliases WHERE user_id = NEW.user_id AND is_deleted = FALSE)
            WHERE id = NEW.user_id;

            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
    `).Error
	if err != nil {
		log.Fatalf("failed to create trigger function: %v", err)
	}

	err = db.Exec(`
        CREATE TRIGGER alias_insert_trigger
        AFTER INSERT ON aliases
        FOR EACH ROW
        EXECUTE FUNCTION update_user_alias_counts();

        CREATE TRIGGER alias_update_trigger
        AFTER UPDATE ON aliases
        FOR EACH ROW
        EXECUTE FUNCTION update_user_alias_counts();

        CREATE TRIGGER alias_delete_trigger
        AFTER DELETE ON aliases
        FOR EACH ROW
        EXECUTE FUNCTION update_user_alias_counts();
    `).Error
	if err != nil {
		log.Fatalf("failed to create triggers: %v", err)
	}

    log.Println("Triggers created")
}
