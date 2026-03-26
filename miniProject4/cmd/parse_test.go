package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestTimeStampParse(t *testing.T){
	data:="2024-02-03T10:15:30Z"
	r,err:=ParseTimeStamps(data)
	require.NoError(t,err)
	require.NotNil(t,r)
	// require.Nil(t,err)
	assert.Equal(t,2024,r.day.year)
	assert.Equal(t,02,r.day.month)
	assert.Equal(t,03,r.day.date)
	assert.Equal(t,10,r.time.hour)
	assert.Equal(t,15,r.time.min)
	assert.Equal(t,30,r.time.sec)
	assert.Equal(t,Done,r.State)

	data="2024-02-30T10:15:30Z"
	r,err=ParseTimeStamps(data)
	require.Error(t,err)
	fmt.Print(r)
	fmt.Print(err)
}