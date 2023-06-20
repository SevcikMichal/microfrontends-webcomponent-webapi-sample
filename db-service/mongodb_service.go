package dbservice

import (
	"ambulance-webapi/models"
	"errors"
	"sync"
)

var syncMap sync.Map

func Connect() {
	syncMap = sync.Map{}

	ambulance := &models.Ambulance{
		Id:         "bobulova",
		Name:       "Ambulancia všeobecného lekára Dr. Bobulová",
		RoomNumber: "211 - 2.posch",
		PredefinedConditions: []models.Condition{
			{Value: "Teploty", Code: "subfebrilia", Reference: "https://zdravoteka.sk/priznaky/zvysena-telesna-teplota/", TypicalDurationMinutes: 20},
			{Value: "Kontrola", Code: "folowup", TypicalDurationMinutes: 15},
			{Value: "Nevoľnosť", Code: "nausea", TypicalDurationMinutes: 45, Reference: "https://zdravoteka.sk/priznaky/nevolnost/"},
		},
		WaitingList: []models.WaitingListEntry{},
	}

	CreateAmbulance(ambulance)
}

func Disconnect() {

}

func CreateAmbulance(ambulance *models.Ambulance) error {

	collection, _ := syncMap.LoadOrStore("Ambulances", &sync.Map{})
	collectionCasted := collection.(*sync.Map)
	_, loaded := collectionCasted.LoadOrStore(ambulance.Id, &models.Ambulance{})

	if loaded {
		errorMessage := "Ambulance with following id already exists: " + ambulance.Id
		return errors.New(errorMessage)
	}

	collectionCasted.Store(ambulance.Id, ambulance)
	syncMap.Store("Ambulances", collectionCasted)

	return nil
}

func GetAmbulance(id string) (*models.Ambulance, error) {
	ambulance, err := getAmbulance(id)
	return ambulance, err
}

func getAmbulance(id string) (*models.Ambulance, error) {

	collection, ok := syncMap.Load("Ambulances")
	if !ok {
		return &models.Ambulance{}, errors.New("ambulance collection does not exist")
	}
	collectionCasted := collection.(*sync.Map)

	ambulances, ok := collectionCasted.Load(id)
	if !ok {
		return &models.Ambulance{}, errors.New("ambulance collection does not exist")
	}
	ambulancesCasted := ambulances.(*models.Ambulance)

	return ambulancesCasted, nil
}

func DeleteWaitingListEntry(ambulanceId string, entryId string) error {
	ambulance, err := getAmbulance(ambulanceId)
	if err != nil {
		return err
	}

	for index, entry := range ambulance.WaitingList {
		if entry.Id == entryId {
			waitingList := removeEntry(ambulance.WaitingList, index)
			ambulance.WaitingList = waitingList

			collection, _ := syncMap.Load("Ambulances")
			collectionCasted := collection.(*sync.Map)

			collectionCasted.Store(ambulance.Id, ambulance)
			syncMap.Store("Ambulances", collectionCasted)
			return nil
		}
	}
	errorMessage := "There is no entry with id: " + entryId + " for ambulance with following id: " + ambulanceId
	return errors.New(errorMessage)
}

func removeEntry(waitingList []models.WaitingListEntry, index int) []models.WaitingListEntry {
	return append(waitingList[:index], waitingList[index+1:]...)
}

func UpdateWaitingListForAmbulance(ambulanceId string, waitingList []models.WaitingListEntry) error {
	collection, _ := syncMap.LoadOrStore("Ambulances", &sync.Map{})
	collectionCasted := collection.(*sync.Map)
	ambulance, _ := collectionCasted.Load(ambulanceId)
	ambulanceCasted := ambulance.(*models.Ambulance)

	ambulanceCasted.WaitingList = waitingList

	collectionCasted.Store(ambulanceCasted.Id, ambulanceCasted)
	syncMap.Store("Ambulances", collectionCasted)

	return nil
}
