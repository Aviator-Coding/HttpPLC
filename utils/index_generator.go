package utils

import (
	"github.com/Aviator-Coding/HttpPLC/configs"
)

func CreateIndex() {
	configs.CreateIndex(configs.GetCollection(configs.DB, "HMIUsers"), "batchid", true)
}
