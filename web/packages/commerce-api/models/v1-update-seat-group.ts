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


import { V1Seat } from './v1-seat';

/**
 * 
 * @export
 * @interface V1UpdateSeatGroup
 */
export interface V1UpdateSeatGroup {
    /**
     * 
     * @type {string}
     * @memberof V1UpdateSeatGroup
     */
    'name'?: string;
    /**
     * 
     * @type {Array<V1Seat>}
     * @memberof V1UpdateSeatGroup
     */
    'seats'?: Array<V1Seat>;
}

