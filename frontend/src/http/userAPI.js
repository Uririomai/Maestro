import { $authHost, $host } from '.'
import { jwtDecode } from 'jwt-decode'

export const registration = async (email, password) => {
	const { data } = await $host.post('api/admin/sign-up', { email, password })
	return jwtDecode(data.token)
}

export const login = async (email, password) => {
	const { data } = await $host.post('api/admin/sign-in', { email, password })
	return jwtDecode(data.token)
}

export const check = async (email, password) => {
	const response = await $host.post('api/admin/sign-up')
	return response
}
