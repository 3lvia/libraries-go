package cmd

var (
	url              string
	topic            string
	system           string
	vaultToken       string
	storageFolder    string
	generationFolder string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", defaultVaultURL, "Vault URL")
	rootCmd.PersistentFlags().StringVar(&system, "system", "", "System is the name of the system owning the topic.")
	rootCmd.PersistentFlags().StringVar(&vaultToken, "vaultToken", "", "Vault token gives direct access to vault.")
	rootCmd.PersistentFlags().StringVar(&topic, "topic", "", "Topic is a topic in COnfluent cloud.")

	rootCmd.PersistentFlags().StringVar(&storageFolder, "storageFolder", "", "Storage folder is the folder where schemas are stored.")
	rootCmd.PersistentFlags().StringVar(&generationFolder, "generationFolder", "", "Generation folder")
}
