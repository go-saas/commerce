/* tslint:disable */
/* eslint-disable */
/**
 * order/api/order/v1/order.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { V1UpdateLocation } from './v1-update-location';

/**
 * 
 * @export
 * @interface V1UpdateLocationRequest
 */
export interface V1UpdateLocationRequest {
    /**
     * 
     * @type {V1UpdateLocation}
     * @memberof V1UpdateLocationRequest
     */
    'location'?: V1UpdateLocation;
    /**
     * 
     * @type {string}
     * @memberof V1UpdateLocationRequest
     */
    'updateMask'?: string;
}

