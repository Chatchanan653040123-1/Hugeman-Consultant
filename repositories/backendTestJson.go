package repositories

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type backendTestsRepo struct {
	FilePath string
}

func NewBackendTestsRepo(FilePath string) *backendTestsRepo {
	return &backendTestsRepo{FilePath: FilePath}
}

func (r *backendTestsRepo) CreateATask(createATask BackendTests) (*BackendTests, error) {
	var existingData []BackendTests

	data, err := ioutil.ReadFile(r.FilePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if len(data) == 0 {
		existingData = []BackendTests{}
	} else {
		err = json.Unmarshal(data, &existingData)
		if err != nil {
			return nil, err
		}
	}

	existingData = append(existingData, createATask)
	newData, err := json.Marshal(existingData)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(r.FilePath, newData, 0644)
	if err != nil {
		return nil, err
	}

	return &createATask, nil
}
func (r *backendTestsRepo) UpdateATask(updateATask BackendTests, id uuid.UUID) (*BackendTests, error) {
	var existingData []BackendTests

	data, err := ioutil.ReadFile(r.FilePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if len(data) == 0 {
		existingData = []BackendTests{}
	} else {
		err = json.Unmarshal(data, &existingData)
		if err != nil {
			return nil, err
		}
	}

	for i, v := range existingData {
		if v.ID == id {
			existingData[i] = updateATask
		}
	}

	newData, err := json.Marshal(existingData)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(r.FilePath, newData, 0644)
	if err != nil {
		return nil, err
	}

	return &updateATask, nil
}
func (r *backendTestsRepo) GetAllOfTasks() ([]BackendTests, error) {
	var existingData []BackendTests

	data, err := ioutil.ReadFile(r.FilePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if len(data) == 0 {
		existingData = []BackendTests{}
	} else {
		err = json.Unmarshal(data, &existingData)
		if err != nil {
			return nil, err
		}
	}

	return existingData, nil
}
