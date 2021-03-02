import asyncHandler from 'express-async-handler';
import axios from 'axios';

const protect = asyncHandler(async (req, res, next) => {
  let token;
  if (req.headers.authorization && req.headers.authorization.startsWith('Bearer')) {
      try {
          token = req.headers.authorization.split(' ')[1];
          const config = {
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            }
          };
          const { data } = await axios.get(`${process.env.USER_SERVICE_URL}/profile`, config);
          req.user = data;
          next();
      } catch (error) {
          res.status(401);
          throw new Error('Not authorized, token failed');
      }
  }

  if (!token) {
      res.status(401);
      throw new Error('Not authorized, no token');
  }
});

export { protect };
