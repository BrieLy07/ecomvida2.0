const AWS = require('aws-sdk');
require('dotenv').config();

const dynamo = new AWS.DynamoDB.DocumentClient({
  region: process.env.AWS_REGION,
  endpoint: process.env.DYNAMO_ENDPOINT
});

module.exports = dynamo;
