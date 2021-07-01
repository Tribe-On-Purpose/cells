/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import InstallGetAgreementResponse from '../model/InstallGetAgreementResponse';
import InstallGetDefaultsResponse from '../model/InstallGetDefaultsResponse';
import InstallInstallEventsResponse from '../model/InstallInstallEventsResponse';
import InstallInstallRequest from '../model/InstallInstallRequest';
import InstallInstallResponse from '../model/InstallInstallResponse';
import InstallPerformCheckRequest from '../model/InstallPerformCheckRequest';
import InstallPerformCheckResponse from '../model/InstallPerformCheckResponse';

/**
* InstallService service.
* @module api/InstallServiceApi
* @version 1.0
*/
export default class InstallServiceApi {

    /**
    * Constructs a new InstallServiceApi. 
    * @alias module:api/InstallServiceApi
    * @class
    * @param {module:ApiClient} apiClient Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }



    /**
     * Load a textual agreement for using the software
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallGetAgreementResponse} and HTTP response
     */
    getAgreementWithHttpInfo() {
      let postBody = null;


      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = InstallGetAgreementResponse;

      return this.apiClient.callApi(
        '/install/agreement', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * Load a textual agreement for using the software
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallGetAgreementResponse}
     */
    getAgreement() {
      return this.getAgreementWithHttpInfo()
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Loads default values for install form
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallGetDefaultsResponse} and HTTP response
     */
    getInstallWithHttpInfo() {
      let postBody = null;


      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = InstallGetDefaultsResponse;

      return this.apiClient.callApi(
        '/install', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * Loads default values for install form
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallGetDefaultsResponse}
     */
    getInstall() {
      return this.getInstallWithHttpInfo()
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallInstallEventsResponse} and HTTP response
     */
    installEventsWithHttpInfo() {
      let postBody = null;


      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = InstallInstallEventsResponse;

      return this.apiClient.callApi(
        '/install/events', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallInstallEventsResponse}
     */
    installEvents() {
      return this.installEventsWithHttpInfo()
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Perform a check during install (like a valid DB connection)
     * @param {module:model/InstallPerformCheckRequest} body 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallPerformCheckResponse} and HTTP response
     */
    performInstallCheckWithHttpInfo(body) {
      let postBody = body;

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling performInstallCheck");
      }


      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = InstallPerformCheckResponse;

      return this.apiClient.callApi(
        '/install/check', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * Perform a check during install (like a valid DB connection)
     * @param {module:model/InstallPerformCheckRequest} body 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallPerformCheckResponse}
     */
    performInstallCheck(body) {
      return this.performInstallCheckWithHttpInfo(body)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Post values to be saved for install
     * @param {module:model/InstallInstallRequest} body 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallInstallResponse} and HTTP response
     */
    postInstallWithHttpInfo(body) {
      let postBody = body;

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling postInstall");
      }


      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = InstallInstallResponse;

      return this.apiClient.callApi(
        '/install', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType
      );
    }

    /**
     * Post values to be saved for install
     * @param {module:model/InstallInstallRequest} body 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallInstallResponse}
     */
    postInstall(body) {
      return this.postInstallWithHttpInfo(body)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


}