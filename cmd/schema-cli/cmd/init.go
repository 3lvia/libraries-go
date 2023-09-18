package cmd

var (
	url              string
	subject          string
	system           string
	vaultToken       string
	storageFolder    string
	generationFolder string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", defaultVaultURL, "Vault URL")
	rootCmd.PersistentFlags().StringVar(&system, "system", "", "System")
	rootCmd.PersistentFlags().StringVar(&vaultToken, "vaultToken", "", "Vault token")
	rootCmd.PersistentFlags().StringVar(&subject, "subject", "", "Subject")

	rootCmd.PersistentFlags().StringVar(&storageFolder, "storageFolder", "", "Storage folder")
	rootCmd.PersistentFlags().StringVar(&generationFolder, "generationFolder", "", "Generation folder")
}
