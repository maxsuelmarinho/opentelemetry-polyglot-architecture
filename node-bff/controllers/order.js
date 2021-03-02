import asyncHandler from 'express-async-handler';
import axios from 'axios';

const createError = (error) => {
  return new Error(error.response && error.response.data.message
    ? error.response.data.message
    : error.message);
};

// @desc    Create new order
// @route   POST /api/orders
// @access  Private
const addOrderItems = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json'
    }
  };

  try {
    const body = {...req.body, userId: req.user._id};
    const response = await axios.post(`${process.env.ORDER_SERVICE_URL}/`, body, config);
    console.log(response.data);
    res.status(response.status);
    res.json(response.data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

// @desc    Get order by ID
// @route   GET /api/orders/:id
// @access  Private
const getOrderById = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json'
    }
  };

  try {
    const { data } = await axios.get(`${process.env.ORDER_SERVICE_URL}/${req.params.id}`, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

// @desc    Update order to paid
// @route   PUT /api/orders/:id/pay
// @access  Private
const updateOrderToPaid = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json'
    }
  };

  try {
    const body = req.body;
    const { data } = await axios.put(`${process.env.ORDER_SERVICE_URL}/${req.params.id}/pay`, body, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

// @desc    Get logged in user orders
// @route   GET /api/orders/myorders
// @access  Private
const getMyOrders = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json'
    }
  };

  try {
    const { data } = await axios.get(`${process.env.ORDER_SERVICE_URL}/?user=${req.user._id}`, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

export { addOrderItems, getOrderById, updateOrderToPaid, getMyOrders };
