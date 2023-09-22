/*
spoonacular API

The spoonacular Nutrition, Recipe, and Food API allows you to access over thousands of recipes, thousands of ingredients, 800,000 food products, over 100,000 menu items, and restaurants. Our food ontology and semantic recipe search engine makes it possible to search for recipes using natural language queries, such as \"gluten free brownies without sugar\" or \"low fat vegan cupcakes.\" You can automatically calculate the nutritional information for any recipe, analyze recipe costs, visualize ingredient lists, find recipes for what's in your fridge, find recipes based on special diets, nutritional requirements, or favorite ingredients, classify recipes into types and cuisines, convert ingredient amounts, or even compute an entire meal plan. With our powerful API, you can create many kinds of food and especially nutrition apps.  Special diets/dietary requirements currently available include: vegan, vegetarian, pescetarian, gluten free, grain free, dairy free, high protein, whole 30, low sodium, low carb, Paleo, ketogenic, FODMAP, and Primal.

API version: 1.1
Contact: mail@spoonacular.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ImageAnalysisByURL200ResponseNutrition struct for ImageAnalysisByURL200ResponseNutrition
type ImageAnalysisByURL200ResponseNutrition struct {
	RecipesUsed int32 `json:"recipesUsed"`
	Calories ImageAnalysisByURL200ResponseNutritionCalories `json:"calories"`
	Fat ImageAnalysisByURL200ResponseNutritionCalories `json:"fat"`
	Protein ImageAnalysisByURL200ResponseNutritionCalories `json:"protein"`
	Carbs ImageAnalysisByURL200ResponseNutritionCalories `json:"carbs"`
}

// NewImageAnalysisByURL200ResponseNutrition instantiates a new ImageAnalysisByURL200ResponseNutrition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAnalysisByURL200ResponseNutrition(recipesUsed int32, calories ImageAnalysisByURL200ResponseNutritionCalories, fat ImageAnalysisByURL200ResponseNutritionCalories, protein ImageAnalysisByURL200ResponseNutritionCalories, carbs ImageAnalysisByURL200ResponseNutritionCalories) *ImageAnalysisByURL200ResponseNutrition {
	this := ImageAnalysisByURL200ResponseNutrition{}
	this.RecipesUsed = recipesUsed
	this.Calories = calories
	this.Fat = fat
	this.Protein = protein
	this.Carbs = carbs
	return &this
}

// NewImageAnalysisByURL200ResponseNutritionWithDefaults instantiates a new ImageAnalysisByURL200ResponseNutrition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAnalysisByURL200ResponseNutritionWithDefaults() *ImageAnalysisByURL200ResponseNutrition {
	this := ImageAnalysisByURL200ResponseNutrition{}
	return &this
}

// GetRecipesUsed returns the RecipesUsed field value
func (o *ImageAnalysisByURL200ResponseNutrition) GetRecipesUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RecipesUsed
}

// GetRecipesUsedOk returns a tuple with the RecipesUsed field value
// and a boolean to check if the value has been set.
func (o *ImageAnalysisByURL200ResponseNutrition) GetRecipesUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipesUsed, true
}

// SetRecipesUsed sets field value
func (o *ImageAnalysisByURL200ResponseNutrition) SetRecipesUsed(v int32) {
	o.RecipesUsed = v
}

// GetCalories returns the Calories field value
func (o *ImageAnalysisByURL200ResponseNutrition) GetCalories() ImageAnalysisByURL200ResponseNutritionCalories {
	if o == nil {
		var ret ImageAnalysisByURL200ResponseNutritionCalories
		return ret
	}

	return o.Calories
}

// GetCaloriesOk returns a tuple with the Calories field value
// and a boolean to check if the value has been set.
func (o *ImageAnalysisByURL200ResponseNutrition) GetCaloriesOk() (*ImageAnalysisByURL200ResponseNutritionCalories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Calories, true
}

// SetCalories sets field value
func (o *ImageAnalysisByURL200ResponseNutrition) SetCalories(v ImageAnalysisByURL200ResponseNutritionCalories) {
	o.Calories = v
}

// GetFat returns the Fat field value
func (o *ImageAnalysisByURL200ResponseNutrition) GetFat() ImageAnalysisByURL200ResponseNutritionCalories {
	if o == nil {
		var ret ImageAnalysisByURL200ResponseNutritionCalories
		return ret
	}

	return o.Fat
}

// GetFatOk returns a tuple with the Fat field value
// and a boolean to check if the value has been set.
func (o *ImageAnalysisByURL200ResponseNutrition) GetFatOk() (*ImageAnalysisByURL200ResponseNutritionCalories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fat, true
}

// SetFat sets field value
func (o *ImageAnalysisByURL200ResponseNutrition) SetFat(v ImageAnalysisByURL200ResponseNutritionCalories) {
	o.Fat = v
}

// GetProtein returns the Protein field value
func (o *ImageAnalysisByURL200ResponseNutrition) GetProtein() ImageAnalysisByURL200ResponseNutritionCalories {
	if o == nil {
		var ret ImageAnalysisByURL200ResponseNutritionCalories
		return ret
	}

	return o.Protein
}

// GetProteinOk returns a tuple with the Protein field value
// and a boolean to check if the value has been set.
func (o *ImageAnalysisByURL200ResponseNutrition) GetProteinOk() (*ImageAnalysisByURL200ResponseNutritionCalories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protein, true
}

// SetProtein sets field value
func (o *ImageAnalysisByURL200ResponseNutrition) SetProtein(v ImageAnalysisByURL200ResponseNutritionCalories) {
	o.Protein = v
}

// GetCarbs returns the Carbs field value
func (o *ImageAnalysisByURL200ResponseNutrition) GetCarbs() ImageAnalysisByURL200ResponseNutritionCalories {
	if o == nil {
		var ret ImageAnalysisByURL200ResponseNutritionCalories
		return ret
	}

	return o.Carbs
}

// GetCarbsOk returns a tuple with the Carbs field value
// and a boolean to check if the value has been set.
func (o *ImageAnalysisByURL200ResponseNutrition) GetCarbsOk() (*ImageAnalysisByURL200ResponseNutritionCalories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Carbs, true
}

// SetCarbs sets field value
func (o *ImageAnalysisByURL200ResponseNutrition) SetCarbs(v ImageAnalysisByURL200ResponseNutritionCalories) {
	o.Carbs = v
}

func (o ImageAnalysisByURL200ResponseNutrition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipesUsed"] = o.RecipesUsed
	}
	if true {
		toSerialize["calories"] = o.Calories
	}
	if true {
		toSerialize["fat"] = o.Fat
	}
	if true {
		toSerialize["protein"] = o.Protein
	}
	if true {
		toSerialize["carbs"] = o.Carbs
	}
	return json.Marshal(toSerialize)
}

type NullableImageAnalysisByURL200ResponseNutrition struct {
	value *ImageAnalysisByURL200ResponseNutrition
	isSet bool
}

func (v NullableImageAnalysisByURL200ResponseNutrition) Get() *ImageAnalysisByURL200ResponseNutrition {
	return v.value
}

func (v *NullableImageAnalysisByURL200ResponseNutrition) Set(val *ImageAnalysisByURL200ResponseNutrition) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAnalysisByURL200ResponseNutrition) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAnalysisByURL200ResponseNutrition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAnalysisByURL200ResponseNutrition(val *ImageAnalysisByURL200ResponseNutrition) *NullableImageAnalysisByURL200ResponseNutrition {
	return &NullableImageAnalysisByURL200ResponseNutrition{value: val, isSet: true}
}

func (v NullableImageAnalysisByURL200ResponseNutrition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAnalysisByURL200ResponseNutrition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


