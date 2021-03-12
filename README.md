# burgers_api

Iimplemented several methods:

POST /v2/burger/create_burger
json body:
{
    "burger_name":"verygoodburger",
    "places":[
        {
            "place_name":"street_chefs",
            "burger_info":{
                "price":4.80,
                "supply":11,
                "date":"11-1992",
                "rating":3.7
            }
        },
        {
            "place_name":"boom_burgers",
            "burger_info":{
                "price":2.10,
                "supply":11,
                "date":"02-1994",
                "rating":10
            }
        },
        {
            "place_name":"aladin",
            "burger_info":{
                "price":6,
                "supply":11,
                "date":"10-1973",
                "rating":2.1
            }
        }
    ]
}

GET /v2/burger/random

GET /v2/burger?page=1&per_page=2&place=aladin_food
