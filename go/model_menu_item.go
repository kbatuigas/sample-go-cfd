/*
 * Continous Food Delievery
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MenuItem struct {
	Id int32 `json:"id"`

	Description string `json:"description"`

	Name string `json:"name"`

	Price float32 `json:"price"`

	// URL to an image of the menu item.  This should be the image from the /image endpoint
	ImageId int32 `json:"imageId"`

	ImageName string

	Image []byte
}
