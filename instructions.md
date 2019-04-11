## Problem Statement: Create a simple ad exchange.

An Ad exchange is a platform where publishers and advertisers sell and buy ads. Think of the Publisher as a website owner who wants to make some revenue by displaying ads on their website. For example, New York Times, StackOverflow, etc and think of the Advertiser as the one who wants to buy these ad spaces and display their ads. For example, Nike, Apple, etc. So in summary, Publishers are the sellers of ad spaces (also known as ad placements) and advertisers are the buyers. There can usually be multiple other parties involved in these transactions but in order to keep this simple, let's assume we just have publishers and advertisers and they talk directly to our ad exchange service.

You need to design and implement a REST API endpoint that a publisher can call to sell his ad placements. This endpoint will conduct an auction between all the supported buyers by the ad exchange. For simplicity, we will assume that the Publisher can only display *image* ads.

- The request will contain one or more ad placements that the publisher wants to sell. An ad placement consists of an ad placement ID and the different image sizes supported by the particular ad placement.
- The ad exchange can then contact different buyers and ask them to bid on the ad placements. For simplicity, in this case we will assume that there exists 3 buyers (Nike, Apple and Amazon) that the ad exchange supports. Ideally we'd make a call to these buyers to get the bids but for this example, assume you can simply mock this functionality such that buyers will always send a rndomly generated bid response.

**Note:** The BidResponse object definition can be found in bid_response.go

- The ad exchange then needs to select the winning bid for each ad placement and return that to the Publisher.
- Remember that an ad placement will only be able to display an ad properly if it is in one of the sizes that is supported by that ad placement.

### Assumptions that you can make:
- Each bidder bids on each ad placement that was requested.
- Prices are in dollars and only a single currency is supported.

**Note:** Feel free to structure/restructure the code as you see fit. You are also allowed to change the request, response objects if you feel the need to.

### Example

* Say, NYT sends the following request to the ad exchange:
```json
{
    "ad_placements": [
        {
            "id": 1,
            "size": [
                {
                    "width": 200,
                    "height": 400
                },
                {
                    "width": 350,
                    "height": 450
                }
            ]
        },
        {
            "id": 2,
            "size": [
                {
                    "width": 200,
                    "height": 400
                },
                {
                    "width": 750,
                    "height": 850
                }
            ]
        },
        {
            "id": 3,
            "size": [
                {
                    "width": 250,
                    "height": 500
                },
                {
                    "width": 350,
                    "height": 400
                }
            ]
        }
    ]
}
```

* Now say, Nike responds with:
```json
{
	"bidder_name": "Nike",
	"bids": [{
			"placement_id": 1,
			"price": 1.10,
			"ad_creatives": [{
					"size": {
						"width": 200,
						"height": 400
					},
					"url": "http://www.nike.com/img/size?w=200&h=400"
				},
				{
					"size": {
						"width": 350,
						"height": 400
					},
					"url": "http://www.nike.com/img/size?w=350&h=400"
				}
			]
		},
		{
			"placement_id": 2,
			"price": 2.00,
			"ad_creatives": [{
					"size": {
						"width": 150,
						"height": 300
					},
					"url": "http://www.nike.com/img/size?w=150&h=300"
				},
				{
					"size": {
						"width": 200,
						"height": 400
					},
					"url": "http://www.nike.com/img/size?w=200&h=400"
				}
			]
		},
		{
			"placement_id": 3,
			"price": 0.50,
			"ad_creatives": [{
					"size": {
						"width": 250,
						"height": 500
					},
					"url": "http://www.nike.com/img/size?w=250&h=500"
				},
				{
					"size": {
						"width": 350,
						"height": 400
					},
					"url": "http://www.nike.com/img/size?w=350&h=400"
				}
			]
		}
	]
}
```

