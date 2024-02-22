package src

import (
	"log"
	"task-test/core/config"
	task_repo "task-test/src/repositories/task"
	user_repo "task-test/src/repositories/user"
	task_service "task-test/src/services/task"
	user_service "task-test/src/services/user"
)

type repositories struct {
	UserRepo *user_repo.UserRepository
	TaskRepo *task_repo.TaskRepository
}

type services struct {
	UserSVC *user_service.UserService
	TaskSVC *task_service.TaskService
}

type Dependency struct {
	Repositories *repositories
	Services     *services
}

func initRepositories() *repositories {
	var r repositories
	var err error

	r.UserRepo, err = user_repo.NewUserRepository(config.DB(), config.Config())
	if err != nil {
		log.Panic(err)
	}

	r.TaskRepo, err = task_repo.NewTaskRepository(config.DB())
	if err != nil {
		log.Panic(err)
	}

	return &r
}

func initServices(r *repositories) *services {

	return &services{
		UserSVC: user_service.NewUserService(r.UserRepo),
		TaskSVC: task_service.NewTaskService(config.DB(), r.TaskRepo),
	}
}

func Dependencies() *Dependency {
	repositories := initRepositories()
	services := initServices(repositories)

	return &Dependency{
		Repositories: repositories,
		Services:     services,
	}
}
