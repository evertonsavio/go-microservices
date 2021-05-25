package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestTest(t *testing.T)  {
	
	request := CreateRepoResquest{
		Name: "golang repo by cli",
		Description: "A microservice thar creates repositories",
		Homepage: "https://github.com",
		Private: true,
		HasIssues: false,
		HasProjects: true,
		HasWiki: false,
	}

	//Marshal takes a input interface and attemps to create a valid json string

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))

	assert.EqualValues(t, `{"name":"golang repo by cli","description":"A microservice thar creates repositories","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`, string(bytes))
}