* Apple responds with:
```json
{
    "bidder_name": "Apple",
    "bids": [
        {
            "placement_id": 1,
            "price": 2.10,
            "ad_creatives": [
                {
                    "size": {
                        "width": 750,
                        "height": 850
                    },
                    "url": "http://www.apple.com/img/size?w=750&h=850"
                }
            ]
        },
        {
            "placement_id": 2,
            "price": 1.00,
            "ad_creatives": [
                {
                    "size": {
                        "width": 750,
                        "height": 850
                    },
                    "url": "http://www.apple.com/img/size?w=750&h=850"
                }
            ]
        },
        {
            "placement_id": 3,
            "price": 5.30,
            "ad_creatives": [
                {
                    "size": {
                        "width": 750,
                        "height": 850
                    },
                    "url": "http://www.apple.com/img/size?w=750&h=850"
                }
            ]
        }
    ]
}
```

* Amazon responds with:
```json
{
    "bidder_name": "Amazon",
    "bids": [
        {
            "placement_id": 1,
            "price": 1.20,
            "ad_creatives": [
                {
                    "size": {
                        "width": 200,
                        "height": 400
                    },
                    "url": "http://www.amazon.com/img/size?w=200&h=400"
                },
                {
                    "size": {
                        "width": 350,
                        "height": 400  
                    },
                    "url": "http://www.amazon.com/img/size?w=350&h=400"
                }
            ]
        },
        {
            "placement_id": 2,
            "price": 0.20,
            "ad_creatives": [
                {
                    "size": {
                        "width": 750,
                        "height": 850  
                    },
                    "url": "http://www.amazon.com/img/size?w=750&h=850"
                }
            ]
        },
        {
            "placement_id": 3,
            "price": 3.50,
            "ad_creatives": [
                {
                    "size": {
                        "width": 150,
                        "height": 500  
                    },
                    "url": "http://www.amazon.com/img/size?w=150&h=500"
                },
                {
                    "size": {
                        "width": 350,
                        "height": 400  
                    },
                    "url": "http://www.amazon.com/img/size?w=350&h=400"
                }
            ]
        }
    ]
}
```

* The ad exchange would then return the following:
```json
{
    "ad_placements": [
        {
            "placement_id": 1,
            "bid_price": 1.20,
            "url": "http://www.amazon.com/img/size?w=200&h=400"
        },
        {
            "placement_id": 2,
            "bid_price": 2.00,
            "url": "http://www.nike.com/img/size?w=200&h=400"
        },
        {
            "placement_id": 3,
            "bid_price": 3.50,
            "url": "http://www.amazon.com/img/size?w=350&h=400"
        }
    ]
}
```

### Implementation

* Since you will be writing GoLang as part of this role, we've provided a GoLang boilerplate that you can use as a starting point to implement this adEx service.
* We understand that GoLang might not be everyone's preferred language and which is why you're welcome to use any other language of your choice too. Although, in that case you'd have to start from scratch.
* Remember that we're looking for how you think and go about designing a service and language isn't a barrier :)
* Once you've implemented the service, make sure to add a `Readme.md` describing any necessary information that we will need to review it. For example, the endpoint that you designed for conducting an auction for a publisher's ad placements.

### The GoLang boilerplate

* In order to use the GoLang boilerplate, make sure you've Go setup on your computer and you've placed the project in the `$GOPATH/src/` directory.
* To compile and run the adEx service, run the following commands: 

```bash
# Install all dependencies using a dependency management tool called dep
dep ensure
# Compile the code into a Go binary
go build
# Finally, run the service
./adEx -port=$ADEX_PORT
```

* Now that you have the adEx service running you can simply use CURL to hit any of the service endpoints. There are two endpoints that are already defined. A `/health` endpoint that's basically a healthcheck endpoint and a `/bidders` endpoint that returns the list of supported bidders by the adEx service.

For example, you can run the following to hit the healthcheck endpoint:
```bash
curl http://localhost:$ADEX_PORT/health
```

This should return:
```bash
Hello World!
```

### Further Questions

* If you have any further questions or need further clarifications, please feel free to reach out to us.