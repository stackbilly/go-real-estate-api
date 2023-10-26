# Real Estate Api
This projects demonstrates a simple real estate api endpoint built with golang. You can use it to obtain real estate data for your applications or to train your model. 
## Table of Contents
- [Introduction](#introduction)
- [Tools](#tools)
- [Usage](#usage)
- [Contribution](#contributions)
- [License](#license)

## Introduction
APIs play a crucial role in the real estate industry by facilitating data integration and data sharing between platforms in real time.
In this project we build a real estate API endpoint which returns upto 1400 json entries of data. The data is obtained through scraping real esate listing websites and written into a mongodb collection.

## Tools
- [Colly](https://github.com/gocolly/colly) - a flexible framework for writing web crawlers in Go
- [mongoimport](https://github.com/stackbilly/mongo-import) - a go package to read and write csv data into mongodb

## Usage
```bash
curl -X GET http://localhost:8080/api/houses
```
sample json output
```json
[
  {
  "Id":"652bb8481be35e2471f72238",
  "Name":"4 Bed House in Nyari",
  "Description":"4 Bed House in Nyari",
  "Location":"Nyari, Westlands",
  "Price":"Price not communicated",
  "Url":"/listings/4-bedroom-house-for-rent-nyari-3623660",
  "Image":"https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg"
  },
  {
  "Id":"652bb8481be35e2471f72239",
  "Name":"4 Bed Townhouse with En Suite at Ridgeways",
  "Description":"4 Bed Townhouse with En Suite at Ridgeways",
  "Location":"Ridgeways","Price":"KSh 285,000",
  "Url":"/listings/4-bedroom-townhouse-for-rent-ridgeways-3639903",
  "Image":"https://i.roamcdn.net/prop/brk/listing-thumb-376w/343e31e67b6789be4d515198a0a9ddc2/-/prod-property-core-backend-media-brk/5760878/a3b24da6-4f29-4c18-9c64-d0dac122eb32.png"
  }
]
```
## Contributions

We welcome and appreciate contributions from the community. If you'd like to contribute to this project, please follow the correct guidelines.
We will review your PR and provide feedback. Your contributions will help improve and grow this project, and we are grateful for your support!
Thank you for considering contributing to this project.

## License
This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details
