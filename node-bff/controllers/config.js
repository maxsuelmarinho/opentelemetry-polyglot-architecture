import asyncHandler from 'express-async-handler';
import axios from 'axios';

// @desc    Fetch PayPal config
// @route   GET /api/config/paypal
// @access  Public
const getPayPal = asyncHandler(async (req, res) => {
  console.log("getPayPal");
  const config = {
    headers: {
        'Content-Type': 'application/json',
    }
  };

  try {
    const { data } = await axios.get(`${process.env.CONFIG_SERVICE_URL}/paypal`, config);
    res.send(data);
  } catch (error) {
    console.error("error:", error);
    res.status(error.response.status);
    throw createError(error);
  }
});

export {
  getPayPal,
};
