package config

import "moises-ba/ms-criptcoin-vote/utils"

const (
	MONGO_SERVER_URL       = "MONGO_SERVER_URL"
	MONGO_USER             = "MONGO_USER"
	MONGO_PASSWORD         = "MONGO_PASSWORD"
	MONGO_CRIPTCOINVOTE_BD = "MONGO_CRIPTCOINVOTE_BD"
	KAFKA_BROKER_URL       = "KAFKA_BROKER_URL"
	VOTE_TOPIC             = "KAFKA_VOTE_TOPIC"

	JWT_PASSWORD = "JWT_PASSWORD"
)

func GetMogoServerURL() string {
	return utils.GetEnv(MONGO_SERVER_URL, "mongodb://localhost:27017")
}

func GetMongoUser() string {
	return utils.GetEnv(MONGO_USER, "root")
}

func GetMongoPassWord() string {
	return utils.GetEnv(MONGO_PASSWORD, "example")
}

func GetVoteTopic() string {
	return utils.GetEnv(VOTE_TOPIC, "vote_topic")
}

func GetKafkaBrokerURL() string {
	return utils.GetEnv(KAFKA_BROKER_URL, "localhost:9092")
}

func GetJWTPassword() string {
	return utils.GetEnv(JWT_PASSWORD, "chave_jwt")
}

func GetMongoCriptCoinDB() string {
	return utils.GetEnv(MONGO_CRIPTCOINVOTE_BD, "criptcoinDB")
}
