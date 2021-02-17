import express from 'express';
const router = express.Router();
import {
    getProducts, 
    getProductById, 
    createProductReview
} from '../controllers/product.js';
import { protect } from '../middleware/auth.js';

router.route('/').get(getProducts);
router.route('/:id').get(getProductById);
router.route('/:id/reviews').post(protect, createProductReview);

export default router;