import React, { useState } from 'react'

import Sidebar from '../Sidebar/Sidebar'

import cabinetStyles from '../Cabinet/cabinet.module.scss'
import styles from './themePage.module.scss'
import Button from '../button/button'

const ThemePage = () => {
	const [isHover, setHoverIndex] = useState(null)

	const onMouseOver = num => {
		setHoverIndex(num)
	}
	const onMouseOut = () => {
		setHoverIndex(null)
	}

	return (
		<>
			<section className={cabinetStyles.personalCabinet + ` container`}>
				<Sidebar />
				<div className={styles.themepage}>
					<h2 className={styles.themepage__title}>Шаблоны сайта</h2>
					<ul className={styles.themepage__list}>
						<li
							className={styles.themepage__item}
							onMouseEnter={() => onMouseOver(1)}
							onMouseLeave={onMouseOut}
						>
							<figure className={styles.themepage__itemFigure}>
								<img
									className={styles.themepage__itemImg}
									src=''
									alt='img template'
								/>
								{isHover === 1 && (
									<div className={styles.themepage__itemButtons}>
										<Button
											width={117}
											height={36}
											colorText={'var(--color-black)'}
											colorBack={'var(--color-light)'}
											buttonText={'Посмотреть'}
										/>
										<Button
											width={97}
											height={36}
											colorBack={'var(--color-black)'}
											colorText={'var(--color-light)'}
											buttonText={'Выбрать'}
										/>
									</div>
								)}
								<figcaption className={styles.themepage__itemText}>
									Шаблон 1
								</figcaption>
							</figure>
						</li>
						<li
							className={styles.themepage__item}
							onMouseOver={() => onMouseOver(2)}
							onMouseOut={onMouseOut}
						>
							<figure className={styles.themepage__itemFigure}>
								<img
									className={styles.themepage__itemImg}
									src=''
									alt='img template'
								/>
								{isHover === 2 && (
									<div className={styles.themepage__itemButtons}>
										<Button
											width={117}
											height={36}
											colorText={'var(--color-black)'}
											colorBack={'var(--color-light)'}
											buttonText={'Посмотреть'}
										/>
										<Button
											width={97}
											height={36}
											colorBack={'var(--color-black)'}
											colorText={'var(--color-light)'}
											buttonText={'Выбрать'}
										/>
									</div>
								)}
								<figcaption className={styles.themepage__itemText}>
									Шаблон 1
								</figcaption>
							</figure>
						</li>
						<li
							className={styles.themepage__item}
							onMouseOver={() => onMouseOver(3)}
							onMouseOut={onMouseOut}
						>
							<figure className={styles.themepage__itemFigure}>
								<img
									className={styles.themepage__itemImg}
									src=''
									alt='img template'
								/>
								{isHover === 3 && (
									<div className={styles.themepage__itemButtons}>
										<Button
											width={117}
											height={36}
											colorText={'var(--color-black)'}
											colorBack={'var(--color-light)'}
											buttonText={'Посмотреть'}
										/>
										<Button
											width={97}
											height={36}
											colorBack={'var(--color-black)'}
											colorText={'var(--color-light)'}
											buttonText={'Выбрать'}
										/>
									</div>
								)}
								<figcaption className={styles.themepage__itemText}>
									Шаблон 1
								</figcaption>
							</figure>
						</li>
						<li
							className={styles.themepage__item}
							onMouseOver={() => onMouseOver(4)}
							onMouseOut={onMouseOut}
						>
							<figure className={styles.themepage__itemFigure}>
								<img
									className={styles.themepage__itemImg}
									src=''
									alt='img template'
								/>
								{isHover === 4 && (
									<div className={styles.themepage__itemButtons}>
										<Button
											width={117}
											height={36}
											colorText={'var(--color-black)'}
											colorBack={'var(--color-light)'}
											buttonText={'Посмотреть'}
										/>
										<Button
											width={97}
											height={36}
											colorBack={'var(--color-black)'}
											colorText={'var(--color-light)'}
											buttonText={'Выбрать'}
										/>
									</div>
								)}
								<figcaption className={styles.themepage__itemText}>
									Шаблон 1
								</figcaption>
							</figure>
						</li>
						<li
							className={styles.themepage__item}
							onMouseOver={() => onMouseOver(5)}
							onMouseOut={onMouseOut}
						>
							<figure className={styles.themepage__itemFigure}>
								<img
									className={styles.themepage__itemImg}
									src=''
									alt='img template'
								/>
								{isHover === 5 && (
									<div className={styles.themepage__itemButtons}>
										<Button
											width={117}
											height={36}
											colorText={'var(--color-black)'}
											colorBack={'var(--color-light)'}
											buttonText={'Посмотреть'}
										/>
										<Button
											width={97}
											height={36}
											colorBack={'var(--color-black)'}
											colorText={'var(--color-light)'}
											buttonText={'Выбрать'}
										/>
									</div>
								)}
								<figcaption className={styles.themepage__itemText}>
									Шаблон 1
								</figcaption>
							</figure>
						</li>
					</ul>
				</div>
			</section>
		</>
	)
}

export default ThemePage
