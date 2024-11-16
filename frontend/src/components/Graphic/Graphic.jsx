import React from 'react'
import { LineChart } from '@mui/x-charts/LineChart'


import styles from './graphic.module.scss'

const Graphic = () => {
	const time = [
		new Date(2025, 1, 0),
		new Date(2025, 2, 0),
		new Date(2025, 3, 0),
		new Date(2025, 4, 0),
		new Date(2025, 5, 0),
		new Date(2025, 6, 0),
		new Date(2025, 7, 0),
		new Date(2025, 8, 0),
		new Date(2025, 9, 0),
		new Date(2025, 10, 0),
	]

	return (
		<LineChart
			className={styles}
			xAxis={[
				{
					scaleType: 'time',
					data: time,
					datakey: 'month',
					min: time[0].getTime(),
					max: time[time.length - 1].getTime(),
				},
			]}
			series={[
				{
					curve: 'linear',
					data: [2, 5.5, 3, 5, 2, 8.5, 1.5, 5, 7, 8],
					area: true,
					color: '#000000',
					showMark: false,
				},
			]}
			width={500}
			height={300}
			grid={{ horizontal: true }}
			RightAxis={{
				tickLabelStyle: {
					fontFamily: 'PublicSans',
					fontSize: 13,
					lineHeight: 16,
				},
			}}
			bottomAxis={{
				tickLabelStyle: {
					fontFamily: 'PublicSans',

					fontSize: 14,
				},
			}}
			sx={{
				'& .MuiLineElement-root': {
					strokeWidth: 2,
				},
				'& .MuiAreaElement-root': {
					fill: 'url(#myGradient)',
				},
			}}
		>
			<defs>
				<linearGradient id='myGradient' gradientTransform='rotate(90)'>
					<stop offset='5%' stopColor='#000000' />
					<stop offset='100%' stopColor='rgba(252, 252, 252, 0)' />
				</linearGradient>
			</defs>
		</LineChart>
	)
}

export default Graphic
