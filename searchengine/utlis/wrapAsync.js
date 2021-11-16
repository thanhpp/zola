module.exports = (func) => {
  return (req, res, next) => {
    func(req, res, next).catch((e) => {
      next(e);
      // console.log(e)
      // return res.status(e.statusCode||500).send({msg:'Error',error:e})
    });
  };
};
