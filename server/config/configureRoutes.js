const restaurantsRouter = require('../routes/restaurantsRouter');
const authRouter = require('../routes/authRouter');

module.exports = server => {
    server.get('/', (req, res) => {
        res.status(200).json({server: 'alive'})
    })

    server.use('/api/res', restaurantsRouter);
    server.use('/auth', authRouter);
}