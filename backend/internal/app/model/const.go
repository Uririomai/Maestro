package model

const UniqueConstraintViolationCode = "23505"

const ImageBucketName = "image-bucket"

const (
	OrderStatusWaitingPayment = iota + 1
	OrderStatusPaid
	// TODO: Дописать список
)
