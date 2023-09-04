package connections

type StorageConnections interface {
	StartLifeConnection(uuid string)
	StopLifeConnection(uuid string)
	IsLifeConnection(uuid string) bool
}
