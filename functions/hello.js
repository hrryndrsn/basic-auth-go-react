// functions/hello.js
const result = require("dotenv").config();

exports.handler = async event => {
  const subject = event.queryStringParameters.name || "World";
  console.log(process.env);

  return {
    statusCode: 200,
    body: `Send email to hello \n\n message \n your mum`
  };
};
