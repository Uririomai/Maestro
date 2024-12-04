import React, { Component } from 'react'
import { Route, Routes, Navigate } from 'react-router-dom'
import { authRoutes, publicRoutes, ROUTES } from '../../utils/routes'

const AppRoutes = () => {
	const isAuth = true

	return (
		<Routes>
			{isAuth &&
				authRoutes.map(({ path, Component }) => (
					<Route exact path={path} Component={Component} />
				))}
			{publicRoutes.map(({ path, Component }) => (
				<Route exact path={path} Component={Component} />
			))}
			
		</Routes>
	)
}

export default AppRoutes
