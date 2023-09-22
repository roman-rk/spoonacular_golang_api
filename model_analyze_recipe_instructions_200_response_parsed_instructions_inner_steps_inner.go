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

// AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner struct for AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner
type AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner struct {
	Number float32 `json:"number"`
	Step string `json:"step"`
	Ingredients []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner `json:"ingredients,omitempty"`
	Equipment []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner `json:"equipment,omitempty"`
}

// NewAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner instantiates a new AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner(number float32, step string) *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner {
	this := AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner{}
	this.Number = number
	this.Step = step
	return &this
}

// NewAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerWithDefaults instantiates a new AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerWithDefaults() *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner {
	this := AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner{}
	return &this
}

// GetNumber returns the Number field value
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) SetNumber(v float32) {
	o.Number = v
}

// GetStep returns the Step field value
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetStep() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetStepOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) SetStep(v string) {
	o.Step = v
}

// GetIngredients returns the Ingredients field value if set, zero value otherwise.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetIngredients() []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner {
	if o == nil || o.Ingredients == nil {
		var ret []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner
		return ret
	}
	return o.Ingredients
}

// GetIngredientsOk returns a tuple with the Ingredients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetIngredientsOk() ([]AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner, bool) {
	if o == nil || o.Ingredients == nil {
		return nil, false
	}
	return o.Ingredients, true
}

// HasIngredients returns a boolean if a field has been set.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) HasIngredients() bool {
	if o != nil && o.Ingredients != nil {
		return true
	}

	return false
}

// SetIngredients gets a reference to the given []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner and assigns it to the Ingredients field.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) SetIngredients(v []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner) {
	o.Ingredients = v
}

// GetEquipment returns the Equipment field value if set, zero value otherwise.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetEquipment() []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner {
	if o == nil || o.Equipment == nil {
		var ret []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner
		return ret
	}
	return o.Equipment
}

// GetEquipmentOk returns a tuple with the Equipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetEquipmentOk() ([]AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner, bool) {
	if o == nil || o.Equipment == nil {
		return nil, false
	}
	return o.Equipment, true
}

// HasEquipment returns a boolean if a field has been set.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) HasEquipment() bool {
	if o != nil && o.Equipment != nil {
		return true
	}

	return false
}

// SetEquipment gets a reference to the given []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner and assigns it to the Equipment field.
func (o *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) SetEquipment(v []AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner) {
	o.Equipment = v
}

func (o AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["step"] = o.Step
	}
	if o.Ingredients != nil {
		toSerialize["ingredients"] = o.Ingredients
	}
	if o.Equipment != nil {
		toSerialize["equipment"] = o.Equipment
	}
	return json.Marshal(toSerialize)
}

type NullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner struct {
	value *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner
	isSet bool
}

func (v NullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) Get() *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner {
	return v.value
}

func (v *NullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) Set(val *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner(val *AnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) *NullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner {
	return &NullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner{value: val, isSet: true}
}

func (v NullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyzeRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


