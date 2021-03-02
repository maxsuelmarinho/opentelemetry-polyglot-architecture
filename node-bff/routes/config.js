import express from 'express';
const router = express.Router();
import {
  getPayPal,
} from '../controllers/config.js';

router.route('/paypal').get(getPayPal);

export default router;
