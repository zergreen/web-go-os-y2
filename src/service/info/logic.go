package info

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func createUserInfo(Id int, FirstName string, LastName string, District string) (result statusDetail, err error) {
	fmt.Printf("\n checkUserInfo... ")
	checkUId, err := checkUserInfo(Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("\n CheckUserInfo fail %s", err)
		return
	}
	fmt.Printf("checkUserId : %v",checkUId)
	if checkUId.Id == 0 {
		fmt.Printf("\n createUser... ")
		err = insertUserInfo(Id, FirstName, LastName)
		if err != nil {
			fmt.Printf("\n InsertUserInfo fail ")
			return

		}
	} else {
		fmt.Printf("\n UserAlreadyCreate... ")
		fmt.Printf("\n checkgeographiesId... ")
		var codeGeo geographyCode
		codeGeo, err = checkgeographiesById(Id)
		if err != nil {
			fmt.Printf("\n checkgeographiesId fail ")
			return
		}
		fmt.Printf("\n updatePopulation... ")
		codeGeo.Population -= 1
		_, err = updatePopulationByGeographyId(codeGeo.GeographyId, codeGeo.Population)
		if err != nil {
			fmt.Printf("\n updatePopulation fail ")
			return
		}
	}

	fmt.Printf("\n updateDistrictInfo... ")
	result, err = updateDistrictInfo(Id, District)
	if err != nil {
		fmt.Printf("\n updateDistrictInfo fail ")
		return
	}

	fmt.Printf("\n checkgeographiesId... ")
	var codeGeo geographyCode
	codeGeo, err = checkgeographiesId(District)
	if err != nil {
		fmt.Printf("\n checkgeographiesId fail ")
		return
	}

	fmt.Printf("\n updatePopulation... ")
	codeGeo.Population += 1
	_, err = updatePopulationByGeographyId(codeGeo.GeographyId, codeGeo.Population)
	if err != nil {
		fmt.Printf("\n updatePopulation fail ")
		return
	}
	fmt.Printf("\n InsertUserInfo succeed ")
	return
}

func updateDistrictInfo(Id int, District string) (result statusDetail, err error) {
	fmt.Printf("\n updateDistrict... ")
	_, err = insertDistrictByUserId(Id, District)
	if err != nil {
		fmt.Printf("\n InsertDistrictByUserId fail ")
		return
	}

	result = statusDetail{
		Status: "succeed",
	}
	fmt.Printf("\n InsertDistrictByUserId succeed ")

	return
}

func showGeography() (result []showAllGeography, err error) {

	fmt.Printf("\n searchGeographyInfo... ")
	result, err = searchGeographyInfo()
	if err != nil {
		fmt.Printf("\n searchGeographyInfo fail ")
		return
	}
	fmt.Printf("\n searchGeographyInfo succeed ")
	return
}

func checkInfo(Id int, FirstName string, LastName string, District string) (result searchInfo, err error) {

	fmt.Printf("\n checkInfo...")
	var addressList address
	addressList, err = checkInfoByDistrict(District)
	if err != nil {
		fmt.Printf("\n checkInfo fail ")
		return
	}
    fmt.Printf("\n firstName:%v",FirstName)
	result = searchInfo{
		Id:         Id,
		FirstName:  FirstName,
		LastName:   LastName,
		District:   District,
		Province:   addressList.NameTh,
		Region:     addressList.NameGeo,
		Population: addressList.Population,
	}
	fmt.Printf("\n checkInfo succeed ")
	return
}
