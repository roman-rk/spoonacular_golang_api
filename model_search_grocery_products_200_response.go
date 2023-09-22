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

// SearchGroceryProducts200Response 
type SearchGroceryProducts200Response struct {
	Products []AutocompleteRecipeSearch200ResponseInner `json:"products"`
	TotalProducts int32 `json:"totalProducts"`
	Type string `json:"type"`
	Offset int32 `json:"offset"`
	Number int32 `json:"number"`
}

// NewSearchGroceryProducts200Response instantiates a new SearchGroceryProducts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchGroceryProducts200Response(products []AutocompleteRecipeSearch200ResponseInner, totalProducts int32, type_ string, offset int32, number int32) *SearchGroceryProducts200Response {
	this := SearchGroceryProducts200Response{}
	this.Products = products
	this.TotalProducts = totalProducts
	this.Type = type_
	this.Offset = offset
	this.Number = number
	return &this
}

// NewSearchGroceryProducts200ResponseWithDefaults instantiates a new SearchGroceryProducts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchGroceryProducts200ResponseWithDefaults() *SearchGroceryProducts200Response {
	this := SearchGroceryProducts200Response{}
	return &this
}

// GetProducts returns the Products field value
func (o *SearchGroceryProducts200Response) GetProducts() []AutocompleteRecipeSearch200ResponseInner {
	if o == nil {
		var ret []AutocompleteRecipeSearch200ResponseInner
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *SearchGroceryProducts200Response) GetProductsOk() ([]AutocompleteRecipeSearch200ResponseInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *SearchGroceryProducts200Response) SetProducts(v []AutocompleteRecipeSearch200ResponseInner) {
	o.Products = v
}

// GetTotalProducts returns the TotalProducts field value
func (o *SearchGroceryProducts200Response) GetTotalProducts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalProducts
}

// GetTotalProductsOk returns a tuple with the TotalProducts field value
// and a boolean to check if the value has been set.
func (o *SearchGroceryProducts200Response) GetTotalProductsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalProducts, true
}

// SetTotalProducts sets field value
func (o *SearchGroceryProducts200Response) SetTotalProducts(v int32) {
	o.TotalProducts = v
}

// GetType returns the Type field value
func (o *SearchGroceryProducts200Response) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SearchGroceryProducts200Response) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SearchGroceryProducts200Response) SetType(v string) {
	o.Type = v
}

// GetOffset returns the Offset field value
func (o *SearchGroceryProducts200Response) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *SearchGroceryProducts200Response) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *SearchGroceryProducts200Response) SetOffset(v int32) {
	o.Offset = v
}

// GetNumber returns the Number field value
func (o *SearchGroceryProducts200Response) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *SearchGroceryProducts200Response) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *SearchGroceryProducts200Response) SetNumber(v int32) {
	o.Number = v
}

func (o SearchGroceryProducts200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["products"] = o.Products
	}
	if true {
		toSerialize["totalProducts"] = o.TotalProducts
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableSearchGroceryProducts200Response struct {
	value *SearchGroceryProducts200Response
	isSet bool
}

func (v NullableSearchGroceryProducts200Response) Get() *SearchGroceryProducts200Response {
	return v.value
}

func (v *NullableSearchGroceryProducts200Response) Set(val *SearchGroceryProducts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchGroceryProducts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchGroceryProducts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchGroceryProducts200Response(val *SearchGroceryProducts200Response) *NullableSearchGroceryProducts200Response {
	return &NullableSearchGroceryProducts200Response{value: val, isSet: true}
}

func (v NullableSearchGroceryProducts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchGroceryProducts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


