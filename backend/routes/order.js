import express from 'express';
const router = express.Router();
import { 
    addOrderItems,
    getOrderById,
} from '../controllers/order.js';
import { protect } from '../middleware/auth.js';

router.route('/').post(protect, addOrderItems);
router.route('/:id').get(protect, getOrderById);

export default router;