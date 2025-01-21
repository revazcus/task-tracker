package initService

import (
	"context"
	userRepo "user-service/adapters/repository/user"
	repositoryInterface "user-service/boundary/repository"
)

type Repositories struct {
	UserRepository repositoryInterface.UserRepository

	initableRepositories []InitableRepository
}

type InitableRepository interface {
	Init(ctx context.Context) error
}

type RepositoriesInit struct {
	dc *DependencyContainer

	repositories  *Repositories
	initFunctions []func(*DependencyContainer) error
}

func NewRepositoriesInit(dc *DependencyContainer) *RepositoriesInit {
	repositories := &Repositories{
		initableRepositories: make([]InitableRepository, 0),
	}
	return &RepositoriesInit{
		dc:           dc,
		repositories: repositories,
		initFunctions: []func(*DependencyContainer) error{
			repositories.initUserRepository,
		},
	}
}

func (i *RepositoriesInit) InitServices() error {
	for _, initFunc := range i.initFunctions {
		if err := initFunc(i.dc); err != nil {
			return err
		}
	}

	i.dc.Repositories = i.repositories
	return nil
}

func (i *RepositoriesInit) StartServices() error {
	for _, initRepo := range i.dc.Repositories.initableRepositories {
		if err := initRepo.Init(context.Background()); err != nil {
			i.dc.LogError(err)
			return err
		}
	}
	return nil
}

func (i *Repositories) initUserRepository(dc *DependencyContainer) error {
	collection, err := dc.EnvRegistry.GetEnv(MongoCollectionEnv)
	if err != nil {
		dc.LogError(err)
		return err
	}

	userRepo, err := userRepo.NewBuilder().
		Collection(collection).
		MongoRepo(dc.MongoRepo).
		Logger(dc.Logger).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.UserRepository = userRepo
	i.initableRepositories = append(i.initableRepositories, userRepo)

	return nil
}
