package service

import (
	"context"
	"errors"

	"github.com/OzkrOssa/radiusx-users/internal/core/domain"
	"github.com/OzkrOssa/radiusx-users/internal/core/port"
	"github.com/OzkrOssa/radiusx-users/internal/core/utils"
)

type UserService struct {
	repo  port.UserRepository
	cache port.CacheRepository
}

func NewUserService(repo port.UserRepository, cache port.CacheRepository) *UserService {
	return &UserService{repo, cache}
}

func (u UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	user.Password = hashedPassword

	user, err = u.repo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, domain.ErrorConflictData) {
			return nil, err
		}
		return nil, domain.ErrorInternal
	}

	key := utils.GenerateCacheKey("user", user.ID)

	serializedUser, err := utils.Serialize(user)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	err = u.cache.Set(ctx, key, serializedUser, 0)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	err = u.cache.DeleteByPrefix(ctx, "users:*")
	if err != nil {
		return nil, domain.ErrorInternal
	}

	return user, nil
}

func (u UserService) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	var user *domain.User
	cacheKey := utils.GenerateCacheKey("user", id)
	cachedUser, err := u.cache.Get(ctx, cacheKey)

	if err == nil {
		err := utils.Deserialize(cachedUser, &user)
		if err != nil {
			return nil, domain.ErrorInternal
		}
		return user, nil
	}

	user, err = u.repo.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorDataNotFound) {
			return nil, err
		}
		return nil, domain.ErrorInternal
	}

	userSerialized, err := utils.Serialize(user)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	err = u.cache.Set(ctx, cacheKey, userSerialized, 0)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	return user, nil

}

func (u UserService) ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error) {

	var users []domain.User

	params := utils.GenerateCacheKeyParams(skip, limit)
	cacheKey := utils.GenerateCacheKey("users", params)

	cachedUsers, err := u.cache.Get(ctx, cacheKey)
	if err == nil {
		err := utils.Deserialize(cachedUsers, &users)
		if err != nil {
			return nil, domain.ErrorInternal
		}
		return users, nil
	}

	users, err = u.repo.ListUsers(ctx, skip, limit)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	usersSerialized, err := utils.Serialize(users)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	err = u.cache.Set(ctx, cacheKey, usersSerialized, 0)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	return users, nil
}

func (u UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	existingUser, err := u.repo.GetUserById(ctx, user.ID)

	if err != nil {
		if errors.Is(err, domain.ErrorDataNotFound) {
			return nil, err
		}
		return nil, domain.ErrorInternal
	}

	emptyData := user.Name == "" &&
		user.Email == "" &&
		user.Password == ""

	sameData := existingUser.Name == user.Name &&
		existingUser.Email == user.Email

	if emptyData || sameData {
		return nil, domain.ErrorNoUpdatedData
	}

	var hashedPassword string

	if user.Password != "" {
		hashedPassword, err = utils.HashPassword(user.Password)
		if err != nil {
			return nil, domain.ErrorConflictData
		}
	}

	user.Password = hashedPassword

	_, err = u.repo.UpdateUser(ctx, user)
	if err != nil {
		if errors.Is(err, domain.ErrorConflictData) {
			return nil, err
		}
		return nil, domain.ErrorInternal
	}

	cacheKey := utils.GenerateCacheKey("user", user.ID)

	err = u.cache.Delete(ctx, cacheKey)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	userSerialized, err := utils.Serialize(user)
	if err != nil {
		return nil, err
	}

	err = u.cache.Set(ctx, cacheKey, userSerialized, 0)
	if err != nil {
		return nil, domain.ErrorInternal
	}

	err = u.cache.DeleteByPrefix(ctx, "users:*")
	if err != nil {
		return nil, domain.ErrorInternal
	}

	return user, nil
}

func (u UserService) DeleteUser(ctx context.Context, id uint64) error {
	_, err := u.repo.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrorDataNotFound) {
			return err
		}
		return domain.ErrorInternal
	}

	cacheKey := utils.GenerateCacheKey("user", id)

	err = u.cache.Delete(ctx, cacheKey)
	if err != nil {
		return domain.ErrorInternal
	}

	err = u.cache.DeleteByPrefix(ctx, "users:*")
	if err != nil {
		return domain.ErrorInternal
	}

	return u.repo.DeleteUser(ctx, id)
}
