package config

import "github.com/kelseyhightower/envconfig"

func NewAPIConfig() (*APIConfig, error) {

	var db DBConfig
	if err := envconfig.Process("DB", &db); err != nil {
		return nil, err
	}

	var dbWriter DBConnection
	if err := envconfig.Process("DB_WRITE", &dbWriter); err != nil {
		return nil, err
	}
	db.Writer = &dbWriter

	var dbReader DBConnection
	if err := envconfig.Process("DB_READ", &dbReader); err != nil {
		return nil, err
	}
	if dbReader.Host == "" {
		db.Reader = &dbWriter
	} else {
		db.Reader = &dbReader
	}

	var apiConfig APIConfig
	if err := envconfig.Process("API", &apiConfig); err != nil {
		return nil, err
	}

	apiConfig.DB = &db

	//var redis RedisConfig
	//if err := envconfig.Process("REDIS", &redis); err != nil {
	//	return nil, err
	//}
	//apiConfig.Redis = &redis
	//
	//var s3 AWSS3Config
	//if err := envconfig.Process("AWS_S3", &s3); err != nil {
	//	return nil, err
	//}
	//apiConfig.Storage = &s3
	//
	//var s3aws AWSConfig
	//if err := envconfig.Process("AWS_S3", &s3aws); err != nil {
	//	return nil, err
	//}
	//s3.AWSConfig = &s3aws
	//
	//var xray XRayConfig
	//if err := envconfig.Process("AWS_XRAY", &xray); err != nil {
	//	return nil, err
	//}
	//apiConfig.XRay = &xray

	return &apiConfig, nil
}
