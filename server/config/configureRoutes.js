module.exports = server => {
    server.get('/', (req, res) => {
        res.status(200).json({server: 'alive'})
    })
}