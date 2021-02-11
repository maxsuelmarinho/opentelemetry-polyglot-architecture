import express from 'express';
const router = express.Router();
import { authUser, getUserProfile, registerUser } from '../controllers/user.js';
import { protect } from '../middleware/auth.js';

router.route('/').post(registerUser);
router.post('/login', authUser);
router.route('/profile').get(protect, getUserProfile);

export default router;