package interfaces

type CallbackRegistrar[dbProvider DataBaseProvider] interface {
	RegisterCallback(dbProvider dbProvider)
}
