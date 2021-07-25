package parser

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ilyavovnenko/shops_categories_ms/internal/attribute"
	"github.com/ilyavovnenko/shops_categories_ms/internal/category"

	log "github.com/sirupsen/logrus"
)

type BolParser struct {
	datamodelURL string
	log          log.Logger
	shopID       uint
	catRepo      category.CategoryRepository
	attrRepo     attribute.AttributeRepository
}

type Attributes struct {
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	Definition          string         `json:"definition"`
	FillingInstructions string         `json:"fillingInstructions"`
	MultiValue          bool           `json:"multiValue"`
	Validation          sql.NullString `json:"Validation"`
	LovId               string         `json:"lovId"`
	EnrichmentLevel     int8           `json:"enrichmentLevel"`
}

type Chunks struct {
	Chunks []Chunk `json:"chunks"`
}

type Chunk struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Attributes []Attribute `json:"attributes"`
}

type ListsOfValues struct {
	ListsOfValues []ListOfValues `json:"lovs"`
}

type ListOfValues struct {
	ID     string   `json:"id"`
	Values []string `json:"values"`
}

func New(shopID uint, log log.Logger, datamodelURL string, catRepo category.CategoryRepository, attrRepo attribute.AttributeRepository) *BolParser {
	return &BolParser{
		datamodelURL: datamodelURL,
		log:          log,
		shopID:       shopID,
		catRepo:      catRepo,
		attrRepo:     attrRepo,
	}
}

func (bp *BolParser) ParseDatamodel() {
	bp.log.Info("Start to download the file")
	response, err := http.Get(bp.datamodelURL)
	if err != nil {
		bp.log.Error(err)
	}

	bp.log.Info("Reading file into memory")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		bp.log.Error(err)
	}

	bp.log.Info("Converting file to objects")
	var chunks Chunks
	json.Unmarshal([]byte(body), &chunks)

	var attributes Attributes
	json.Unmarshal([]byte(body), &attributes)

	var listsOfValues ListsOfValues
	json.Unmarshal([]byte(body), &listsOfValues)

	bp.log.Info("Deactivating all categories and attributes")
	bp.catRepo.DeactivateShopCategories(bp.shopID)
	bp.attrRepo.DeactivateShopAttributes(bp.shopID)

	bp.log.Info("Storing data into DB")
	err = bp.StoreCategoriesData(&attributes, &chunks, &listsOfValues)
	if err != nil {
		bp.log.Error(err)
	}

	bp.log.Info("Remove not active categories, attributes and their values")
	// todo: add removing function
}

func (bp *BolParser) StoreCategoriesData(attributes *Attributes, chunks *Chunks, listsOfValues *ListsOfValues) error {
	// this is here, because I don't want to loop every time for find some attribute by ID
	assocAttributes := make(map[string]Attribute)
	for _, assocAttr := range attributes.Attributes {
		assocAttributes[assocAttr.ID] = assocAttr
	}

	// this is here, because I don't want to loop every time for find some list of values by ID
	assocListsOfValues := make(map[string]ListOfValues)
	for _, assocLOV := range listsOfValues.ListsOfValues {
		assocListsOfValues[assocLOV.ID] = assocLOV
	}

	for _, chunk := range chunks.Chunks {
		var currentCategory category.Category

		currentCategory.ShopID = bp.shopID
		currentCategory.ShopExternalId = chunk.ID
		currentCategory.Name = chunk.Name
		currentCategory.UpdatedAt = time.Now()

		cat, err := bp.catRepo.StoreCategory(currentCategory)
		if err != nil {
			log.Error(err)
		}

		for _, chunkAttr := range chunk.Attributes {
			attributeInfo := assocAttributes[chunkAttr.ID]
			currentAttribute := prepareCurrentAttribute(cat, chunkAttr, attributeInfo)

			attr, err := bp.attrRepo.StoreAttribute(currentAttribute)
			if err != nil {
				log.Error(err)
			}

			if chunkAttr.LovId != "" {
				for _, attrValue := range assocListsOfValues[chunkAttr.LovId].Values {
					var currentAttrVal attribute.AttributeValue

					currentAttrVal.AttributeID = attr.ID
					currentAttrVal.TechName = attrValue
					currentAttrVal.Name = attrValue
					currentAttrVal.UpdatedAt = time.Now()

					_, err := bp.attrRepo.StoreAttributeValue(currentAttrVal)
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
	}

	return nil
}

func prepareCurrentAttribute(category category.Category, chunkAttr Attribute, attributeInfo Attribute) attribute.Attribute {
	var currentAttribute attribute.Attribute

	currentAttribute.CategoryID = category.ID
	currentAttribute.TechName = chunkAttr.ID
	currentAttribute.Name = attributeInfo.Name
	currentAttribute.Type = "string"
	currentAttribute.UpdatedAt = time.Now()
	currentAttribute.Validation = attributeInfo.Validation

	if attributeInfo.MultiValue {
		currentAttribute.Multivalue = 1
	}

	if chunkAttr.EnrichmentLevel == 0 || chunkAttr.EnrichmentLevel == 1 {
		currentAttribute.Mandatory = 1
	} else {
		currentAttribute.Priority = sql.NullInt64{Int64: 1, Valid: true}
	}

	return currentAttribute
}
