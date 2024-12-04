import { useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import { ROUTES } from '../utils/routes'

export function useAuth() {
	const isAuth = useSelector(state => state.user.isAuth)
   
	console.log('aut ')
	console.log(isAuth)

	return isAuth
}
