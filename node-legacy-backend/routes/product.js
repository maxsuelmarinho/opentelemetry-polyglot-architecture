import express from 'express';
const router = express.Router();
import {
    getProducts, 
    getProductById, 
    createProductReview,
    getTopProducts,
} from '../controllers/product.js';
import { protect } from '../middleware/auth.js';

router.route('/').get(getProducts);
router.route('/top').get(getTopProducts);
router.route('/:id').get(getProductById);
router.route('/:id/reviews').post(protect, createProductReview);

export default router;