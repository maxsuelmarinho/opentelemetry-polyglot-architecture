import asyncHandler from 'express-async-handler';
import axios from 'axios';

const createError = (error) => {
  return new Error(error.response && error.response.data.message
    ? error.response.data.message
    : error.message);
};
// @desc    Auth user & get token
// @route   POST /api/users/login
// @access  Public
const authUser = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json'
    }
  };

  try {
    const body = req.body;
    const { data } = await axios.post(`${process.env.USER_SERVICE_URL}/login`, body, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

// @desc    Register a new user
// @route   POST /api/users
// @access  Public
const registerUser = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json'
    }
  };

  try {
    const body = req.body;
    const { data } = await axios.post(`${process.env.USER_SERVICE_URL}/`, body, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

// @desc    Get user profile
// @route   GET /api/users/profile
// @access  Private
const getUserProfile = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json',
        Authorization: req.headers.authorization,
    }
  };

  try {
    const { data } = await axios.get(`${process.env.USER_SERVICE_URL}/profile`, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

// @desc    Update user profile
// @route   PUT /api/users/profile
// @access  Private
const updateUserProfile = asyncHandler(async (req, res) => {
  const config = {
    headers: {
        'Content-Type': 'application/json',
        Authorization: req.headers.authorization
    }
  };

  try {
    const body = req.body;
    const { data } = await axios.put(`${process.env.USER_SERVICE_URL}/profile`, body, config);
    res.json(data);
  } catch (error) {
    res.status(error.response.status);
    throw createError(error);
  }
});

export { authUser, getUserProfile, updateUserProfile, registerUser };
