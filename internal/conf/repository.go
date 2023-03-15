package conf

import "cuckoo/internal/adapter/repository"

func GetRepositoryOpt() repository.RepositoryOpt {
	return repositoryOpt
}
