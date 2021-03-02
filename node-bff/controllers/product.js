import asyncHandler from 'express-async-handler';
import axios from 'axios';
//import Product from '../models/product.js';
// @desc    Fetch all products
// @route   GET /api/products
// @access  Public
const getProducts = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json',
    }
  };

  try {
    const { data } = await axios.get(`${process.env.PRODUCT_SERVICE_URL}${req.url}`, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});


// @desc    Fetch single product
// @route   GET /api/products/:id
// @access  Public
const getProductById = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json',
    }
  };

  try {
    const { data } = await axios.get(`${process.env.PRODUCT_SERVICE_URL}/${req.params.id}`, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

// @desc    Create new review
// @route   POST /api/products/:id/reviews
// @access  Private
const createProductReview = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json',
        Authorization: req.headers.authorization,
    }
  };

  try {
    const body = req.body;
    const { data } = await axios.post(`${process.env.PRODUCT_SERVICE_URL}/${req.params.id}/reviews`, body, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

// @desc    Get Top Products
// @route   GET /api/products/top
// @access  Public
const getTopProducts = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json',
    }
  };

  try {
    const { data } = await axios.get(`${process.env.PRODUCT_SERVICE_URL}/top`, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

export {
    getProducts,
    getProductById,
    createProductReview,
    getTopProducts
};
