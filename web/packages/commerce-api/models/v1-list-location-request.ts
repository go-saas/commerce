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


import { V1LocationFilter } from './v1-location-filter';

/**
 * 
 * @export
 * @interface V1ListLocationRequest
 */
export interface V1ListLocationRequest {
    /**
     * 
     * @type {number}
     * @memberof V1ListLocationRequest
     */
    'pageOffset'?: number;
    /**
     * 
     * @type {number}
     * @memberof V1ListLocationRequest
     */
    'pageSize'?: number;
    /**
     * 
     * @type {string}
     * @memberof V1ListLocationRequest
     */
    'search'?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1ListLocationRequest
     */
    'sort'?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof V1ListLocationRequest
     */
    'fields'?: string;
    /**
     * 
     * @type {V1LocationFilter}
     * @memberof V1ListLocationRequest
     */
    'filter'?: V1LocationFilter;
}

