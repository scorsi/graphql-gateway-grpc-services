module.exports = {
    resolvers: {
        ProductModel: {
            reviews: {
                selectionSet: `{ id }`,
                resolve: async ({id: productId}, _args, {ReviewService}) => {
                    return (await ReviewService.api.reviewServiceGetReviewsByProduct({input: {productId}})).reviews;
                }
            }
        }
    }
}