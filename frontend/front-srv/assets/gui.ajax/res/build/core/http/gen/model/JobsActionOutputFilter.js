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

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

var _ServiceQuery = require('./ServiceQuery');

var _ServiceQuery2 = _interopRequireDefault(_ServiceQuery);

/**
* The JobsActionOutputFilter model module.
* @module model/JobsActionOutputFilter
* @version 1.0
*/

var JobsActionOutputFilter = (function () {
    /**
    * Constructs a new <code>JobsActionOutputFilter</code>.
    * @alias module:model/JobsActionOutputFilter
    * @class
    */

    function JobsActionOutputFilter() {
        _classCallCheck(this, JobsActionOutputFilter);

        this.Query = undefined;
        this.Label = undefined;
        this.Description = undefined;
    }

    /**
    * Constructs a <code>JobsActionOutputFilter</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/JobsActionOutputFilter} obj Optional instance to populate.
    * @return {module:model/JobsActionOutputFilter} The populated <code>JobsActionOutputFilter</code> instance.
    */

    JobsActionOutputFilter.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobsActionOutputFilter();

            if (data.hasOwnProperty('Query')) {
                obj['Query'] = _ServiceQuery2['default'].constructFromObject(data['Query']);
            }
            if (data.hasOwnProperty('Label')) {
                obj['Label'] = _ApiClient2['default'].convertToType(data['Label'], 'String');
            }
            if (data.hasOwnProperty('Description')) {
                obj['Description'] = _ApiClient2['default'].convertToType(data['Description'], 'String');
            }
        }
        return obj;
    };

    /**
    * @member {module:model/ServiceQuery} Query
    */
    return JobsActionOutputFilter;
})();

exports['default'] = JobsActionOutputFilter;
module.exports = exports['default'];

/**
* @member {String} Label
*/

/**
* @member {String} Description
*/