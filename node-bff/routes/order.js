import express from 'express';
const router = express.Router();
import { 
    addOrderItems,
    getOrderById,
    updateOrderToPaid,
    getMyOrders,
} from '../controllers/order.js';
import { protect } from '../middleware/auth.js';

router.route('/').post(protect, addOrderItems);
router.route('/myorders').get(protect, getMyOrders);
router.route('/:id').get(protect, getOrderById);
router.route('/:id/pay').put(protect, updateOrderToPaid);

export default router;