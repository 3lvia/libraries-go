package cmd

var (
	url              string
	topic            string
	system           string
	vaultToken       string
	storageFolder    string
	generationFolder string
	schemaID         int
)

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", defaultVaultURL, "Vault URL")
	rootCmd.PersistentFlags().StringVar(&system, "system", "", "System is the name of the system owning the topic.")
	rootCmd.PersistentFlags().StringVar(&vaultToken, "vaultToken", "", "Vault token gives direct access to vault.")
	rootCmd.PersistentFlags().StringVar(&topic, "topic", "", "Topic is a topic in Confluent cloud.")
	rootCmd.PersistentFlags().IntVar(&schemaID, "schemaID", 0, "Schema ID is the ID of the schema to be generated.")

	rootCmd.PersistentFlags().StringVar(&storageFolder, "storageFolder", "", "Storage folder is the folder where schemas are stored.")
	rootCmd.PersistentFlags().StringVar(&generationFolder, "generationFolder", "", "Generation folder")
}
