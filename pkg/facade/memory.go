package facade

type Memory interface {
	read(position int) byte
	write(position int, data byte)
	free(position int)
	cleanAll()
}
