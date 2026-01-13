package domain

type DashboardSummary struct {
	Orders []ExternalOrder
	Wallet ExternalWallet
	User ExternalUser
}
