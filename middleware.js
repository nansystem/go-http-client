module.exports = (req, res, next) => {
  console.info('req.headers:', req.headers)
  next()
}
