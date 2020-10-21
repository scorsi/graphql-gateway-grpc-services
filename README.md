# graphql-gateway-grpc-services

To make it up and running:
 - `npm install`
 - `make run_products` and wait the compilation (may be long)
 - `make run_reviews` and wait the compilation (this time, very short)
 - `make run_gateway`
 
Now, by visiting [localhost:4000](http://localhost:4000/) you can now execute the following query:
```graphql
{
  productServiceGetProducts(input: {productIds: ["1"]}) {
    products {
      title
      reviews {
        userName
        rating
      }
    }
  }
}
```
