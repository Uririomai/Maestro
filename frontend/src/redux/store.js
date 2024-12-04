import {configureStore} from '@reduxjs/toolkit'
import userReducer from './slices/userSlice'
import {useSelector, useDicpatch} from 'react-redux'

export const store = configureStore({
    reducer: {
        user: userReducer,
    },
})