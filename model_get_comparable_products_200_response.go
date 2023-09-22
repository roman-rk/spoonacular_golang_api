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

// GetComparableProducts200Response 
type GetComparableProducts200Response struct {
	ComparableProducts GetComparableProducts200ResponseComparableProducts `json:"comparableProducts"`
}

// NewGetComparableProducts200Response instantiates a new GetComparableProducts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetComparableProducts200Response(comparableProducts GetComparableProducts200ResponseComparableProducts) *GetComparableProducts200Response {
	this := GetComparableProducts200Response{}
	this.ComparableProducts = comparableProducts
	return &this
}

// NewGetComparableProducts200ResponseWithDefaults instantiates a new GetComparableProducts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetComparableProducts200ResponseWithDefaults() *GetComparableProducts200Response {
	this := GetComparableProducts200Response{}
	return &this
}

// GetComparableProducts returns the ComparableProducts field value
func (o *GetComparableProducts200Response) GetComparableProducts() GetComparableProducts200ResponseComparableProducts {
	if o == nil {
		var ret GetComparableProducts200ResponseComparableProducts
		return ret
	}

	return o.ComparableProducts
}

// GetComparableProductsOk returns a tuple with the ComparableProducts field value
// and a boolean to check if the value has been set.
func (o *GetComparableProducts200Response) GetComparableProductsOk() (*GetComparableProducts200ResponseComparableProducts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparableProducts, true
}

// SetComparableProducts sets field value
func (o *GetComparableProducts200Response) SetComparableProducts(v GetComparableProducts200ResponseComparableProducts) {
	o.ComparableProducts = v
}

func (o GetComparableProducts200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["comparableProducts"] = o.ComparableProducts
	}
	return json.Marshal(toSerialize)
}

type NullableGetComparableProducts200Response struct {
	value *GetComparableProducts200Response
	isSet bool
}

func (v NullableGetComparableProducts200Response) Get() *GetComparableProducts200Response {
	return v.value
}

func (v *NullableGetComparableProducts200Response) Set(val *GetComparableProducts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetComparableProducts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetComparableProducts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetComparableProducts200Response(val *GetComparableProducts200Response) *NullableGetComparableProducts200Response {
	return &NullableGetComparableProducts200Response{value: val, isSet: true}
}

func (v NullableGetComparableProducts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetComparableProducts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


