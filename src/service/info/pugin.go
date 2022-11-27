package info

import (
	mssql "golang-101-master/src/database/mssql"
)

func checkUserInfo(Id int) (result checkuser, err error) {
	if err = mssql.DB.Table("dbo.user").Select("id").Where("id = ?", Id).Find(&result).Error; err != nil {
		return
	}
	return
}

func insertUserInfo(Id int, FirstName string, LastName string) (err error) {
	var req userList
	req.Id = Id
	req.FirstName = FirstName
	req.LastName = LastName

	if err = mssql.DB.Table("dbo.user").Create(&req).Error; err != nil {
		return
	}
	return
}

func insertDistrictByUserId(Id int, District string) (_, err error) {
	if err = mssql.DB.Table("dbo.user").Where("id = ?", Id).Update(map[string]interface{}{
		"district": District,
	}).Error; err != nil {
		return
	}
	return
}

func searchGeographyInfo() (result []showAllGeography, err error) {
	if err = mssql.DB.Table("dbo.geographies").Select("*").Find(&result).Error; err != nil {
		return
	}
	return
}

func checkInfoByDistrict(District string) (result address, err error) {

	if err = mssql.DB.Select("p.name_th, g.name_geo, g.population").Table("dbo.amphures AS a").
		Joins("INNER JOIN dbo.provinces AS p ON p.id = a.province_id ").
		Joins("INNER JOIN dbo.geographies AS g ON g.id = p.geography_id ").
		Where("a.name_th = ?", District).Find(&result).Error; err != nil {
		return
	}
	return
}

func checkgeographiesId(District string) (result geographyCode, err error) {

	if err = mssql.DB.Select("p.geography_id, g.population").Table("dbo.amphures AS a").
		Joins("INNER JOIN dbo.provinces AS p ON p.id = a.province_id ").
		Joins("INNER JOIN dbo.geographies AS g ON g.id = p.geography_id ").
		Where("a.name_th = ?", District).Find(&result).Error; err != nil {
		return
	}
	return
}

func checkgeographiesById(Id int) (result geographyCode, err error) {
	if err = mssql.DB.Select("p.geography_id, g.population").Table("dbo.[user] AS u").
		Joins("INNER JOIN dbo.amphures AS a ON a.name_th = u.district").
		Joins("INNER JOIN dbo.provinces AS p ON p.id = a.province_id").
		Joins("INNER JOIN dbo.geographies AS g ON g.id = p.geography_id").
		Where("u.id = ?", Id).Find(&result).Error; err != nil {
		return
	}
	return
}

func updatePopulationByGeographyId(geographyId int, newPopulation int) (r, err error) {

	if err = mssql.DB.Table("dbo.geographies").Where("id = ?", geographyId).Update(map[string]interface{}{
		"population": newPopulation,
	}).Error; err != nil {
		return
	}
	return
}
