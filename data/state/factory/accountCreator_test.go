package factory_test

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-sandbox/data/mock"
	"github.com/ElrondNetwork/elrond-go-sandbox/data/state"
	"github.com/ElrondNetwork/elrond-go-sandbox/data/state/factory"
	"github.com/stretchr/testify/assert"
)

func TestAccountCreator_CreateAccountNilAddress(t *testing.T) {
	t.Parallel()

	shardC := mock.ShardCoordinatorMock{
		SelfID:     0,
		NrOfShards: 1,
	}
	accF, err := factory.NewAccountFactoryCreator(shardC)
	assert.Nil(t, err)

	_, ok := accF.(*factory.AccountCreator)
	assert.Equal(t, true, ok)

	acc, err := accF.CreateAccount(nil, &mock.AccountTrackerStub{})

	assert.Nil(t, acc)
	assert.Equal(t, err, state.ErrNilAddressContainer)
}

func TestAccountCreator_CreateAccountNilAccountTraccer(t *testing.T) {
	t.Parallel()

	shardC := mock.ShardCoordinatorMock{
		SelfID:     0,
		NrOfShards: 1,
	}
	accF, err := factory.NewAccountFactoryCreator(shardC)
	assert.Nil(t, err)

	_, ok := accF.(*factory.AccountCreator)
	assert.Equal(t, true, ok)

	acc, err := accF.CreateAccount(&mock.AddressMock{}, nil)

	assert.Nil(t, acc)
	assert.Equal(t, err, state.ErrNilAccountTracker)
}

func TestAccountCreator_CreateAccountOk(t *testing.T) {
	t.Parallel()

	shardC := mock.ShardCoordinatorMock{
		SelfID:     0,
		NrOfShards: 1,
	}
	accF, err := factory.NewAccountFactoryCreator(shardC)
	assert.Nil(t, err)

	_, ok := accF.(*factory.AccountCreator)
	assert.Equal(t, true, ok)

	acc, err := accF.CreateAccount(&mock.AddressMock{}, &mock.AccountTrackerStub{})

	assert.NotNil(t, acc)
	assert.Nil(t, err)
}