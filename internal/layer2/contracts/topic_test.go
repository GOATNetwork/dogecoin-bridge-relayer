package contracts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopic(t *testing.T) {
	assert.Equal(t, AddressRegisteredTopic.String(), "0x0ab661710c67363885e0e51920050375aff9dcd587adf3e2e468e060ee8f0e1e")
	assert.Equal(t, TaskSubmittedTopic.String(), "0xc43825e1299c5fbb71e3a2b57d3a8d8178eaddc032e876d32dc37514160bd376")
	assert.Equal(t, TaskUpdatedTopic.String(), "0x3357e4d22dded01b8f5952481c84ccbcd9818d31adbb1bb1e8df6557f58e1ab7")
	assert.Equal(t, SubmitterRotationRequestedTopic.String(), "0x810bb46f7f5182d661c517393732ca0639393a548c222be3f52830dbd81b5584")
	assert.Equal(t, SubmitterChosenTopic.String(), "0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d")
	assert.Equal(t, ParticipantAddedTopic.String(), "0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b")
	assert.Equal(t, ParticipantRemovedTopic.String(), "0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc")
	assert.Equal(t, ParticipantsResetTopic.String(), "0x32e9d8d19fb1e71c8dc610e5f45fd7f1e2f81babf8ea90e267475a708e09c35e")
	assert.Equal(t, DepositRecordedTopic.String(), "0xb707b36bff7df581533b6b1939ae22c92c6fb817647b5a7bc7b7a42914dd7927")
	assert.Equal(t, WithdrawalRecordedTopic.String(), "0x59a01e2a8864dd3fae2368280a0cb84e97d4960b6c44fdde8fda335c715ce9e4")
	assert.Equal(t, NIP20TokenEventBurnbTopic.String(), "0xebe23dd93b970477278ceb9abd3082df92d977d6131fb0ef75f18c3d353b565a")
	assert.Equal(t, NIP20TokenEventMintbTopic.String(), "0x685c530c280ee1f7a4e96d082303ee9ebf21cec512259c6a943eda3854e05102")
	assert.Equal(t, AssetListedTopic.String(), "0x9a321153d23feb212bd45898d567f7472a2ecc6c752d8e6757ea0914ab2b7009")
	assert.Equal(t, AssetUpdatedTopic.String(), "0xc91c349180cb8d82f404c5b3cb776276676bba19867657b1e45e84e3103d0b36")
	assert.Equal(t, AssetDelistedTopic.String(), "0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6")
	assert.Equal(t, LinkTokenTopic.String(), "0x1dc913129df13b37a9f34933a8f82f346b89e0625048b9f1488c3d1e2af063c6")
	assert.Equal(t, ResetLinkedTokenTopic.String(), "0x1f036e69048c80897a67509c6ce12379bd2f893606404ec226c2db215f15eb32")
	assert.Equal(t, TokenSwitchTopic.String(), "0x744b3741ce814ba0b60e84403861ca5d5c3910c693b612f751af4eac581d6363")
}
