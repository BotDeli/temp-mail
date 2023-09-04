package connections

import "time"

func (r *RedisConnections) StartLifeConnection(uuid string) {
	r.Client.Set(r.Ctx, uuid, true, 30*time.Minute)
}

func (r *RedisConnections) StopLifeConnection(uuid string) {
	r.Client.Set(r.Ctx, uuid, false, 30*time.Minute)
}

func (r *RedisConnections) IsLifeConnection(uuid string) bool {
	cmd := r.Client.Get(r.Ctx, uuid)
	life, err := cmd.Bool()
	return err == nil && life
}
