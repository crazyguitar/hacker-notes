#!/usr/bin/env node

'use strict'

/* eslint-disable no-multi-spaces, key-spacing */

const fs      = require('fs')
const path    = require('path')
const util    = require('util')
const crypto  = require('crypto')
const request = require('request')
const program = require('commander')

program.on('--help', () => {
  const ex = `
  Example:

    $ s3-upload.js -a <access key> -s <secret key> -b <bucket> -r <region> file
  `
  console.log(ex)
})

program
  .version('0.0.1')
  .usage('[options] <file>')
  .option('-a, --access-key-id [id]', 'AWS access key id')
  .option('-s, --secret-access-key [key]', 'AWS secret access key')
  .option('-b, --bucket [bucket]', 'AWS bucket name')
  .option('-r, --region [region]', 'AWS region')
  .parse(process.argv)

const accessKeyId     = program.accessKeyId
const secretAccessKey = program.secretAccessKey
const region          = program.region
const bucket          = program.bucket
const file            = program.args[0]

try {
  if (!accessKeyId) {
    throw new Error('access key id cannot be empty')
  }
  if (!secretAccessKey) {
    throw new Error('secret access key cannot be empty')
  }
  if (!region) {
    throw new Error('region cannot be empty')
  }
  if (!bucket) {
    throw new Error('bucket cannot be empty')
  }
  if (!file) {
    throw new Error('file cannot be empty')
  }
} catch (e) {
  console.log(e.message)
  process.exit(1)
}

/* eslint valid-jsdoc: "error" */

/**
 * Create a sha256 hmac.
 * @param {string} key The key which uses for creating hmac.
 * @param {string} data The data which uses for createing hmac.
 * @param {string} encoding The encoding. ex: base64, hex.
 * @returns {string} Return the result of the hmac.
 */
function hmac (key, data, encoding) {
  return crypto
    .createHmac('sha256', key)
    .update(data, 'utf8')
    .digest(encoding)
}

/**
 * Convert time to YYYYMMDDThhmmssZ. ex: 20180520T014936Z
 * @param {Date} time A Date object.
 * @returns {string} Return the string of time.
 */
function toTime (time) {
  return new Date(time).toISOString().replace(/[:-]|\.\d{3}/g, '')
}

/**
 * Convert time to YYYYMMDD. ex: 20180520
 * @param {Date} time A Date object.
 * @returns {string} Return the string of time.
 */
function toDate (time) {
  return toTime(time).substring(0, 8)
}

/**
 * Create a credential: <aws access key id>/region/service/aws4_request
 * @param {string} accessKey The access key id.
 * @param {Date} time The date object.
 * @param {string} region The region of aws. ex: us-west-1
 * @param {string} service The aws service. ex: s3
 * @returns {string} Return the credential.
 */
function credential (accessKey, time, region, service) {
  return [accessKey, toDate(time), region, service, 'aws4_request'].join('/')
}

/**
 * Create a AWS Signature Version 4
 * @param {string} secret The secret access key.
 * @param {Date} time The date object.
 * @param {string} region The region of aws. ex: us-west-1
 * @param {service} service the aws service. ex: s3
 * @param {string} str2Sign The string which is used to sign.
 * @returns {string} Return the result of signature.
 */
function signature (secret, time, region, service, str2Sign) {
  const kDate = hmac('AWS4' + secret, toDate(time))
  const kRegion = hmac(kDate, region)
  const kService = hmac(kRegion, service)
  const kSigning = hmac(kService, 'aws4_request')
  return hmac(kSigning, str2Sign, 'hex')
}

/**
 * POST file to s3.
 * @param {string} url The s3 bucket url.
 * @param {object} formData The http form-data.
 * @returns {undefined} Not return
 */
async function upload (url, formData) {
  const httpost = util.promisify(request.post)
  try {
    await httpost({ url: url, formData: formData })
  } catch (e) {
    console.error(e.message)
    process.exit(1)
  }
}

const date   = new Date()
const expire = new Date(date)

// expire at 1 hour later
expire.setHours(expire.getHours() + 1)

const policy = {
  expiration: expire.toISOString(),
  conditions: [
    { acl: 'private' },
    { bucket: bucket },
    ['starts-with', '$key', ''],
    ['content-length-range', 1, 10485760], // 1 Byte to 10 MB
    { 'x-amz-date': toTime(date) },
    { 'x-amz-credential': credential(accessKeyId, expire, region, 's3') },
    { 'x-amz-algorithm': 'AWS4-HMAC-SHA256' }
  ]
}

const str2Sign = Buffer.from(JSON.stringify(policy)).toString('base64')
const sig      = signature(secretAccessKey, expire, region, 's3', str2Sign)
const key      = path.basename(file)
const formData = {
  key               : key,
  acl               : 'private',
  'x-amz-credential': credential(accessKeyId, expire, region, 's3'),
  'x-amz-signature' : sig,
  'x-amz-algorithm' : 'AWS4-HMAC-SHA256',
  'x-amz-date'      : toTime(date),
  Policy            : str2Sign,
  file              : fs.createReadStream(file)
}

const url = `https://${bucket}.s3.amazonaws.com/`

upload(url, formData)
