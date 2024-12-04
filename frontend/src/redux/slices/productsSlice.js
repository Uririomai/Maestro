import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { $host } from '../../http'

const initialState = {
	list: [],
	isLoading: null,
}

export const getProducts = createAsyncThunk(
	'product/getProducts',
	async (_, thunkAPI) => {
		try {
			const res = await $host.get('api/product/get-active-by-alias')
			return res.data
		} catch (err) {
			alert(err.response.data.error)
			console.log(err)
			return thunkAPI.rejectWithValue(err)
		}
	}
)

const productsSlice = createSlice({
	name: 'products',
	initialState,

	extraReducers: builder => {
		builder
			.addCase(getProducts.pending, state => {
				state.status = 'loading'
			})
			.addCase(getProducts.fulfilled, (state, action) => {
				state.list = action.payload
				state.status = 'resolved'
			})
			.addCase(getProducts.rejected, state => {
				state.status = 'rejected'
				state.isAuth = false
			})
	},
})

export default productsSlice.reducer
