// Tiny helper to always return a date 10 minutes in the future.
module.exports = () => {
  let d = new Date();
  d.setMinutes(d.getMinutes() + 10);
  return d;
};
