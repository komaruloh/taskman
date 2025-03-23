/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the task manager",
	Long: `Initialize the task manager. This command will initialize database and
configuration. Usage example:

taskman init`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		initAll()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initAll() {
	if err := initConfig(); err != nil {
		fmt.Println("Error initializing config", err)
		os.Exit(1)
	}

	if err := initDatabase(); err != nil {
		fmt.Println("Error initializing database", err)
		os.Exit(1)
	}
}

// check and create config directory, if needed
func checkPath(path string) error {
	if path == "" {
		return fmt.Errorf("Path is empty")
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Attempting to create directory: %s\n", path)
			err := os.Mkdir(path, 0755)
			if err != nil {
				return fmt.Errorf("Error creating directory: %s. Err: %v", path, err)
			}
			return nil
		}
		return fmt.Errorf("Error checking path: %s. Err: %v", path, err)
	}
	return nil
}
func defaultConfig() {
	viper.SetDefault("app.base_dir", "$HOME/.taskman")

	// database configuration
	viper.SetDefault("database.filename", "taskman.db")
}

func initDatabase() error {
	// Initialize database
	dbPath := viper.GetString("app.base_dir")
	dbPath = os.ExpandEnv(dbPath)
	fmt.Printf("Database path: %s\n", dbPath)

	err := checkPath(dbPath)
	if err != nil {
		return err
	}

	dbPath = dbPath + "/" + viper.GetString("database.filename")
	fmt.Printf("Database file: %s\n", dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	query := `CREATE TABLE 'tasks' (
		'id' INTEGER PRIMARY KEY AUTOINCREMENT, 
		'parent_id' VARCHAR(255) NULL, 
		'user_id' VARCHAR(64) NULL, 
		'status' VARCHAR(10) NOT NULL, 
		'title' VARCHAR(255) NOT NULL, 
		'detail' TEXT NULL, 
		'created_at' TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		'created_by' INTEGER NOT NULL, 
		'modified_at' TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		'modified_by' INTEGER NULL
);`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func initConfig() error {
	// Initialize configuration
	configPath := "$HOME/.config/taskman"
	configPath = os.ExpandEnv(configPath)

	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := checkPath(configPath)
	if err != nil {
		return err
	}

	err = viper.ReadInConfig()
	if err != nil && !os.IsNotExist(err) {
		defaultConfig()
		if err := viper.WriteConfigAs(configPath + "/config.yaml"); err != nil {
			return err
		}
	}

	return nil
}
