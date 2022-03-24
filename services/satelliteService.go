package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/oskarincon/operation-quasar-go/constants"
	"github.com/oskarincon/operation-quasar-go/models"
)

var cacheKenobi ttlcache.SimpleCache = ttlcache.NewCache()
var cacheSkywalker ttlcache.SimpleCache = ttlcache.NewCache()
var cacheSato ttlcache.SimpleCache = ttlcache.NewCache()

func PostInfoSatellite(satelliteName string, data models.Satellite) (response models.Satellite, err error) {
	fmt.Printf("[PostInfoSatellite] - data: %#v\n", data)
	switch strings.ToLower(satelliteName) {
	case "kenobi":
		data.Name = "kenobi"
		cacheKenobi.Set("kenobi", data)
	case "skywalker":
		data.Name = "skywalker"
		cacheSkywalker.Set("skywalker", data)
	case "sato":
		data.Name = "sato"
		cacheSato.Set("sato", data)
	default:
		return response, errors.New(constants.SATELLITE_ERROR)
	}
	response = data
	return response, err
}

func GetCache(key string, cache ttlcache.SimpleCache) (models.Satellite, bool) {
	var emp models.Satellite
	data, err := cache.Get(key)
	if err != nil {
		return emp, false
	}
	emp = data.(models.Satellite)
	return emp, true
}

func GetInfoSatellite() (response models.Satellites, err error) {
	dataKenobi, foundKenobi := GetCache("kenobi", cacheKenobi)
	fmt.Printf("[GetInfoSatellite] - dataKenobi: %#v\n", dataKenobi)
	dataSkywalker, foundSkywalker := GetCache("skywalker", cacheSkywalker)
	fmt.Printf("[GetInfoSatellite] - dataSkywalker: %#v\n", dataSkywalker)
	dataSato, foundSato := GetCache("sato", cacheSato)
	fmt.Printf("[GetInfoSatellite] - dataSato: %#v\n", dataSato)
	if !foundKenobi || !foundSkywalker || !foundSato {
		return response, errors.New(constants.DATA_NOT_FOUND_ERROR)
	}
	data := []models.Satellite{dataKenobi, dataSkywalker, dataSato}
	fmt.Printf("[GetInfoSatellite] - data res: %#v\n", data)
	response.Satellites = data
	return response, nil
}
