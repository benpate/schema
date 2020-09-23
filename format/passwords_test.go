package format

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasUpperCase(t *testing.T) {

	require.NotNil(t, HasUppercase("")("there are no uppercase"))

	require.Nil(t, HasUppercase("")("There is one uppercase"))
	require.Nil(t, HasUppercase("")("there is oNe uppercase"))
	require.Nil(t, HasUppercase("")("there is one uppercasE"))

	require.Nil(t, HasUppercase("1")("There is one uppercase"))
	require.Nil(t, HasUppercase("1")("there is oNe uppercase"))
	require.Nil(t, HasUppercase("1")("there is one uppercasE"))

	require.NotNil(t, HasUppercase("2")("There is one uppercase"))
	require.NotNil(t, HasUppercase("2")("there is oNe uppercase"))
	require.NotNil(t, HasUppercase("2")("there is one uppercasE"))

	require.Nil(t, HasUppercase("4")("THEre Is one uppercase"))
	require.Nil(t, HasUppercase("4")("tHEre Is oNe uppercase"))
	require.Nil(t, HasUppercase("4")("tHEre Is one uppercasE"))
}

func TestHasLowerCase(t *testing.T) {

	require.NotNil(t, HasLowercase("")("THERE ARE NO LOWERCASE"))

	require.Nil(t, HasLowercase("")("tHERE IS ONE LOWERCASE"))
	require.Nil(t, HasLowercase("")("THERE IS oNE LOWERCASE"))
	require.Nil(t, HasLowercase("")("THERE IS ONE LOWERCASe"))

	require.Nil(t, HasLowercase("1")("tHERE IS ONE LOWERCASE"))
	require.Nil(t, HasLowercase("1")("THERE IS OnE LOWERCASE"))
	require.Nil(t, HasLowercase("1")("THERE IS ONE LOWERCASe"))

	require.Nil(t, HasLowercase("2")("tHERE aRE TWO LOWERCASE"))
	require.Nil(t, HasLowercase("2")("THERE aRE TWO LOWERCASe"))
	require.Nil(t, HasLowercase("2")("THERe ARE TWO lOWERCASE"))
	require.Nil(t, HasLowercase("2")("tHERE ARE TWO LOWERCASe"))

	require.NotNil(t, HasLowercase("4")("tHERE aRE TWO LOWERCASE"))
	require.NotNil(t, HasLowercase("4")("THERE aRE TWO LOWERCASe"))
	require.NotNil(t, HasLowercase("4")("THERe ARE TWO lOWERCASE"))
	require.NotNil(t, HasLowercase("4")("tHERE ARE TWO LOWERCASe"))
}
