package repositories

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testBackend/logs"

	"github.com/google/uuid"
	"go.uber.org/zap"
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
		logs.Error("Error while reading data from the file", zap.Error(err))
		return nil, err
	}

	if len(data) == 0 {
		existingData = []BackendTests{}
	} else {
		err = json.Unmarshal(data, &existingData)
		if err != nil {
			logs.Error("Error while unmarshalling data from the file", zap.Error(err))
			return nil, err
		}
	}

	existingData = append(existingData, createATask)
	newData, err := json.Marshal(existingData)
	if err != nil {
		logs.Error("Error while marshalling data", zap.Error(err))
		return nil, err
	}

	err = ioutil.WriteFile(r.FilePath, newData, 0644)
	if err != nil {
		logs.Error("Error while writing data to the file", zap.Error(err))
		return nil, err
	}

	logs.Info("Task created successfully")
	return &createATask, nil
}

func (r *backendTestsRepo) UpdateATask(updateATask BackendTests, id uuid.UUID) (*BackendTests, error) {
	var existingData []BackendTests

	data, err := ioutil.ReadFile(r.FilePath)
	if err != nil && !os.IsNotExist(err) {
		logs.Error("Error while reading data from the file", zap.Error(err))
		return nil, err
	}

	if len(data) == 0 {
		existingData = []BackendTests{}
	} else {
		err = json.Unmarshal(data, &existingData)
		if err != nil {
			logs.Error("Error while unmarshalling data from the file", zap.Error(err))
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
		logs.Error("Error while marshalling data", zap.Error(err))
		return nil, err
	}

	err = ioutil.WriteFile(r.FilePath, newData, 0644)
	if err != nil {
		logs.Error("Error while writing data to the file", zap.Error(err))
		return nil, err
	}

	logs.Info("Task updated successfully")
	return &updateATask, nil
}

func (r *backendTestsRepo) GetAllOfTasks() ([]BackendTests, error) {
	var existingData []BackendTests

	data, err := ioutil.ReadFile(r.FilePath)
	if err != nil && !os.IsNotExist(err) {
		logs.Error("Error while reading data from the file", zap.Error(err))
		return nil, err
	}

	if len(data) == 0 {
		existingData = []BackendTests{}
	} else {
		err = json.Unmarshal(data, &existingData)
		if err != nil {
			logs.Error("Error while unmarshalling data from the file", zap.Error(err))
			return nil, err
		}
	}

	logs.Info("Retrieved all tasks successfully")
	return existingData, nil
}
