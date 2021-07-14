package main

func removeEnemyAt(i int, list []Enemy) []Enemy {
	return append(list[:i], list[i+1:]...)
}

func removeBulletAt(i int, list []Bullet) []Bullet {
	return append(list[:i], list[i+1:]...)
}
