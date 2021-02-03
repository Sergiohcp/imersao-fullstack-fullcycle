package model_test

import (
	uuid "github.com/satori/go.uuid"
	"testing"

	"github.com/Sergiohcp/imersao-fullstack-fullcycle/domain/model"

	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T)  {
	name := "Sergio Henrique"
	email := "sergio@testeimersao.com"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser("", email)
	require.NotNil(t, err)
}
