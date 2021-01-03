# Request & Response Examples

## API Resources

* [GET /get](https://murmuring-harbor-28221.herokuapp.com/TFSA/get)
* [GET /get/[year]](https://murmuring-harbor-28221.herokuapp.com/TFSA/get/2020)
* [GET /refresh](https://murmuring-harbor-28221.herokuapp.com/TFSA/refresh)

### GET /get

Example: https://murmuring-harbor-28221.herokuapp.com/TFSA/get

Response body:
```json
{
  "2009": 5000,
  "2010": 5000,
  "2011": 5000,
  "2012": 5000,
  "2013": 5500,
  "2014": 5500,
  "2015": 10000,
  "2016": 5500,
  "2017": 5500,
  "2018": 5500,
  "2019": 6000,
  "2020": 6000,
  "2021": 6000
}
```

### GET /get/[year]
Example: https://murmuring-harbor-28221.herokuapp.com/TFSA/get/2020

Response body:
```
6000
```

### GET /refresh/
Example: https://murmuring-harbor-28221.herokuapp.com/TFSA/refresh

Response body:
```
Refreshed
```

