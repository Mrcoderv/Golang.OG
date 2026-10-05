# SOLID Principle in Go

SOLID is an approach to how we **build and organize a software system**.

It has 5 main principles:

### 1. S — Single Responsibility

Give each component **one main responsibility**.

> One component → One main job.

### 2. O — Open/Closed

When adding a new feature, the existing working code should **not be unnecessarily modified**.

> Add new behavior → Don't break old behavior.

### 3. L — Liskov Substitution

If we replace an old implementation with a new implementation, the system should **still work as expected**.

> Replace → Still works.

### 4. I — Interface Segregation

Give a component **only the essential features it actually needs**, instead of giving it a large interface with unnecessary methods.

> Give only what is needed.
// a api is used to fetch user data from a database

in the previous the userreportry has the varoius method .
type UserRepository interface {
	GetUser(id int)
	CreateUser()
	UpdateUser()
	DeleteUser()
}
// now the user reader interface is used to fetch user data from a database  mwe can replace that with the user repository interface and it will still work as expected.

type UserReader interface {
    GetUser(id int)
}

### 5. D — Dependency Inversion

The system should not depend directly on one specific implementation or brand.

> Depend on abstraction, not one implementation.






# How do we implement SOLID in Go?

In Go, we mainly use:

```text
Structs
   ↓
Interfaces
   ↓
Small responsibilities
   ↓
Dependency Injection
   ↓
Loose coupling
```

For example:

```go
type UserRepository interface {
    Save(name string) error
}

type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
    return UserService{
        repo: repo,
    }
}
```

Here:

* `UserService` handles user/business logic.
* `UserRepository` defines the required behavior.
* The service does not care whether the implementation is MySQL or PostgreSQL.
* The repository is passed into the service through **Dependency Injection**.

So in Go, SOLID is mainly achieved by **separating responsibilities, using small interfaces, and injecting dependencies instead of tightly coupling components**.
