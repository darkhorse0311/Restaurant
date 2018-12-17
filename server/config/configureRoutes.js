const restaurantsRouter = require('../routes/restaurantsRouter');

module.exports = server => {
    server.get('/', (req, res) => {
        res.status(200).json({server: 'alive'})
    })

    server.use('/api/res', restaurantsRouter);
}