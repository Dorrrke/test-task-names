package services

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/Dorrrke/test-task-names/internal/domain/models"
	"github.com/Dorrrke/test-task-names/internal/logger"
	"go.uber.org/zap"
)

type EnrichService struct {
	client http.Client
}

type Result struct {
	model models.NameData
	err   error
}

func New() *EnrichService {
	client := &http.Client{}
	return &EnrichService{
		client: *client,
	}
}

func (es *EnrichService) addAge(wgIter int, wg *sync.WaitGroup, name models.NameData) Result {
	logger.Log.Debug("Age enrich name", zap.String("Name", name.Name))
	res := Result{
		model: name,
	}
	req, err := http.NewRequest("GET", "https://api.agify.io/?name="+res.model.Name, nil)
	if err != nil {
		logger.Log.Error("Create age api request error", zap.Error(err))
		res.err = err
		return res
	}
	resp, err := es.client.Do(req)
	if err != nil {
		logger.Log.Error("Create age api response error", zap.Error(err))
		res.err = err
		return res
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error("Error convert age req to byte", zap.Error(err))
		res.err = err
		return res
	}
	err = json.Unmarshal(body, &res.model)
	if err != nil {
		logger.Log.Error("Error decoding body", zap.Error(err))
		res.err = err
		return res
	}
	wg.Done()
	return res
}

func (es *EnrichService) addGender(wgIter int, wg *sync.WaitGroup, name models.NameData) Result {
	logger.Log.Debug("Gender enrich name", zap.String("Name", name.Name))
	res := Result{
		model: name,
	}
	req, err := http.NewRequest("GET", "https://api.genderize.io/?name="+res.model.Name, nil)
	if err != nil {
		logger.Log.Error("Create gender api request error", zap.Error(err))
		res.err = err
		return res
	}
	resp, err := es.client.Do(req)
	if err != nil {
		logger.Log.Error("Create gender api response error", zap.Error(err))
		res.err = err
		return res
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error("Error convert gender req to byte", zap.Error(err))
		res.err = err
		return res
	}
	err = json.Unmarshal(body, &res.model)
	if err != nil {
		logger.Log.Error("Error decoding gender body", zap.Error(err))
		res.err = err
		return res
	}
	wg.Done()
	return res
}

func (es *EnrichService) addNational(wgIter int, wg *sync.WaitGroup, name models.NameData) Result {
	logger.Log.Debug("National enrich name", zap.String("Name", name.Name))
	res := Result{
		model: name,
	}
	req, err := http.NewRequest("GET", "https://api.nationalize.io/?name="+res.model.Name, nil)
	if err != nil {
		logger.Log.Error("Create national api request error", zap.Error(err))
		res.err = err
		return res
	}
	resp, err := es.client.Do(req)
	if err != nil {
		logger.Log.Error("Create national api response error", zap.Error(err))
		res.err = err
		return res
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error("Error convert national req to byte", zap.Error(err))
		res.err = err
		return res
	}
	var national models.NationalApiModel
	err = json.Unmarshal(body, &national)
	if err != nil {
		logger.Log.Error("Error decoding national body", zap.Error(err))
		res.err = err
		return res
	}
	res.model.National = national.National[0].CountryID
	wg.Done()
	return res
}

func (es *EnrichService) EnrichName(name models.NameData) (models.NameData, error) {
	logger.Log.Debug("Enrich name", zap.String("Name", name.Name))
	var wg sync.WaitGroup
	var nameValue = name
	var resAge Result
	var resGender Result
	var resNational Result
	wg.Add(3)
	go func() {
		resAge = es.addAge(1, &wg, nameValue)
	}()
	go func() {
		resGender = es.addGender(2, &wg, nameValue)
	}()

	go func() {
		resNational = es.addNational(3, &wg, nameValue)
	}()

	wg.Wait()
	if resAge.err != nil {
		logger.Log.Error("enrich name error", zap.Error(resAge.err))
		return models.NameData{}, resAge.err
	}
	if resGender.err != nil {
		logger.Log.Error("enrich name error", zap.Error(resGender.err))
		return models.NameData{}, resGender.err
	}
	if resNational.err != nil {
		logger.Log.Error("enrich name error", zap.Error(resNational.err))
		return models.NameData{}, resNational.err
	}
	nameValue.Age = resAge.model.Age
	nameValue.Gender = resGender.model.Gender
	nameValue.National = resNational.model.National

	return nameValue, nil
}
