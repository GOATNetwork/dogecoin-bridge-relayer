package layer2

const (
	SubmitterChosen            = "SubmitterChosen"
	SubmitterRotationRequested = "SubmitterRotationRequested"
	TaskSubmitted              = "TaskSubmitted"
	TaskUpdated                = "TaskUpdated"
	AddressRegistered          = "AddressRegistered"
	ParticipantAdded           = "ParticipantAdded"
	ParticipantRemoved         = "ParticipantRemoved"
	ParticipantReset           = "ParticipantReset"
	DepositRecorded            = "DepositRecorded"
	WithdrawalRecorded         = "WithdrawalRecorded"

	// #nosec G101: This is not a credential
	NIP20TokenMintbEvent = "NIP20TokenEvent_mintb"
	// #nosec G101: This is not a credential
	NIP20TokenBurnbEvent = "NIP20TokenEvent_burnb"

	AssetListed   = "AssetListed"
	AssetUpdated  = "AssetUpdated"
	AssetDelisted = "AssetDelisted"
)
