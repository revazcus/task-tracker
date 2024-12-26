package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/test/adapters/repository"

	"testing"
)

type UserUseCaseShould struct {
	suite.Suite
	userUseCase  usecase.UserUseCaseInterface
	userRepoMock repository.UserRepositoryMock
}

func TestUserUseCaseShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserUseCaseShould))
}

func (c *UserUseCaseShould) SetupTest() {
	//c.userUseCase = usecase.UserUseCase{}
	c.userRepoMock = repository.UserRepositoryMock{}
}

func (c *UserUseCaseShould) TestGivenUserDoesNotExist_WhenCreateUserWithValidData_ThenSuccessfullyCreated() {
	assert.Fail(c.T(), "implement me")
}

func (c *UserUseCaseShould) TestGivenUserDoesNotExist_WhenCreateUserWithMissingRequiredFields_ThenValidationError() {
	assert.Fail(c.T(), "implement me")
}

func (c *UserUseCaseShould) TestGivenUserDoesNotExist_WhenCreateUserWithDuplicateEmail_ThenConflictError() {
	assert.Fail(c.T(), "implement me")
}

func (c *UserUseCaseShould) TestGivenUserExist_WhenGetUser_ThenSuccessfullyReturned() {
	assert.Fail(c.T(), "implement me")
}

func (c *UserUseCaseShould) TestGivenUserDoesNotExist_WhenGetUser_ThenNotFoundError() {
	assert.Fail(c.T(), "implement me")
}

func (c *UserUseCaseShould) TestGivenUserExist_WhenUpdateUserWithValidData_ThenSuccessfullyUpdated() {
	assert.Fail(c.T(), "implement me")
}

func (c *UserUseCaseShould) TestGivenUserExist_WhenUpdateUserWithInvalidEmail_ThenValidationError() {
	assert.Fail(c.T(), "implement me")
}

func (c *UserUseCaseShould) TestGivenUserDoesNotExist_WhenUpdateUser_ThenNotFoundError() {
	assert.Fail(c.T(), "implement me")
}

func (c *UserUseCaseShould) TestGivenUserExist_WhenDeleteUser_ThenSuccessfullyDeleted() {
	//err := c.userUseCase.DeleteUser(1)
	//assert.Nil(c.T(), err)
}

func (c *UserUseCaseShould) TestGivenUserDoesNotExist_WhenDeleteUser_ThenNotFoundError() {
	//err := c.userUseCase.DeleteUser(1)
	//assert.EqualValues(c.T(), "implement me", err)
}
