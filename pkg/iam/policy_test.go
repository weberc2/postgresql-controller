package iam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	region     = "eu-west-1"
	accountID  = "000000000000"
	rolePrefix = "iam_developer_"
)

func Test_AddUsersToDocument(t *testing.T) {

	assert := assert.New(t)

	document := NewPolicyDocument("2012-10-17")
	document.Add(region, accountID, rolePrefix, "user1")
	document.Add(region, accountID, rolePrefix, "user2")

	assert.Equal(2, document.Count())
	assert.True(document.Exists("user1"))
	assert.True(document.Exists("user2"))
}

func Test_RemoveUsersFromDocument(t *testing.T) {

	assert := assert.New(t)

	document := NewPolicyDocument("2012-10-17")
	document.Add(region, accountID, rolePrefix, "user1")
	document.Add(region, accountID, rolePrefix, "user2")
	document.Add(region, accountID, rolePrefix, "user3")
	document.Remove("user2")

	assert.Equal(2, document.Count())
	assert.True(document.Exists("user1"))
	assert.False(document.Exists("user2"))
	assert.True(document.Exists("user3"))
}
