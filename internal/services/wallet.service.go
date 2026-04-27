func CreateDefaultWallets(userID primitive.ObjectID) []Wallet {
	now := time.Now()

	return []Wallet{
		{
			UserID:     userID,
			WalletType: WalletMoniFlex,
			WalletName: "MoniFlex Wallet",
			Balance:    0,
			Status:     WalletActive,
			CreatedAt:  now,
		},
		{
			UserID:     userID,
			WalletType: WalletMoniBank,
			WalletName: "MoniBank Wallet",
			Balance:    0,
			Status:     WalletActive,
			CreatedAt:  now,
		},
		{
			UserID:     userID,
			WalletType: WalletMoniTarget,
			WalletName: "MoniTarget Wallet",
			Balance:    0,
			Status:     WalletActive,
			CreatedAt:  now,
		},
	}
}